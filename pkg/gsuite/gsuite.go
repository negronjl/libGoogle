package gsuite

import (
	"encoding/json"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/api/admin/directory/v1"
	"net/http"
)

func GetBackupCodes(httpClient *http.Client, userEmail string) ([]string, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return nil, errors.New("Empty userEmail ")
	}
	service, err := admin.New(httpClient)
	if err != nil {
		return nil, err
	}
	response, err := service.VerificationCodes.List(userEmail).Do()
	if err != nil {
		return nil, err
	}
	verCodesJSON, err := response.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var verCodes admin.VerificationCodes
	err = json.Unmarshal(verCodesJSON, &verCodes)
	if err != nil {
		return nil, err
	}
	var verCodeArray []string
	for _, value := range verCodes.Items {
		if len(value.VerificationCode) > 0 {
			verCodeArray = append(verCodeArray, value.VerificationCode)
		}
	}
	return verCodeArray, nil
}

func GetGsuiteUsers(httpClient *http.Client, query string) ([]*admin.User, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if query == "" {
		return nil, errors.New("Empty query ")
	}
	service, err := admin.New(httpClient)
	if err != nil {
		return nil, err
	}
	var allUsers []*admin.User
	_ = service.Users.List().
		Projection("full").
		Customer("my_customer").
		MaxResults(500).
		OrderBy("email").
		Query(query).
		Pages(context.Background(),
			func(users *admin.Users) error {
				allUsers = append(allUsers, users.Users...)
				return nil
			})
	if allUsers == nil {
		return nil, errors.New("Received no data ")
	}
	return allUsers, nil
}

func GetGsuiteUser(httpClient *http.Client, userEmail string) (*admin.User, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return nil, errors.New("Empty userEmail ")
	}
	service, err := admin.New(httpClient)
	if err != nil {
		return nil, err
	}
	response, err := service.Users.Get(userEmail).Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received empty response ")
	}
	return response, nil
}

func GetGsuiteGroups(httpClient *http.Client, query string) ([]*admin.Group, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if query == "" {
		return nil, errors.New("Empty query ")
	}
	service, err := admin.New(httpClient)
	if err != nil {
		return nil, err
	}
	var allGroups []*admin.Group
	_ = service.Groups.List().
		Customer("my_customer").
		MaxResults(500).
		OrderBy("email").
		Query(query).
		Pages(context.Background(),
			func(groups *admin.Groups) error {
				allGroups = append(allGroups, groups.Groups...)
				return nil
			})
	if allGroups == nil {
		return nil, errors.New("Received no data ")
	}
	return allGroups, nil
}

func GetGsuiteGroup(httpClient *http.Client, groupEmail string) (*admin.Group, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if groupEmail == "" {
		return nil, errors.New("Empty groupEmail ")
	}
	service, err := admin.New(httpClient)
	if err != nil {
		return nil, err
	}
	response, err := service.Groups.Get(groupEmail).Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received an empty response ")
	}
	return response, nil
}

func UpdateGsuiteUser(httpClient *http.Client, userEmail string, userUpdate *admin.User) (*admin.User, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return nil, errors.New("Empty userEmail ")
	}
	if userUpdate == nil {
		return nil, errors.New("Empty userUpdate ")
	}
	service, err := admin.New(httpClient)
	if err != nil {
		return nil, err
	}
	response, err := service.Users.Update(userEmail, userUpdate).Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received an empty response ")
	}
	return response, nil
}