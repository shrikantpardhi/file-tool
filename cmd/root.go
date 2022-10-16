package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mkCmd = &cobra.Command{
	Use:   "mkfile",
	Short: "mkfile is CLI tool to create a file",
	Long: `A Fast and Flexible way to create multiple files at a time.
			created using spf13/cobra and Go.
			Complete documentation is available at https://github.com/shrikantpardhi/file-tool`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, element := range args {
			newfile, err := os.Create(element)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Println("File created: " + element)
			newfile.Close()
		}
	},
}

func Execute() {
	if err := mkCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
