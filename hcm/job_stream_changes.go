package hcm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/config"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

func (s *Server) StreamJobIntoBigQuery(c *gin.Context) {
	s.Logger.Info("Stream job Data Into BigQuery")

	config.CreateConfig(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
	session := &types.Session{}

	s.Logger.Info("initDB")
	err := config.Init(session)
	if err != nil {
		e := fmt.Sprintf("unable to init session")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	err = resetSession(session)
	if err != nil {
		e := fmt.Sprintf("Could not create replication connection")
		s.Logger.WithError(err).Error(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": e,
		})
		return
	}

	go func() {
		wsErr := make(chan error, 1)

		s.Logger.Infof("Starting replication for slot '%s' from LSN %s", session.SlotName, pgx.FormatLSN(session.RestartLSN))
		err = session.ReplConn.StartReplication(session.SlotName, session.RestartLSN, -1, "\"include-lsn\" 'on'", "\"pretty-print\" 'off'")
		if err != nil {
			e := fmt.Sprintf("Could not Start replication for slot '%s' from LSN %s", session.SlotName, pgx.FormatLSN(session.RestartLSN))
			s.Logger.WithError(err).Error(e)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": e,
			})
			return
		}

		// start sending periodic status heartbeats to postgres
		go config.SendPeriodicHeartbeats(session)

		for {
			if !session.ReplConn.IsAlive() {
				e := fmt.Sprintf("Looks like the connection is dead")
				s.Logger.WithError(session.ReplConn.CauseOfDeath()).Error(e)
			}
			s.Logger.Info("Waiting for LR message")

			ctx := session.Ctx
			message, err := session.ReplConn.WaitForReplicationMessage(ctx)
			if err != nil {
				// check whether the error is because of the context being cancelled
				if ctx.Err() != nil {
					// context cancelled, exit
					s.Logger.Warn("Websocket closed")
					return
				}

				s.Logger.WithError(err).Errorf("%s", reflect.TypeOf(err))
			}

			if message.WalMessage != nil {
				if message == nil {
					s.Logger.Error("Message nil")
					continue
				}
				walData := message.WalMessage.WalData
				s.Logger.Infof("Received replication message: %s", string(walData))

				var wData types.WalData
				if err := json.Unmarshal(walData, &wData); err != nil {
					e := fmt.Sprintf("faild to decode waldata:%s", string(walData))
					s.Logger.WithError(session.ReplConn.CauseOfDeath()).Error(e)
				}

				for _, value := range wData.Change {
					if value.Kind == "delete" {
						continue
					}

					if value.Table == "job" {
						u := s.BigQueryClient.Dataset("hcm").Table("job").Inserter()
						item := ParseValueToJob(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert job data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "job_competency" {
						u := s.BigQueryClient.Dataset("hcm").Table("job_competency").Inserter()
						item := ParseValueToJobCompetency(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert job data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "job_education_specialization" {
						u := s.BigQueryClient.Dataset("hcm").Table("job_education_specialization").Inserter()
						item := ParseValueToJobEducationSpecialization(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert job data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "job_language" {
						u := s.BigQueryClient.Dataset("hcm").Table("job_language").Inserter()
						item := ParseValueToJobLanguage(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert job data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "job_nationality" {
						u := s.BigQueryClient.Dataset("hcm").Table("job_nationality").Inserter()
						item := ParseValueToJobNationality(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert job data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "degree_level" {
						u := s.BigQueryClient.Dataset("hcm").Table("degree_level").Inserter()
						item := ParseValueToDegreeLevel(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert job data")
							s.Logger.WithError(err).Error(e)
						}
					}

				}

			}

			if message.ServerHeartbeat != nil {
				s.Logger.Info("Received server heartbeat")
				// set the flushed LSN (and other LSN values) in the standby status and send to PG
				s.Logger.Info(message.ServerHeartbeat)
				// send Standby Status if the server is requesting for a reply
				if message.ServerHeartbeat.ReplyRequested == 1 {
					s.Logger.Info("Status requested")
					err = config.SendStandbyStatus(session)
					if err != nil {
						s.Logger.WithError(err).Error("Unable to send standby status")
					}
				}
			}

		}

		select {
		case <-wsErr: // ws closed
			s.Logger.Warn("Cancelling context.")
			// cancel session context
			session.CancelFunc()

			err = session.ReplConn.Close()
			if err != nil {
				s.Logger.WithError(err).Error("Could not close replication connection")
			}

		}
	}()
}

// Cancel the currently running session
// Recreate replication connection
func resetSession(session *types.Session) error {
	var err error
	// cancel the currently running session
	if session.CancelFunc != nil {
		session.CancelFunc()
	}

	// close websocket connection
	if session.WSConn != nil {
		//err = session.WSConn.Close()
		if err != nil {
			return err
		}
	}

	// create new context
	ctx, cancelFunc := context.WithCancel(context.Background())
	session.Ctx = ctx
	session.CancelFunc = cancelFunc

	// create the replication connection
	err = config.CheckAndCreateReplConn(session)
	if err != nil {
		return err
	}

	return nil

}
