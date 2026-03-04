package cmd

import (
	"os"
	"strings"

	"github.com/rkrebs/sonar/internal/display"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sonar",
	Short: "Detect services listening on localhost ports",
	RunE: func(cmd *cobra.Command, args []string) error {
		return listRun(cmd, args)
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable colored output")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if nc, _ := cmd.Flags().GetBool("no-color"); nc {
			display.NoColor = true
		}
	}

	// Register list flags on root too since `sonar` delegates to listRun
	rootCmd.Flags().BoolVar(&jsonFlag, "json", false, "Output as JSON")
	rootCmd.Flags().StringVar(&filterFlag, "filter", "", "Filter by type: docker, user, system")
	rootCmd.Flags().StringVar(&sortFlag, "sort", "port", "Sort by: port, pid, name, type")
	rootCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "Include desktop apps (hidden by default)")
	rootCmd.Flags().StringVarP(&columnsFlag, "columns", "c", "",
		"Columns to display (comma-separated: "+strings.Join(display.AllColumns, ", ")+")")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
