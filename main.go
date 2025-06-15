package main

import "network/cmd"

func main() {
	cmd.Execute()

	/*
		? $./network --help
		Here’s a curated list of free, no-authentication-required APIs specifically useful for developers, covering code execution, utilities, testing, and more

		Usage:
		  network [flags]
		  network [command]

		Available Commands:
		  completion  Generate the autocompletion script for the specified shell
		  help        Help about any command
		  ip          Information of IP
		  joke        Random joke

		Flags:
			  --about   Developer info
		  -h, --help    help for network

		Use "network [command] --help" for more information about a command.
	*/

	/*
		? $./network ip -h
		IP address and location details

		Usage:
		  network ip [flags]

		Flags:
		  -a, --address string   More info about IP
		  -h, --help             help for ip


		? $./network ip
		ip : 103.199.152.60


		? $./network ip -a 103.199.152.60
		region : C
		regionName : Dhaka Division
		city : Kāfrul
		zip : 1207
		query : 103.199.152.60
		status : success
		countryCode : BD
		lat : 23.7882
		lon : 90.3736
		timezone : Asia/Dhaka
		isp : Md. Nasir Uddin T/A SPEED TECH ONLINE
		org :
		as : AS134749 SPEED TECH ONLINE
		country : Bangladesh
	*/

	/*
		? $./network joke -h
		Random joke free

		Usage:
		  network joke [flags]

		Flags:
		  -h, --help   help for joke


		? $./network joke
		punchline : It needed a root canal.
		id : 335
		type : general
		setup : Why did the tree go to the dentist?
	*/

	/*
		? $./network --about
		Develop By:
		Mahadi Hassan MITHUN
		mithun.2121@yahoo.com
		https://github.com/mithun1st
	*/
}
