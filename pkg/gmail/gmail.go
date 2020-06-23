package gmail

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
	"net/http"
)

func ListLabels(httpClient *http.Client, userEmail string) ([]string, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return nil, errors.New("Empty userEmail ")
	}
	service, err := gmail.New(httpClient)
	if err != nil {
		return nil, err
	}
	response, err := service.Users.Labels.List(userEmail).Do()
	if err != nil {
		return nil, err
	}
	var returnValue []string
	for _, value := range response.Labels {
		returnValue = append(returnValue, value.Name)
	}
	return returnValue, nil
}

func ListMessagesMatchingQuery(httpClient *http.Client, userEmail string, query string) ([]gmail.Message, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return nil, errors.New("Empty userEmail ")
	}
	if query == "" {
		return nil, errors.New("Empty query ")
	}
	service, err := gmail.New(httpClient)
	if err != nil {
		return nil, err
	}
	var allMessages []*gmail.Message
	_ = service.Users.Messages.
		List(userEmail).
		Q(query).
		MaxResults(500).
		Pages(context.Background(),
			func(message *gmail.ListMessagesResponse) error {
				allMessages = append(allMessages, message.Messages...)
				return nil
			})
	if allMessages == nil {
		return nil, errors.New(" Received no data")
	}
	var returnValue []gmail.Message
	for _, value := range allMessages {
		returnValue = append(returnValue, *value)
	}
	return returnValue, nil
}

func GetMessageById(httpClient *http.Client, userEmail string, messageId string) (gmail.Message, error) {
	if httpClient == nil {
		return gmail.Message{}, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return gmail.Message{}, errors.New("Empty userEmail ")
	}
	if messageId == "" {
		return gmail.Message{}, errors.New("Empty messageId ")
	}
	service, err := gmail.New(httpClient)
	if err != nil {
		return gmail.Message{}, err
	}
	message, err := service.Users.Messages.Get(userEmail, messageId).Format("raw").Do()
	if err != nil {
		return gmail.Message{}, err
	}
	return *message, nil
}

func ArchiveMessage(httpClient *http.Client, userEmail string, messageId string) (gmail.Message, error) {
	if httpClient == nil {
		return gmail.Message{}, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return gmail.Message{}, errors.New("Empty userEmail ")
	}
	if messageId == "" {
		return gmail.Message{}, errors.New("Empty messageId ")
	}
	service, err := gmail.New(httpClient)
	if err != nil {
		return gmail.Message{}, err
	}

	response, err := service.Users.Messages.Modify(
		userEmail,
		messageId,
		&gmail.ModifyMessageRequest{
			RemoveLabelIds: []string{"INBOX"},
		}).Do()
	if err != nil {
		return gmail.Message{}, err
	}
	return *response, nil
}

func UnArchiveMessage(httpClient *http.Client, userEmail string, messageId string) (gmail.Message, error) {
	if httpClient == nil {
		return gmail.Message{}, errors.New("Empty httpClient ")
	}
	if userEmail == "" {
		return gmail.Message{}, errors.New("Empty userEmail ")
	}
	if messageId == "" {
		return gmail.Message{}, errors.New("Empty messageId ")
	}
	service, err := gmail.New(httpClient)
	if err != nil {
		return gmail.Message{}, err
	}

	response, err := service.Users.Messages.Modify(
		userEmail,
		messageId,
		&gmail.ModifyMessageRequest{
			AddLabelIds: []string{"INBOX"},
		}).Do()
	if err != nil {
		return gmail.Message{}, err
	}
	return *response, nil
}