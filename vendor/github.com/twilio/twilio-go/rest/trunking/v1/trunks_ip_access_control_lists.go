/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Trunking
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

// Optional parameters for the method 'CreateIpAccessControlList'
type CreateIpAccessControlListParams struct {
	// The SID of the [IP Access Control List](https://www.twilio.com/docs/voice/sip/api/sip-ipaccesscontrollist-resource) that you want to associate with the trunk.
	IpAccessControlListSid *string `json:"IpAccessControlListSid,omitempty"`
}

func (params *CreateIpAccessControlListParams) SetIpAccessControlListSid(IpAccessControlListSid string) *CreateIpAccessControlListParams {
	params.IpAccessControlListSid = &IpAccessControlListSid
	return params
}

// Associate an IP Access Control List with a Trunk
func (c *ApiService) CreateIpAccessControlList(TrunkSid string, params *CreateIpAccessControlListParams) (*TrunkingV1IpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.IpAccessControlListSid != nil {
		data.Set("IpAccessControlListSid", *params.IpAccessControlListSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrunkingV1IpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an associated IP Access Control List from a Trunk
func (c *ApiService) DeleteIpAccessControlList(TrunkSid string, Sid string) error {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
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

func (c *ApiService) FetchIpAccessControlList(TrunkSid string, Sid string) (*TrunkingV1IpAccessControlList, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists/{Sid}"
	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)
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

	ps := &TrunkingV1IpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIpAccessControlList'
type ListIpAccessControlListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIpAccessControlListParams) SetPageSize(PageSize int) *ListIpAccessControlListParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIpAccessControlListParams) SetLimit(Limit int) *ListIpAccessControlListParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IpAccessControlList records from the API. Request is executed immediately.
func (c *ApiService) PageIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams, pageToken, pageNumber string) (*ListIpAccessControlListResponse, error) {
	path := "/v1/Trunks/{TrunkSid}/IpAccessControlLists"

	path = strings.Replace(path, "{"+"TrunkSid"+"}", TrunkSid, -1)

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

	ps := &ListIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IpAccessControlList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams) ([]TrunkingV1IpAccessControlList, error) {
	response, errors := c.StreamIpAccessControlList(TrunkSid, params)

	records := make([]TrunkingV1IpAccessControlList, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams IpAccessControlList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIpAccessControlList(TrunkSid string, params *ListIpAccessControlListParams) (chan TrunkingV1IpAccessControlList, chan error) {
	if params == nil {
		params = &ListIpAccessControlListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrunkingV1IpAccessControlList, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageIpAccessControlList(TrunkSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamIpAccessControlList(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamIpAccessControlList(response *ListIpAccessControlListResponse, params *ListIpAccessControlListParams, recordChannel chan TrunkingV1IpAccessControlList, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.IpAccessControlLists
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListIpAccessControlListResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListIpAccessControlListResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListIpAccessControlListResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
