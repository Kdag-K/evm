package run

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/Kdag-K/evm/src/consensus/solo"
	"github.com/Kdag-K/evm/src/engine"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var genesisTemplate = `
{
	"alloc": {
			"{{.}}": {
					"balance": "1337000000000000000000"
			}
	}
}
`

var genesisAddress string

//AddSoloFlags adds flags to the Solo command
func AddSoloFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&genesisAddress, "genesisaddress", "", "create genesis file specifying pre-funded account with given address")
	viper.BindPFlags(cmd.Flags())
}

//NewSoloCmd returns the command that starts EVM-Lite with Solo consensus
func NewSoloCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solo",
		Short: "Run the evm-lite node with Solo consensus (no consensus)",
		PreRunE: func(cmd *cobra.Command, args []string) error {

			config.SetDataDir(config.DataDir)

			logger.WithFields(logrus.Fields{
				"Eth":            config,
				"genesisAddress": genesisAddress,
			}).Debug("Config")

			if cmd.Flags().Changed("genesisaddress") {
				logger.Debug("Writing genesis file")
				if err := createGenesis(config.Genesis, genesisAddress); err != nil {
					return err
				}
			}

			return nil
		},
		RunE: runSolo,
	}
	AddSoloFlags(cmd)
	return cmd
}


