package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/cmd/hcm"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/cmd/profile"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/cmd/rest"
	"gitlab.com/target-smart-data-ai-search/pg-log-extractor/version"
)

var (
	verbose bool
	rootCmd *cobra.Command
)

func init() {
	rootCmd = &cobra.Command{
		Use:   "log-extractor",
		Short: "utilities and services",
		Long:  "Top level command for utilities and services of the pg-log-extractor app",
	}

	rootCmd.AddCommand(
		versionCmd,
		rest.Cmd,
		hcm.Cmd,
		profile.Cmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version and exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"tag-connector-be Version: %s \n API Version: %s \n Go Version: %s \n Go OS/ARCH: %s %s",
			version.Version,
			version.APIVersion,
			runtime.Version(),
			runtime.GOOS,
			runtime.GOARCH,
		)
	},
}
