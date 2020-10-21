package run

import (
	_config "github.com/Kdag-K/evm/src/config"
	"github.com/Kdag-K/evm/src/version"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	config = _config.DefaultConfig()
	logger = logrus.New()
)

//RunCmd is launches a node
var RunCmd = &cobra.Command{
	Use:              "run",
	Short:            "Run a node",
	TraverseChildren: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if err := bindFlagsLoadViper(cmd); err != nil {
			return err
		}

		config, err = parseConfig()
		if err != nil {
			return err
		}

		logger = logrus.New()
		logger.Level = logLevel(config.LogLevel)

		logger.WithField("Version", version.Version).Info("Run")

		config.SetDataDir(config.DataDir)

		logger.WithFields(logrus.Fields{
			"Base": config}).Debug("Config")

		return nil
	},
}

func init() {
	//Subcommands
	RunCmd.AddCommand(
		NewSoloCmd())

	//Base config
	RunCmd.PersistentFlags().StringP("datadir", "d", config.DataDir, "Top-level directory for configuration and data")
	RunCmd.PersistentFlags().String("log", config.LogLevel, "debug, info, warn, error, fatal, panic")

	//Eth config
	RunCmd.PersistentFlags().String("eth.genesis", config.Genesis, "Location of genesis file")
	RunCmd.PersistentFlags().String("eth.db", config.DbFile, "Eth database file")
	RunCmd.PersistentFlags().String("eth.listen", config.EthAPIAddr, "Address of HTTP API service")
	RunCmd.PersistentFlags().Int("eth.cache", config.Cache, "Megabytes of memory allocated to internal caching (min 16MB / database forced)")

}

//------------------------------------------------------------------------------

func logLevel(l string) logrus.Level {
	switch l {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}
