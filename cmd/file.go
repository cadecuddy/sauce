package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search using a local file.",
	Long: `File searches support most visual file mediums under 25MB including:
- pdf
- jpeg/jpg
- gif
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("file called")
	},
	Args: cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
