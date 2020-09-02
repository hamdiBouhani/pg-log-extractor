package profile

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/profile"
)

var Cmd = &cobra.Command{
	Use:   "profile-extractor",
	Short: "Start serving a hcm extractor web server",
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

	projectID      string
	serviceAccount string
)

func init() {
	Cmd.Flags().StringVar(&dbName, "db", "", "Name of the database to connect to")
	Cmd.Flags().StringVar(&pgUser, "user", "postgres", "Postgres user name")
	Cmd.Flags().StringVar(&pgPass, "password", "postgres", "Postgres password")
	Cmd.Flags().StringVar(&pgHost, "pgHost", "localhost", "Postgres server hostname")
	Cmd.Flags().IntVar(&pgPort, "pgPort", 5432, "Postgres server port")

	Cmd.Flags().StringVar(&serverHost, "serverHost", "localhost", "Host to listen on")
	Cmd.Flags().StringVar(&serverPort, "serverPort", "8081", "Port to listen on")

	Cmd.Flags().StringVar(&projectID, "project-id", "", "Sets your Google Cloud Platform project ID.")
	Cmd.Flags().StringVarP(&serviceAccount, "service-account", "", "", "Path to Service Account .json file")

}

func run(cmd *cobra.Command, args []string) {
	// Create a logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Create server
	s, err := profile.NewServer(logger, dbName, pgUser, pgPass, pgHost, pgPort, projectID, serviceAccount)
	if err != nil {
		logger.Fatalln("couldn't create REST server:", err)
	}
	logger.Infof("Starting server for database %s; serving at %s:%d", dbName, serverHost, serverPort)
	if err = s.Run(serverHost + ":" + serverPort); err != nil {
		logger.Fatalln("server error:", err)
	}

}
