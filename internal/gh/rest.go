package gh

import (
	"github.com/cli/go-gh/v2/pkg/api"
)

var restClient *api.RESTClient

func init() {
	var err error
	restClient, err = api.DefaultRESTClient()
	if err != nil {
		panic(err)
	}
}

func GetLoginUser() UserResponse {
	response := UserResponse{}
	err := restClient.Get("user", &response)
	if err != nil {
		panic(err)
	}
	return response
}

type UserResponse struct {
	Login string
}
