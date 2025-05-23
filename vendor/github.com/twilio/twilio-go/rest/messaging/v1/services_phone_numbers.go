/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// Optional parameters for the method 'CreatePhoneNumber'
type CreatePhoneNumberParams struct {
	// The SID of the Phone Number being added to the Service.
	PhoneNumberSid *string `json:"PhoneNumberSid,omitempty"`
}

func (params *CreatePhoneNumberParams) SetPhoneNumberSid(PhoneNumberSid string) *CreatePhoneNumberParams {
	params.PhoneNumberSid = &PhoneNumberSid
	return params
}

func (c *ApiService) CreatePhoneNumber(ServiceSid string, params *CreatePhoneNumberParams) (*MessagingV1PhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.PhoneNumberSid != nil {
		data.Set("PhoneNumberSid", *params.PhoneNumberSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeletePhoneNumber(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchPhoneNumber(ServiceSid string, Sid string) (*MessagingV1PhoneNumber, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPhoneNumber'
type ListPhoneNumberParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPhoneNumberParams) SetPageSize(PageSize int) *ListPhoneNumberParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPhoneNumberParams) SetLimit(Limit int) *ListPhoneNumberParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of PhoneNumber records from the API. Request is executed immediately.
func (c *ApiService) PagePhoneNumber(ServiceSid string, params *ListPhoneNumberParams, pageToken, pageNumber string) (*ListPhoneNumberResponse, error) {
	path := "/v1/Services/{ServiceSid}/PhoneNumbers"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists PhoneNumber records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPhoneNumber(ServiceSid string, params *ListPhoneNumberParams) ([]MessagingV1PhoneNumber, error) {
	response, errors := c.StreamPhoneNumber(ServiceSid, params)

	records := make([]MessagingV1PhoneNumber, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams PhoneNumber records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPhoneNumber(ServiceSid string, params *ListPhoneNumberParams) (chan MessagingV1PhoneNumber, chan error) {
	if params == nil {
		params = &ListPhoneNumberParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1PhoneNumber, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePhoneNumber(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPhoneNumber(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamPhoneNumber(response *ListPhoneNumberResponse, params *ListPhoneNumberParams, recordChannel chan MessagingV1PhoneNumber, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.PhoneNumbers
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPhoneNumberResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPhoneNumberResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPhoneNumberResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
