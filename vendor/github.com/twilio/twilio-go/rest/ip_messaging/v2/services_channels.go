/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Ip_messaging
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateChannel'
type CreateChannelParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	UniqueName *string `json:"UniqueName,omitempty"`
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	Type *string `json:"Type,omitempty"`
	//
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	//
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	//
	CreatedBy *string `json:"CreatedBy,omitempty"`
}

func (params *CreateChannelParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *CreateChannelParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *CreateChannelParams) SetFriendlyName(FriendlyName string) *CreateChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateChannelParams) SetUniqueName(UniqueName string) *CreateChannelParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateChannelParams) SetAttributes(Attributes string) *CreateChannelParams {
	params.Attributes = &Attributes
	return params
}
func (params *CreateChannelParams) SetType(Type string) *CreateChannelParams {
	params.Type = &Type
	return params
}
func (params *CreateChannelParams) SetDateCreated(DateCreated time.Time) *CreateChannelParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *CreateChannelParams) SetDateUpdated(DateUpdated time.Time) *CreateChannelParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *CreateChannelParams) SetCreatedBy(CreatedBy string) *CreateChannelParams {
	params.CreatedBy = &CreatedBy
	return params
}

func (c *ApiService) CreateChannel(ServiceSid string, params *CreateChannelParams) (*IpMessagingV2Channel, error) {
	path := "/v2/Services/{ServiceSid}/Channels"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.Type != nil {
		data.Set("Type", *params.Type)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.CreatedBy != nil {
		data.Set("CreatedBy", *params.CreatedBy)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteChannel'
type DeleteChannelParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
}

func (params *DeleteChannelParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *DeleteChannelParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}

func (c *ApiService) DeleteChannel(ServiceSid string, Sid string, params *DeleteChannelParams) error {
	path := "/v2/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}
	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchChannel(ServiceSid string, Sid string) (*IpMessagingV2Channel, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{Sid}"
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

	ps := &IpMessagingV2Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannel'
type ListChannelParams struct {
	//
	Type *[]string `json:"Type,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListChannelParams) SetType(Type []string) *ListChannelParams {
	params.Type = &Type
	return params
}
func (params *ListChannelParams) SetPageSize(PageSize int) *ListChannelParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListChannelParams) SetLimit(Limit int) *ListChannelParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Channel records from the API. Request is executed immediately.
func (c *ApiService) PageChannel(ServiceSid string, params *ListChannelParams, pageToken, pageNumber string) (*ListChannelResponse, error) {
	path := "/v2/Services/{ServiceSid}/Channels"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Type != nil {
		for _, item := range *params.Type {
			data.Add("Type", item)
		}
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

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Channel records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListChannel(ServiceSid string, params *ListChannelParams) ([]IpMessagingV2Channel, error) {
	response, errors := c.StreamChannel(ServiceSid, params)

	records := make([]IpMessagingV2Channel, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Channel records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamChannel(ServiceSid string, params *ListChannelParams) (chan IpMessagingV2Channel, chan error) {
	if params == nil {
		params = &ListChannelParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan IpMessagingV2Channel, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageChannel(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamChannel(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamChannel(response *ListChannelResponse, params *ListChannelParams, recordChannel chan IpMessagingV2Channel, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Channels
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListChannelResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListChannelResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListChannelResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateChannel'
type UpdateChannelParams struct {
	// The X-Twilio-Webhook-Enabled HTTP request header
	XTwilioWebhookEnabled *string `json:"X-Twilio-Webhook-Enabled,omitempty"`
	//
	FriendlyName *string `json:"FriendlyName,omitempty"`
	//
	UniqueName *string `json:"UniqueName,omitempty"`
	//
	Attributes *string `json:"Attributes,omitempty"`
	//
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	//
	DateUpdated *time.Time `json:"DateUpdated,omitempty"`
	//
	CreatedBy *string `json:"CreatedBy,omitempty"`
}

func (params *UpdateChannelParams) SetXTwilioWebhookEnabled(XTwilioWebhookEnabled string) *UpdateChannelParams {
	params.XTwilioWebhookEnabled = &XTwilioWebhookEnabled
	return params
}
func (params *UpdateChannelParams) SetFriendlyName(FriendlyName string) *UpdateChannelParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateChannelParams) SetUniqueName(UniqueName string) *UpdateChannelParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *UpdateChannelParams) SetAttributes(Attributes string) *UpdateChannelParams {
	params.Attributes = &Attributes
	return params
}
func (params *UpdateChannelParams) SetDateCreated(DateCreated time.Time) *UpdateChannelParams {
	params.DateCreated = &DateCreated
	return params
}
func (params *UpdateChannelParams) SetDateUpdated(DateUpdated time.Time) *UpdateChannelParams {
	params.DateUpdated = &DateUpdated
	return params
}
func (params *UpdateChannelParams) SetCreatedBy(CreatedBy string) *UpdateChannelParams {
	params.CreatedBy = &CreatedBy
	return params
}

func (c *ApiService) UpdateChannel(ServiceSid string, Sid string, params *UpdateChannelParams) (*IpMessagingV2Channel, error) {
	path := "/v2/Services/{ServiceSid}/Channels/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.Attributes != nil {
		data.Set("Attributes", *params.Attributes)
	}
	if params != nil && params.DateCreated != nil {
		data.Set("DateCreated", fmt.Sprint((*params.DateCreated).Format(time.RFC3339)))
	}
	if params != nil && params.DateUpdated != nil {
		data.Set("DateUpdated", fmt.Sprint((*params.DateUpdated).Format(time.RFC3339)))
	}
	if params != nil && params.CreatedBy != nil {
		data.Set("CreatedBy", *params.CreatedBy)
	}

	if params != nil && params.XTwilioWebhookEnabled != nil {
		headers["X-Twilio-Webhook-Enabled"] = *params.XTwilioWebhookEnabled
	}
	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &IpMessagingV2Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
