package config

import (
	"github.com/jackc/pgx"
	"github.com/prometheus/common/log"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

func InitSession(dbName, pgUser, pgPass, pgHost string, pgPort int) (*types.Session, error) {
	// Initialize the database configuration

	session := types.Session{}
	bdConfig := pgx.ConnConfig{}

	bdConfig.Database = dbName
	bdConfig.Host = pgHost
	bdConfig.Port = uint16(pgPort)
	bdConfig.User = pgUser
	bdConfig.Password = pgPass

	// - creates a db connection
	// create a regular pg connection for use by transactions
	log.Info("Creating regular connection to db")
	pgConn, err := pgx.Connect(bdConfig)
	if err != nil {
		return nil, err
	}

	session.PGConn = pgConn

	// - creates a replication connection
	replConn, err := pgx.ReplicationConnect(BDConfig)
	if err != nil {
		return nil, err
	}

	session.ReplConn = replConn

	// - creates a new replication slot
	slotName := generateSlotName()
	session.SlotName = slotName

	log.Info("Creating replication slot ", slotName)
	consistentPoint, snapshotName, err := session.ReplConn.CreateReplicationSlotEx(slotName, "wal2json")
	if err != nil {
		return nil, err
	}

	log.Infof("Created replication slot \"%s\" with consistent point LSN = %s, snapshot name = %s",
		slotName, consistentPoint, snapshotName)

	lsn, _ := pgx.ParseLSN(consistentPoint)

	session.RestartLSN = lsn
	session.SnapshotName = snapshotName
	return &session, nil
}
