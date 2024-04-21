package main

import (
	"fmt"

	"github.com/chaitsgithub/goprojects/zerodha/services/testservice/constants"
	kiteconnect "github.com/zerodha/gokiteconnect/v4"
)

func main() {
	// Create a new Kite connect instance
	kc := kiteconnect.New(constants.ZERODHA_API_KEY)

	// Login URL from which request token can be obtained
	fmt.Println(kc.GetLoginURL())
	kc.Get

	// Obtained request token after Kite Connect login flow
	fmt.Println(constants.REQUEST_TOKEN)
	// Get user details and access token
	data, err := kc.GenerateSession(constants.REQUEST_TOKEN, constants.ZERODHA_API_SECRET)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Set access token
	kc.SetAccessToken(data.AccessToken)

	// Get margins
	margins, err := kc.GetUserMargins()
	if err != nil {
		fmt.Printf("Error getting margins: %v", err)
	}
	fmt.Println("margins: ", margins)
}
