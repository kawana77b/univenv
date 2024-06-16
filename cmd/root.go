package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/kawana77b/univenv/internal/builder"
	"github.com/kawana77b/univenv/internal/common/lo"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
	"github.com/spf13/cobra"
)

var Fs embed.FS
var Version string = "0.0.1"

type rootArgs struct {
	shell string
}

type rootOptions struct {
	file      string
	target    string
	noComment bool
}

var rootArgsVal rootArgs
var rootOptsVal rootOptions

var shells = lo.AsStrings(shell.All())

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "univenv",
	Short: "",
	Long: fmt.Sprintf(`univenv is a tool for universal shell platforms that allows you to output your yml as bash, fish or powershell (pwsh).

Available arguments: %s
	`, strings.Join(shells, ", ")),
	ValidArgs: shells,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	PreRunE:   preRunRoot,
	RunE:      runRoot,
}

func Execute() {
	rootCmd.Version = Version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rootOptsVal.file, "file", "f", "", "Set the file to be used.")
	rootCmd.Flags().StringVarP(&rootOptsVal.target, "target", "t", "", "Set the target name. If the given string is dev, it tries to read config.dev.yml. If a file is specified, it is ignored.")
	rootCmd.Flags().BoolVarP(&rootOptsVal.noComment, "no-comment", "n", false, "Suppresses comments inserted before or after the generated script.")
}

func preRunRoot(cmd *cobra.Command, args []string) error {
	rootArgsVal.shell = args[0]
	return nil
}

func runRoot(cmd *cobra.Command, args []string) error {
	b := &builder.ScriptBuilder{}
	b.DetectOS()

	b.SetFS(Fs)
	b.SetShell(rootArgsVal.shell)
	b.SetNoComment(rootOptsVal.noComment)
	b.SetTarget(rootOptsVal.target)
	b.SetFile(rootOptsVal.file)

	script, err := b.Build()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", script)

	return nil
}
