package gh

import (
	"github.com/cli/go-gh/v2/pkg/api"
)

var client *api.RESTClient

func init() {
	var err error
	client, err = api.DefaultRESTClient()
	if err != nil {
		panic(err)
	}
}

func GetLoginUser() UserResponse {
	response := UserResponse{}
	err := client.Get("user", &response)
	if err != nil {
		panic(err)
	}
	return response
}

func GetUserEvents(username string) []*UserEventResponse {
	var responses []*UserEventResponse
	err := client.Get("users/"+username+"/events", &responses)
	if err != nil {
		panic(err)
	}
	return responses
}
