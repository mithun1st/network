package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func about() {
	var about string = `
Develop By:
Mahadi Hassan MITHUN
mithun.2121@yahoo.com
https://github.com/mithun1st
	`
	fmt.Println(about)
}

var rootCmd = &cobra.Command{
	Use:   "network",
	Short: "Free API service",
	Long:  `Here’s a curated list of free, no-authentication-required APIs specifically useful for developers, covering code execution, utilities, testing, and more`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		hasAbout, _ := cmd.Flags().GetBool("about")
		if hasAbout {
			about()
		}
	},
}

func init() {
	rootCmd.Flags().Bool("about", false, "Developer info")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
