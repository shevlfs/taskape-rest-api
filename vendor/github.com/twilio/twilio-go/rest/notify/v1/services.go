/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Notify
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

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings.
	ApnCredentialSid *string `json:"ApnCredentialSid,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings.
	GcmCredentialSid *string `json:"GcmCredentialSid,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/quickstart#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// Deprecated.
	FacebookMessengerPageId *string `json:"FacebookMessengerPageId,omitempty"`
	// The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultApnNotificationProtocolVersion *string `json:"DefaultApnNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultGcmNotificationProtocolVersion *string `json:"DefaultGcmNotificationProtocolVersion,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings.
	FcmCredentialSid *string `json:"FcmCredentialSid,omitempty"`
	// The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultFcmNotificationProtocolVersion *string `json:"DefaultFcmNotificationProtocolVersion,omitempty"`
	// Whether to log notifications. Can be: `true` or `false` and the default is `true`.
	LogEnabled *bool `json:"LogEnabled,omitempty"`
	// Deprecated.
	AlexaSkillId *string `json:"AlexaSkillId,omitempty"`
	// Deprecated.
	DefaultAlexaNotificationProtocolVersion *string `json:"DefaultAlexaNotificationProtocolVersion,omitempty"`
	// URL to send delivery status callback.
	DeliveryCallbackUrl *string `json:"DeliveryCallbackUrl,omitempty"`
	// Callback configuration that enables delivery callbacks, default false
	DeliveryCallbackEnabled *bool `json:"DeliveryCallbackEnabled,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateServiceParams) SetApnCredentialSid(ApnCredentialSid string) *CreateServiceParams {
	params.ApnCredentialSid = &ApnCredentialSid
	return params
}
func (params *CreateServiceParams) SetGcmCredentialSid(GcmCredentialSid string) *CreateServiceParams {
	params.GcmCredentialSid = &GcmCredentialSid
	return params
}
func (params *CreateServiceParams) SetMessagingServiceSid(MessagingServiceSid string) *CreateServiceParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *CreateServiceParams) SetFacebookMessengerPageId(FacebookMessengerPageId string) *CreateServiceParams {
	params.FacebookMessengerPageId = &FacebookMessengerPageId
	return params
}
func (params *CreateServiceParams) SetDefaultApnNotificationProtocolVersion(DefaultApnNotificationProtocolVersion string) *CreateServiceParams {
	params.DefaultApnNotificationProtocolVersion = &DefaultApnNotificationProtocolVersion
	return params
}
func (params *CreateServiceParams) SetDefaultGcmNotificationProtocolVersion(DefaultGcmNotificationProtocolVersion string) *CreateServiceParams {
	params.DefaultGcmNotificationProtocolVersion = &DefaultGcmNotificationProtocolVersion
	return params
}
func (params *CreateServiceParams) SetFcmCredentialSid(FcmCredentialSid string) *CreateServiceParams {
	params.FcmCredentialSid = &FcmCredentialSid
	return params
}
func (params *CreateServiceParams) SetDefaultFcmNotificationProtocolVersion(DefaultFcmNotificationProtocolVersion string) *CreateServiceParams {
	params.DefaultFcmNotificationProtocolVersion = &DefaultFcmNotificationProtocolVersion
	return params
}
func (params *CreateServiceParams) SetLogEnabled(LogEnabled bool) *CreateServiceParams {
	params.LogEnabled = &LogEnabled
	return params
}
func (params *CreateServiceParams) SetAlexaSkillId(AlexaSkillId string) *CreateServiceParams {
	params.AlexaSkillId = &AlexaSkillId
	return params
}
func (params *CreateServiceParams) SetDefaultAlexaNotificationProtocolVersion(DefaultAlexaNotificationProtocolVersion string) *CreateServiceParams {
	params.DefaultAlexaNotificationProtocolVersion = &DefaultAlexaNotificationProtocolVersion
	return params
}
func (params *CreateServiceParams) SetDeliveryCallbackUrl(DeliveryCallbackUrl string) *CreateServiceParams {
	params.DeliveryCallbackUrl = &DeliveryCallbackUrl
	return params
}
func (params *CreateServiceParams) SetDeliveryCallbackEnabled(DeliveryCallbackEnabled bool) *CreateServiceParams {
	params.DeliveryCallbackEnabled = &DeliveryCallbackEnabled
	return params
}

func (c *ApiService) CreateService(params *CreateServiceParams) (*NotifyV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ApnCredentialSid != nil {
		data.Set("ApnCredentialSid", *params.ApnCredentialSid)
	}
	if params != nil && params.GcmCredentialSid != nil {
		data.Set("GcmCredentialSid", *params.GcmCredentialSid)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.FacebookMessengerPageId != nil {
		data.Set("FacebookMessengerPageId", *params.FacebookMessengerPageId)
	}
	if params != nil && params.DefaultApnNotificationProtocolVersion != nil {
		data.Set("DefaultApnNotificationProtocolVersion", *params.DefaultApnNotificationProtocolVersion)
	}
	if params != nil && params.DefaultGcmNotificationProtocolVersion != nil {
		data.Set("DefaultGcmNotificationProtocolVersion", *params.DefaultGcmNotificationProtocolVersion)
	}
	if params != nil && params.FcmCredentialSid != nil {
		data.Set("FcmCredentialSid", *params.FcmCredentialSid)
	}
	if params != nil && params.DefaultFcmNotificationProtocolVersion != nil {
		data.Set("DefaultFcmNotificationProtocolVersion", *params.DefaultFcmNotificationProtocolVersion)
	}
	if params != nil && params.LogEnabled != nil {
		data.Set("LogEnabled", fmt.Sprint(*params.LogEnabled))
	}
	if params != nil && params.AlexaSkillId != nil {
		data.Set("AlexaSkillId", *params.AlexaSkillId)
	}
	if params != nil && params.DefaultAlexaNotificationProtocolVersion != nil {
		data.Set("DefaultAlexaNotificationProtocolVersion", *params.DefaultAlexaNotificationProtocolVersion)
	}
	if params != nil && params.DeliveryCallbackUrl != nil {
		data.Set("DeliveryCallbackUrl", *params.DeliveryCallbackUrl)
	}
	if params != nil && params.DeliveryCallbackEnabled != nil {
		data.Set("DeliveryCallbackEnabled", fmt.Sprint(*params.DeliveryCallbackEnabled))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NotifyV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
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

func (c *ApiService) FetchService(Sid string) (*NotifyV1Service, error) {
	path := "/v1/Services/{Sid}"
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

	ps := &NotifyV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// The string that identifies the Service resources to read.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetFriendlyName(FriendlyName string) *ListServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParams) SetLimit(Limit int) *ListServiceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken, pageNumber string) (*ListServiceResponse, error) {
	path := "/v1/Services"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
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

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]NotifyV1Service, error) {
	response, errors := c.StreamService(params)

	records := make([]NotifyV1Service, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan NotifyV1Service, chan error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NotifyV1Service, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageService(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamService(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamService(response *ListServiceResponse, params *ListServiceParams, recordChannel chan NotifyV1Service, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Services
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListServiceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListServiceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateService'
type UpdateServiceParams struct {
	// A descriptive string that you create to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for APN Bindings.
	ApnCredentialSid *string `json:"ApnCredentialSid,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for GCM Bindings.
	GcmCredentialSid *string `json:"GcmCredentialSid,omitempty"`
	// The SID of the [Messaging Service](https://www.twilio.com/docs/sms/quickstart#messaging-services) to use for SMS Bindings. This parameter must be set in order to send SMS notifications.
	MessagingServiceSid *string `json:"MessagingServiceSid,omitempty"`
	// Deprecated.
	FacebookMessengerPageId *string `json:"FacebookMessengerPageId,omitempty"`
	// The protocol version to use for sending APNS notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultApnNotificationProtocolVersion *string `json:"DefaultApnNotificationProtocolVersion,omitempty"`
	// The protocol version to use for sending GCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultGcmNotificationProtocolVersion *string `json:"DefaultGcmNotificationProtocolVersion,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/notify/api/credential-resource) to use for FCM Bindings.
	FcmCredentialSid *string `json:"FcmCredentialSid,omitempty"`
	// The protocol version to use for sending FCM notifications. Can be overridden on a Binding by Binding basis when creating a [Binding](https://www.twilio.com/docs/notify/api/binding-resource) resource.
	DefaultFcmNotificationProtocolVersion *string `json:"DefaultFcmNotificationProtocolVersion,omitempty"`
	// Whether to log notifications. Can be: `true` or `false` and the default is `true`.
	LogEnabled *bool `json:"LogEnabled,omitempty"`
	// Deprecated.
	AlexaSkillId *string `json:"AlexaSkillId,omitempty"`
	// Deprecated.
	DefaultAlexaNotificationProtocolVersion *string `json:"DefaultAlexaNotificationProtocolVersion,omitempty"`
	// URL to send delivery status callback.
	DeliveryCallbackUrl *string `json:"DeliveryCallbackUrl,omitempty"`
	// Callback configuration that enables delivery callbacks, default false
	DeliveryCallbackEnabled *bool `json:"DeliveryCallbackEnabled,omitempty"`
}

func (params *UpdateServiceParams) SetFriendlyName(FriendlyName string) *UpdateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateServiceParams) SetApnCredentialSid(ApnCredentialSid string) *UpdateServiceParams {
	params.ApnCredentialSid = &ApnCredentialSid
	return params
}
func (params *UpdateServiceParams) SetGcmCredentialSid(GcmCredentialSid string) *UpdateServiceParams {
	params.GcmCredentialSid = &GcmCredentialSid
	return params
}
func (params *UpdateServiceParams) SetMessagingServiceSid(MessagingServiceSid string) *UpdateServiceParams {
	params.MessagingServiceSid = &MessagingServiceSid
	return params
}
func (params *UpdateServiceParams) SetFacebookMessengerPageId(FacebookMessengerPageId string) *UpdateServiceParams {
	params.FacebookMessengerPageId = &FacebookMessengerPageId
	return params
}
func (params *UpdateServiceParams) SetDefaultApnNotificationProtocolVersion(DefaultApnNotificationProtocolVersion string) *UpdateServiceParams {
	params.DefaultApnNotificationProtocolVersion = &DefaultApnNotificationProtocolVersion
	return params
}
func (params *UpdateServiceParams) SetDefaultGcmNotificationProtocolVersion(DefaultGcmNotificationProtocolVersion string) *UpdateServiceParams {
	params.DefaultGcmNotificationProtocolVersion = &DefaultGcmNotificationProtocolVersion
	return params
}
func (params *UpdateServiceParams) SetFcmCredentialSid(FcmCredentialSid string) *UpdateServiceParams {
	params.FcmCredentialSid = &FcmCredentialSid
	return params
}
func (params *UpdateServiceParams) SetDefaultFcmNotificationProtocolVersion(DefaultFcmNotificationProtocolVersion string) *UpdateServiceParams {
	params.DefaultFcmNotificationProtocolVersion = &DefaultFcmNotificationProtocolVersion
	return params
}
func (params *UpdateServiceParams) SetLogEnabled(LogEnabled bool) *UpdateServiceParams {
	params.LogEnabled = &LogEnabled
	return params
}
func (params *UpdateServiceParams) SetAlexaSkillId(AlexaSkillId string) *UpdateServiceParams {
	params.AlexaSkillId = &AlexaSkillId
	return params
}
func (params *UpdateServiceParams) SetDefaultAlexaNotificationProtocolVersion(DefaultAlexaNotificationProtocolVersion string) *UpdateServiceParams {
	params.DefaultAlexaNotificationProtocolVersion = &DefaultAlexaNotificationProtocolVersion
	return params
}
func (params *UpdateServiceParams) SetDeliveryCallbackUrl(DeliveryCallbackUrl string) *UpdateServiceParams {
	params.DeliveryCallbackUrl = &DeliveryCallbackUrl
	return params
}
func (params *UpdateServiceParams) SetDeliveryCallbackEnabled(DeliveryCallbackEnabled bool) *UpdateServiceParams {
	params.DeliveryCallbackEnabled = &DeliveryCallbackEnabled
	return params
}

func (c *ApiService) UpdateService(Sid string, params *UpdateServiceParams) (*NotifyV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.ApnCredentialSid != nil {
		data.Set("ApnCredentialSid", *params.ApnCredentialSid)
	}
	if params != nil && params.GcmCredentialSid != nil {
		data.Set("GcmCredentialSid", *params.GcmCredentialSid)
	}
	if params != nil && params.MessagingServiceSid != nil {
		data.Set("MessagingServiceSid", *params.MessagingServiceSid)
	}
	if params != nil && params.FacebookMessengerPageId != nil {
		data.Set("FacebookMessengerPageId", *params.FacebookMessengerPageId)
	}
	if params != nil && params.DefaultApnNotificationProtocolVersion != nil {
		data.Set("DefaultApnNotificationProtocolVersion", *params.DefaultApnNotificationProtocolVersion)
	}
	if params != nil && params.DefaultGcmNotificationProtocolVersion != nil {
		data.Set("DefaultGcmNotificationProtocolVersion", *params.DefaultGcmNotificationProtocolVersion)
	}
	if params != nil && params.FcmCredentialSid != nil {
		data.Set("FcmCredentialSid", *params.FcmCredentialSid)
	}
	if params != nil && params.DefaultFcmNotificationProtocolVersion != nil {
		data.Set("DefaultFcmNotificationProtocolVersion", *params.DefaultFcmNotificationProtocolVersion)
	}
	if params != nil && params.LogEnabled != nil {
		data.Set("LogEnabled", fmt.Sprint(*params.LogEnabled))
	}
	if params != nil && params.AlexaSkillId != nil {
		data.Set("AlexaSkillId", *params.AlexaSkillId)
	}
	if params != nil && params.DefaultAlexaNotificationProtocolVersion != nil {
		data.Set("DefaultAlexaNotificationProtocolVersion", *params.DefaultAlexaNotificationProtocolVersion)
	}
	if params != nil && params.DeliveryCallbackUrl != nil {
		data.Set("DeliveryCallbackUrl", *params.DeliveryCallbackUrl)
	}
	if params != nil && params.DeliveryCallbackEnabled != nil {
		data.Set("DeliveryCallbackEnabled", fmt.Sprint(*params.DeliveryCallbackEnabled))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NotifyV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
