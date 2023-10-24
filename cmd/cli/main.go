package main

import (
	"os"
	"smartklink/pkg/klink"

	_ "github.com/joho/godotenv/autoload"
)

const AppId = "435dfff8ab1f4ec89c41a48d5b2e91d5"
const AppSecret = "e21dcf86ae8d6cdf1246dcb57356ef8c"

func main() {
	klinkClient := klink.NewKlinkClient(
		"435dfff8ab1f4ec89c41a48d5b2e91d5",
		"e21dcf86ae8d6cdf1246dcb57356ef8c",
		"8f194778-7294-11ee-b962-0242ac120002",
	)

	klinkClient.Login(os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
}
