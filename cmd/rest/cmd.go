package rest

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/rest"
)

var Cmd = &cobra.Command{
	Use:   "rest",
	Short: "Start serving a ws server",
	Run:   run,
}

var (
	dbName string
	pgUser string
	pgPass string
	pgHost string
	pgPort int

	serverHost string
	serverPort string

	topic          string
	projectID      string
	serviceAccount string

	mongoAddr  string
	mongoHosts string
	repSetName string
)

func init() {
	//Cmd.Flags().StringVarP(&addr, "address", "a", ":9502", "Address to listen on")
	Cmd.Flags().StringVar(&dbName, "db", "", "Name of the database to connect to")
	Cmd.Flags().StringVar(&pgUser, "user", "postgres", "Postgres user name")
	Cmd.Flags().StringVar(&pgPass, "password", "postgres", "Postgres password")
	Cmd.Flags().StringVar(&pgHost, "pgHost", "localhost", "Postgres server hostname")
	Cmd.Flags().IntVar(&pgPort, "pgPort", 5432, "Postgres server port")

	Cmd.Flags().StringVar(&serverHost, "serverHost", "0.0.0.0", "Host to listen on")
	Cmd.Flags().StringVar(&serverPort, "serverPort", "8080", "Port to listen on")

	Cmd.Flags().StringVar(&topic, "topic", "default_topic", "The Kafka topic to use")
	Cmd.Flags().StringVar(&projectID, "project-id", "", "Sets your Google Cloud Platform project ID.")
	Cmd.Flags().StringVarP(&serviceAccount, "service-account", "", "", "Path to Service Account .json file")

	Cmd.Flags().StringVar(&mongoAddr, "mongo-addr", "", "mongo hostname")
	Cmd.Flags().StringVar(&mongoHosts, "mongo-hosts", "", "mongo hosts in replicat set mode : separted with ,")
	Cmd.Flags().StringVar(&repSetName, "rs-name", "", "mongo replica set name")
}

func run(cmd *cobra.Command, args []string) {
	// Create a logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Create server
	s, err := rest.NewServer(logger, dbName, pgUser, pgPass, pgHost, pgPort, topic, projectID, serviceAccount, mongoAddr, repSetName, mongoHosts)
	if err != nil {
		logger.Fatalln("couldn't create REST server:", err)
	}
	logger.Infof("Starting server for database %s; serving at %s:%d", dbName, serverHost, serverPort)
	if err = s.Run(serverHost + ":" + serverPort); err != nil {
		logger.Fatalln("server error:", err)
	}

}
