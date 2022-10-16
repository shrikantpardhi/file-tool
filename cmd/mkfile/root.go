package mkfile

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra-cli/cmd"
)

var rootCmd =  &cobra.Command{
	Use: "mkfile"
	Short: "mkfile - a simple CLI tool to create files"
	Long: `mkfile is super simple CLI tool to create multiple files at a time.`
	Run: func (cmd *cobra.Commandm args []string)  {
		
	},
}

func Execute () {
	if err: rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while creating a file '%s'", err)
		os.Exit(1)
		
	}
}