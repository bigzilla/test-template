package cmd

import (
	"os"

	"github.com/injustease/test-template/repository/postgres"
	"github.com/injustease/test-template/server/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgDB postgres.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-template",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := postgres.Open(cfgDB)
		if err != nil {
			log.Fatalf("postgres: %v", err)
		}
		defer db.Close()
		log.Info("connected to postgres database")

		port := viper.GetString("PORT")
		srv := http.NewServer(port)

		log.Infof("test http server running on port %s", port)
		log.Fatal(srv.ListenAndServe())
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	cfgDB = postgres.Config{
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetInt("POSTGRES_PORT"),
		User:     viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		DB:       viper.GetString("POSTGRES_DB"),
	}
}
