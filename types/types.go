package types

import (
	"context"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/pubsub"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type MongoConn struct {
	MongoClt  *mongo.Client
	MongoAddr string
}

type PubSubCon struct {
	TopicName      string
	ProjectID      string
	ServiceAccount string

	Topic  *pubsub.Topic
	Client *pubsub.Client
}

// Session stores the context, active db and ws connections, and replication slot state
type Session struct {
	Ctx        context.Context
	CancelFunc context.CancelFunc

	ReplConn *pgx.ReplicationConn
	PGConn   *pgx.Conn

	WSConn *websocket.Conn

	PSConn *PubSubCon

	MbdConn *MongoConn

	BigQueryClient *bigquery.Client

	SlotName     string
	SnapshotName string
	RestartLSN   uint64 //The pg_lsn data type can be used to store LSN (Log Sequence Number) data which is a pointer to a location in the XLOG. This type is a representation of XLogRecPtr and an internal system type of PostgreSQL.
}

// SnapshotDataJSON is the struct that binds with an incoming request for snapshot data
type SnapshotDataJSON struct {
	// SlotName is the name of the replication slot for which the snapshot data needs to be fetched
	// (not used as of now, will be useful in multi client setup)
	SlotName string `json:"slotName" binding:"omitempty"`

	Table   string   `json:"table" binding:"required"`
	Offset  *uint    `json:"offset" binding:"exists"`
	Limit   *uint    `json:"limit" binding:"exists"`
	OrderBy *OrderBy `json:"order_by" binding:"exists"`
}

type OrderBy struct {
	Column string `json:"column" binding:"exists"`
	Order  string `json:"order" binding:"exists"`
	// Nulls TODO
}

type Wal2JSONEvent struct {
	NextLSN string `json:"nextlsn"`
	Change  []map[string]interface{}
}
