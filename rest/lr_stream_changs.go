package rest

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/config"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

func lrStream(session *types.Session) error {
	logger := logrus.New()
	// reset the connections
	err := resetSession(session)
	if err != nil {
		logger.WithError(err).Error("Could not create replication connection")
		return fmt.Errorf("Could not create replication connection")
	}

	wsErr := make(chan error, 1)
	go config.LRListenAck(session, wsErr) // concurrently listen on the ws for ack messages
	go config.LRStream(session)           // listen for WAL messages and send them over ws

	select {
	/*case <-c.Writer.CloseNotify(): // ws closed // ?this doesn't work?
	  log.Warn("Websocket connection closed. Cancelling context.")
	  cancelFunc()
	*/
	case <-wsErr: // ws closed
		logger.Warn("Websocket connection closed. Cancelling context.")
		// cancel session context
		session.CancelFunc()
		// close connections
		err = session.WSConn.Close()
		if err != nil {
			logger.WithError(err).Error("Could not close websocket connection")
		}

		err = session.ReplConn.Close()
		if err != nil {
			logger.WithError(err).Error("Could not close replication connection")
		}

	}

	return nil
}

// validate fields in the snapshot data request JSON
func validateSnapshotDataJSON(requestData *types.SnapshotDataJSON) error {
	ob := requestData.OrderBy
	if ob != nil {
		if ob.Column == "" {
			return fmt.Errorf("required field 'column' missing in 'order_by'")
		}
		if !(strings.EqualFold(ob.Order, "asc") || strings.EqualFold(ob.Order, "desc")) {
			return fmt.Errorf("order_by order direction can only be either 'ASC' or 'DESC'")
		}
	}
	return nil
}

func snapshotData(session *types.Session, requestData *types.SnapshotDataJSON) ([]map[string]interface{}, error) {
	return config.SnapshotData(session, requestData)
}
