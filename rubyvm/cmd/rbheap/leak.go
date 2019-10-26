package cmd

import (
	"fmt"

	"github.com/jimeh/rbheap/leak"
	"github.com/spf13/cobra"
)

var leakOpts = struct {
	Format  string
	Verbose bool
}{}

// leakCmd represents the leak command
var leakCmd = &cobra.Command{
	Use:   "leak [flags] <dump-A> <dump-B> <dump-C>",
	Short: "Find objects which are likely leaked memory.",
	Long: `Find objects which are likely leaked memory.

Compares the objects in three different dumps (A, B, C), to identify which
objects are present in both B and C, and not present in A.`,

	// Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			usage_er(cmd, fmt.Sprintf("requires 3 args, received %d", len(args)))
		}

		finder := leak.NewFinder(args[0], args[1], args[2])
		finder.Verbose = leakOpts.Verbose

		err := finder.Process()
		if err != nil {
			er(err)
		}

		switch leakOpts.Format {
		case "hex":
			finder.PrintLeakedAddresses()
		case "json":
			err := finder.PrintLeakedObjects()
			if err != nil {
				er(err)
			}
		default:
			usage_er(
				cmd,
				fmt.Sprintf("\"%s\" is not a valid format", leakOpts.Format),
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(leakCmd)

	leakCmd.PersistentFlags().StringVarP(
		&leakOpts.Format,
		"format", "f", "hex",
		"output format: \"hex\" / \"json\"",
	)

	leakCmd.PersistentFlags().BoolVarP(
		&leakOpts.Verbose,
		"verbose", "v", false,
		"print verbose information",
	)
}
