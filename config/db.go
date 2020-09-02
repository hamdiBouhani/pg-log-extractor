package config

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgx"
	"github.com/prometheus/common/log"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

var BDConfig = pgx.ConnConfig{}

// Initialize the database configuration
func CreateConfig(dbName, pgUser, pgPass, pgHost string, pgPort int) {
	BDConfig.Database = dbName
	BDConfig.Host = pgHost
	BDConfig.Port = uint16(pgPort)
	BDConfig.User = pgUser
	BDConfig.Password = pgPass
}

// Init function
// - creates a db connection
// - creates a replication connection
// - delete existing replication slots
// - creates a new replication slot
// - gets the consistent point LSN and snapshot name
// - populates the Session object
func Init(session *types.Session) error {

	// create a regular pg connection for use by transactions
	log.Info("Creating regular connection to db")
	pgConn, err := pgx.Connect(BDConfig)
	if err != nil {
		return err
	}

	session.PGConn = pgConn

	log.Info("Creating replication connection to ", BDConfig.Database)
	if session.ReplConn != nil {
		log.Info("Closing existing replication connection")
		session.ReplConn.Close()
	}

	replConn, err := pgx.ReplicationConnect(BDConfig)
	if err != nil {
		return err
	}

	session.ReplConn = replConn

	// delete all existing slots
	err = deleteAllSlots(session)
	if err != nil {
		log.Errorf("could not delete replication slots : %v", err)
	}

	// create new slots
	slotName := generateSlotName()
	session.SlotName = slotName

	log.Info("Creating replication slot ", slotName)
	consistentPoint, snapshotName, err := session.ReplConn.CreateReplicationSlotEx(slotName, "wal2json")
	if err != nil {
		return err
	}

	log.Infof("Created replication slot \"%s\" with consistent point LSN = %s, snapshot name = %s",
		slotName, consistentPoint, snapshotName)

	lsn, _ := pgx.ParseLSN(consistentPoint)

	session.RestartLSN = lsn
	session.SnapshotName = snapshotName

	return nil
}

// CheckAndCreateReplConn creates a new replication connection
func CheckAndCreateReplConn(session *types.Session) error {
	if session.ReplConn != nil {
		if session.ReplConn.IsAlive() {
			// reuse the existing connection (or close it nonetheless?)
			return nil
		}
	}

	replConn, err := pgx.ReplicationConnect(BDConfig)
	if err != nil {
		return err
	}
	session.ReplConn = replConn

	return nil
}

// generates a random slot name which can be remembered
func generateSlotName() string {
	// list of random words
	strs := []string{
		"gigantic",
		"scold",
		"greasy",
		"shaggy",
		"wasteful",
		"few",
		"face",
		"pet",
		"ablaze",
		"mundane",
	}

	rand.Seed(time.Now().Unix())

	// generate name such as delta_gigantic20
	name := fmt.Sprintf("delta_%s%d", strs[rand.Intn(len(strs))], rand.Intn(100))

	return name
}

// delete all old slots that were created by us
func deleteAllSlots(session *types.Session) error {
	rows, err := session.PGConn.Query("SELECT slot_name FROM pg_replication_slots")
	if err != nil {
		return err
	}
	for rows.Next() {
		var slotName string
		rows.Scan(&slotName)

		// only delete slots created by this program
		if !strings.Contains(slotName, "delta_") {
			continue
		}

		log.Infof("Deleting replication slot %s", slotName)
		err = session.ReplConn.DropReplicationSlot(slotName)
		//_,err = session.PGConn.Exec(fmt.Sprintf("SELECT pg_drop_replication_slot(\"%s\")", slotName))
		if err != nil {
			log.With("could not delete slot ", slotName).Error(err)
		}
	}
	return nil
}
