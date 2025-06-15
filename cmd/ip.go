package cmd

import (
	"fmt"
	"network/service"
	"network/utils"

	"github.com/spf13/cobra"
)

func ip() {
	dataMap, err := service.GetMethod("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(utils.CmtOutput(dataMap))
}

func ipAddress(address string) {
	dataMap, err := service.GetMethod("http://ip-api.com/json/" + address)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(utils.CmtOutput(dataMap))
}

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Information of IP",
	Long:  `IP address and location details`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")

		if address != "" {
			ipAddress(address)
		} else {
			ip()

		}
	},
}

func init() {
	rootCmd.AddCommand(ipCmd)

	ipCmd.Flags().StringP("address", "a", "", "More info about IP")
}
