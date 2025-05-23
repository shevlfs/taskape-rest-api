/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
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

// Optional parameters for the method 'FetchAvailablePhoneNumberCountry'
type FetchAvailablePhoneNumberCountryParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchAvailablePhoneNumberCountryParams) SetPathAccountSid(PathAccountSid string) *FetchAvailablePhoneNumberCountryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

func (c *ApiService) FetchAvailablePhoneNumberCountry(CountryCode string, params *FetchAvailablePhoneNumberCountryParams) (*ApiV2010AvailablePhoneNumberCountry, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers/{CountryCode}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"CountryCode"+"}", CountryCode, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010AvailablePhoneNumberCountry{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAvailablePhoneNumberCountry'
type ListAvailablePhoneNumberCountryParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) requesting the available phone number Country resources.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListAvailablePhoneNumberCountryParams) SetPathAccountSid(PathAccountSid string) *ListAvailablePhoneNumberCountryParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListAvailablePhoneNumberCountryParams) SetPageSize(PageSize int) *ListAvailablePhoneNumberCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListAvailablePhoneNumberCountryParams) SetLimit(Limit int) *ListAvailablePhoneNumberCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of AvailablePhoneNumberCountry records from the API. Request is executed immediately.
func (c *ApiService) PageAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams, pageToken, pageNumber string) (*ListAvailablePhoneNumberCountryResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/AvailablePhoneNumbers.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListAvailablePhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists AvailablePhoneNumberCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams) ([]ApiV2010AvailablePhoneNumberCountry, error) {
	response, errors := c.StreamAvailablePhoneNumberCountry(params)

	records := make([]ApiV2010AvailablePhoneNumberCountry, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams AvailablePhoneNumberCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamAvailablePhoneNumberCountry(params *ListAvailablePhoneNumberCountryParams) (chan ApiV2010AvailablePhoneNumberCountry, chan error) {
	if params == nil {
		params = &ListAvailablePhoneNumberCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010AvailablePhoneNumberCountry, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageAvailablePhoneNumberCountry(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamAvailablePhoneNumberCountry(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamAvailablePhoneNumberCountry(response *ListAvailablePhoneNumberCountryResponse, params *ListAvailablePhoneNumberCountryParams, recordChannel chan ApiV2010AvailablePhoneNumberCountry, errorChannel chan error) {
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

		record, err := client.GetNext(c.baseURL, response, c.getNextListAvailablePhoneNumberCountryResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListAvailablePhoneNumberCountryResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListAvailablePhoneNumberCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAvailablePhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
