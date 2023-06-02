// package main

// import twilio "github.com/twilio/twilio-go"
// import openapi "github.com/twilio/twilio-go/rest/api/v2010"
// import "os"
// import "fmt"

// func main() {
//     client := twilio.NewRestClient()

//     params := &openapi.CreateMessageParams{}
//     params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
//     params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
//     params.SetBody("Hello from Golang!")

//     _, err := client.Api.CreateMessage(params)
//     if err != nil {
//         fmt.Println(err.Error())
//     } else {
//         fmt.Println("SMS sent successfully!")
//     }
// }

// Download the helper library from https://www.twilio.com/docs/go/install

// package main

// import (
// 	"fmt"
// 	"github.com/twilio/twilio-go"
// 	api "github.com/twilio/twilio-go/rest/api/v2010"
// )

// func main() {
// 	// Find your Account SID and Auth Token at twilio.com/console
// 	// and set the environment variables. See http://twil.io/secure
// 	client := twilio.NewRestClient()

// 	params := &api.CreateMessageParams{}
// 	params.SetBody("Hi there")
// 	params.SetFrom("+13613094103")
// 	params.SetTo("+916353749808")

// 	resp, err := client.Api.CreateMessage(params)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		if resp.Sid != nil {
// 			fmt.Println(*resp.Sid)
// 		} else {
// 			fmt.Println(resp.Sid)
// 		}
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	accountSid := "AC6dd335aa2fb21752156a1d6a1bde43ba"
	authToken := "312266203da5eaa2b5eec4f3c53d88ad"
	//authToken:= os.Getenv("auth_token")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetFrom("+13613094103")
	params.SetTo("+916353749808")
	params.SetBody("Hello from Go its me here!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
