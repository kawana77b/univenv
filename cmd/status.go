package cmd

import (
	"fmt"

	"github.com/kawana77b/univenv/internal/config"
	"github.com/kawana77b/univenv/internal/sysutil"
	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display information about using this tool",
	Long:  `Display information about using this tool`,
	RunE:  runStatus,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func runStatus(cmd *cobra.Command, args []string) error {
	// ENV_UNIVENV_CONFIG_DIR
	opener := &config.ConfigOpener{}
	fmt.Printf("%s:\t%s\n", config.ENV_UNIVENV_CONFIG_DIR, opener.Dir())

	// config exists
	configExists := func() string {
		if sysutil.FileExists(opener.FilePath()) {
			return "TRUE"
		}
		return "FALSE"
	}
	fmt.Printf("%s EXISTS:\t%s\n", opener.FilePath(), configExists())

	// OS
	var o ostype.OS
	o, ok := ostype.GetOS()
	if !ok {
		o = "unknown"
	}
	fmt.Printf("OS:\t%s\n", o.String())

	return nil
}
