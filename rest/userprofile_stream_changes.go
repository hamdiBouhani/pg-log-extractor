package rest

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

func (s *Server) StreamUserProfileIntoBigQuery(c *gin.Context) {
	s.Logger.Info("Stream User Profile Data Into BigQuery")

	config.CreateConfig(s.DbName, s.PgUser, s.PgPass, s.PgHost, s.PgPort)
	session := &types.Session{}

	s.Logger.Info("initDB")
	err := initDB(session)
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

					if value.Table == "user_profile" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("user_profile").Inserter()
						item := ParseValueToUserProfile(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert user_profile data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "user_career_aspiration" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("user_profile_career_aspiration").Inserter()
						item := ParseValueToUserCareerAspiration(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert user_profile_career_aspiration data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "experience" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("experience").Inserter()
						item := ParseValueToExperience(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert experience data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "user_competency" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("user_competency").Inserter()
						item := ParseValueToUserCompetency(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert user_competency data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "education_specialization" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("education_specialization").Inserter()
						item := ParseValueToEducationSpecialization(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert education_specialization data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "user_language" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("language_bridge").Inserter()
						item := ParseValueToUserLanguage(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert user_language data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "user_education" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("user_education").Inserter()
						item := ParseValueToUserEducation(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert user_education data")
							s.Logger.WithError(err).Error(e)
						}
					}

					if value.Table == "degree_level" {
						u := s.BigQueryClient.Dataset(DatasetID).Table("degree_level").Inserter()
						item := ParseValueToDegreeLevel(value.GetValue())
						if err := u.Put(context.Background(), item); err != nil {
							e := fmt.Sprintf("unable to insert degree_level data")
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
