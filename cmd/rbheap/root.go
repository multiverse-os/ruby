package cmd

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// BuildInfo represents info collected as build-time.
type BuildInfo struct {
	Version string
	Commit  string
	Date    string
}

var rootCmd = &cobra.Command{
	Use:           "rbheap",
	Short:         "rbheap analyzes ObjectSpace dumps from Ruby processes.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func versionString(info *BuildInfo) string {
	var buffer bytes.Buffer
	var meta []string

	buffer.WriteString(info.Version)

	if info.Commit != "unknown" {
		meta = append(meta, info.Commit)
	}

	meta = append(meta, runtime.Version())

	if info.Date != "unknown" {
		meta = append(meta, info.Date)
	}

	if len(meta) > 0 {
		buffer.WriteString(fmt.Sprintf(" (%s)", strings.Join(meta, ", ")))
	}

	return buffer.String()
}

func Execute(info *BuildInfo) {
	rootCmd.Version = versionString(info)
	rootCmd.SetVersionTemplate("{{.Use}} {{.Version}}\n")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Show version.")
}
