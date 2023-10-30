package main

import (
	"log"
	"os"
	"smartklink/pkg/klink"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	klinkClient := klink.NewKlinkClient(
		os.Getenv("APP_ID"),
		os.Getenv("APP_SECRET"),
		os.Getenv("UNIQUE_ID"),
	)

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	creds, err := klinkClient.Login(username, password)
	if err != nil {
		log.Fatalln("Error login:", err)
	}
	log.Println(creds)

	// token := os.Getenv("TEMP_ACCESS_TOKEN")
	// // info, err := klinkClient.GetUserInfo(token)
	// // if err != nil {
	// // 	log.Fatal("Error fetching user info:", err)
	// // }
	// // log.Println(info)
	// lock := os.Getenv("TEMP_LOCK")

	// room, err := klinkClient.GetRoom(lock, token)
	// if err != nil {
	// 	log.Fatal("Error fetching rooom:", err)
	// }
	// log.Println(room)
}
