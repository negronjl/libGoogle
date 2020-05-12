package gdocs

import (
	"errors"
	"google.golang.org/api/sheets/v4"
	"net/http"
)

func CreateGoogleSpreadsheet(httpClient *http.Client, title string) (*sheets.Spreadsheet, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if title == "" {
		return nil, errors.New("Empty title ")
	}
	service, err := sheets.New(httpClient)
	if err != nil {
		return nil, err
	}
	sheet := service.Spreadsheets
	if sheet == nil {
		return nil, errors.New("Unable to instantiate a Spreadsheet service ")
	}
	newSpreadsheet := sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}
	response, err := sheet.Create(&newSpreadsheet).Fields("spreadsheetId").Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received an empty response ")
	}
	return response, nil
}

func AddGoogleSheet(httpClient *http.Client, id string, title string) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if title == "" {
		return nil, errors.New("Empty title ")
	}
	if id == "" {
		return nil, errors.New("Empty id ")
	}
	service, err := sheets.New(httpClient)
	if err != nil {
		return nil, err
	}
	sheet := service.Spreadsheets
	if sheet == nil {
		return nil, errors.New("Unable to instantiate a Spreadsheet service ")
	}
	bacthUpdate := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title: title,
					},
				},
			},
		},
	}
	response, err := sheet.BatchUpdate(id, &bacthUpdate).Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received an empty response ")
	}
	return response, nil
}

func UpdateGoogleSheet(httpClient *http.Client, spreadsheetId string, spreadsheetRange string,
	spreadsheetValues *sheets.ValueRange) (*sheets.UpdateValuesResponse, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if spreadsheetId == "" {
		return nil, errors.New("Empty spreadsheetId ")
	}
	if spreadsheetRange == "" {
		return nil, errors.New("Empty spreadsheetRange ")
	}
	if spreadsheetValues == nil {
		return nil, errors.New("Empty Values ")
	}
	service, err := sheets.New(httpClient)
	if err != nil {
		return nil, err
	}
	sheet := service.Spreadsheets
	if sheet == nil {
		return nil, errors.New("Unable to instantiate a Spreadsheet service ")
	}
	response, err := sheet.Values.Update(spreadsheetId, spreadsheetRange, spreadsheetValues).
		ValueInputOption("RAW").Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received an empty response ")
	}
	return response, nil
}

func GetGoogleSheet(httpClient *http.Client, spreadsheetId string, spreadsheetRange string) (*sheets.ValueRange, error) {
	if httpClient == nil {
		return nil, errors.New("Empty httpClient ")
	}
	if spreadsheetRange == "" {
		return nil, errors.New("Empty range ")
	}
	service, err := sheets.New(httpClient)
	if err != nil {
		return nil, err
	}
	sheet := service.Spreadsheets
	if sheet == nil {
		return nil, errors.New("Unable to instantiate a Spreadsheet service ")
	}
	response, err := sheet.Values.Get(spreadsheetId, spreadsheetRange).Do()
	if err != nil {
		return nil, err
	}
	if response == nil {
		return nil, errors.New("Received an empty response ")
	}
	return response, nil
}
