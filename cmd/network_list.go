package cmd

import (
	//"fmt"
	//"os"

	"github.com/spf13/cobra"
	//"golang.org/x/exp/slog"

	// "../cgroups"
	// "../container"
	"hind/network"
)

/*
type network_createOptions struct {
	Subnet        string
	Driver        string
	Name        string
}
*/

func network_listCommand() *cobra.Command {
	//opts := network_createOptions{}

	var cmd = &cobra.Command{
		Use:   "network_list",
		Short: "List all container network",
		//Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// runBoot()
			// runRun(opts)
			network.Init()
			network.ListNetwork()
		},
	}

	/*
	flags := cmd.Flags()

	flags.StringVar(&opts.Subnet, "subnet", "", "Subnet,eg: 192.168.0.0/24 ")
	flags.StringVar(&opts.Driver, "driver", "", "Driver,eg: bridge")
	flags.StringVar(&opts.Name, "name", "", "Name,eg: testbridgenet")

	// flags.SetInterspersed(false)
	*/

	return cmd
}

func init() {
	rootCmd.AddCommand(network_listCommand())
}
