package cmd

import (
	"fmt"
	"network/service"
	"network/utils"

	"github.com/spf13/cobra"
)

func joke() {
	dataMap, err := service.GetMethod("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(utils.CmtOutput(dataMap))
}

var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Random joke",
	Long:  `Random joke free`,
	Run: func(cmd *cobra.Command, args []string) {
		joke()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}
