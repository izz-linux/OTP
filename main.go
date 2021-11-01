package main

import (
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	// using .env file for actual values
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Something bad happened loading environment file::error %s", err)
	}

	// lastpass
	lpsec := os.Getenv("LASTPASS_SECRET")
	showOTP("lastpass", lpsec)

	// dropbox
	dbsec := os.Getenv("DROPBOX_SECRET")
	showOTP("dropbox", dbsec)
}
