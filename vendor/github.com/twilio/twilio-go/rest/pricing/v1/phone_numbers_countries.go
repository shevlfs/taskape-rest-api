/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Pricing
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

func (c *ApiService) FetchPhoneNumberCountry(IsoCountry string) (*PricingV1PhoneNumberCountryInstance, error) {
	path := "/v1/PhoneNumbers/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1PhoneNumberCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPhoneNumberCountry'
type ListPhoneNumberCountryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPhoneNumberCountryParams) SetPageSize(PageSize int) *ListPhoneNumberCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPhoneNumberCountryParams) SetLimit(Limit int) *ListPhoneNumberCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of PhoneNumberCountry records from the API. Request is executed immediately.
func (c *ApiService) PagePhoneNumberCountry(params *ListPhoneNumberCountryParams, pageToken, pageNumber string) (*ListPhoneNumberCountryResponse, error) {
	path := "/v1/PhoneNumbers/Countries"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists PhoneNumberCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPhoneNumberCountry(params *ListPhoneNumberCountryParams) ([]PricingV1PhoneNumberCountry, error) {
	response, errors := c.StreamPhoneNumberCountry(params)

	records := make([]PricingV1PhoneNumberCountry, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams PhoneNumberCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPhoneNumberCountry(params *ListPhoneNumberCountryParams) (chan PricingV1PhoneNumberCountry, chan error) {
	if params == nil {
		params = &ListPhoneNumberCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan PricingV1PhoneNumberCountry, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePhoneNumberCountry(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPhoneNumberCountry(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamPhoneNumberCountry(response *ListPhoneNumberCountryResponse, params *ListPhoneNumberCountryParams, recordChannel chan PricingV1PhoneNumberCountry, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Countries
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPhoneNumberCountryResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPhoneNumberCountryResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPhoneNumberCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
