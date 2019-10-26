package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func usage_er(cmd *cobra.Command, msg interface{}) {
	cmd.Usage()
	fmt.Println("")
	er(msg)
}

func er(msg interface{}) {
	fmt.Println("ERROR:", msg)
	os.Exit(1)
}
