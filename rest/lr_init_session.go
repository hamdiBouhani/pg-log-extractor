package rest

import (
	"context"

	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/config"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/types"
)

//*******init********

// create replication slot ex and get snapshot name, consistent point
// return slotname
func initDB(session *types.Session) error {
	var err error
	// initilize the connections for the session
	resetSession(session)
	err = config.Init(session)
	if err != nil {
		return err
	}

	return nil
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
