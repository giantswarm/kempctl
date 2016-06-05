package main

import (
	kemp "github.com/giantswarm/kemp-client"
	"github.com/spf13/cobra"
)

var (
	globalFlags struct {
		debug    bool
		verbose  bool
		user     string
		password string
		endpoint string
	}

	KempCtlCmd = &cobra.Command{
		Use:   "kempctl",
		Short: "Manage kemp loadmaster via api",
		Long:  "Manage kemp loadmaster via api",
		Run:   KempCtlRun,
	}

	projectVersion string
	projectBuild   string
)

func init() {
	KempCtlCmd.PersistentFlags().BoolVarP(&globalFlags.debug, "debug", "d", false, "Print debug output")
	KempCtlCmd.PersistentFlags().BoolVarP(&globalFlags.verbose, "verbose", "v", false, "Print verbose output")
	KempCtlCmd.PersistentFlags().StringVar(&globalFlags.user, "user", "", "User name to access the loadbalancer api")
	KempCtlCmd.PersistentFlags().StringVar(&globalFlags.password, "password", "", "Password to access the loadbalancer api")
	KempCtlCmd.PersistentFlags().StringVar(&globalFlags.endpoint, "endpoint", "", "Endpoint of the loadbalancer api (eg. https://1.2.3.4/access/)")
}

func createClient() *kemp.Client {
	return kemp.NewClient(kemp.Config{
		User:     globalFlags.user,
		Password: globalFlags.password,
		Endpoint: globalFlags.endpoint,
		Debug:    globalFlags.debug,
	})
}

func KempCtlRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	KempCtlCmd.AddCommand(versionCmd)
	KempCtlCmd.AddCommand(getCmd)
	KempCtlCmd.AddCommand(setCmd)
	KempCtlCmd.AddCommand(virtualCmd)
	KempCtlCmd.AddCommand(statsCmd)
	virtualCmd.AddCommand(virtualListCmd)
	virtualCmd.AddCommand(virtualAddCmd)
	virtualCmd.AddCommand(virtualDeleteCmd)
	virtualCmd.AddCommand(virtualUpdateCmd)
	virtualCmd.AddCommand(virtualShowCmd)
	KempCtlCmd.AddCommand(realCmd)
	realCmd.AddCommand(realAddCmd)

	KempCtlCmd.Execute()
}
