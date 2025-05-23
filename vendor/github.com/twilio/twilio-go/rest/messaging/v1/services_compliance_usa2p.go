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

// Optional parameters for the method 'CreateUsAppToPerson'
type CreateUsAppToPersonParams struct {
	// A2P Brand Registration SID
	BrandRegistrationSid *string `json:"BrandRegistrationSid,omitempty"`
	// A short description of what this SMS campaign does. Min length: 40 characters. Max length: 4096 characters.
	Description *string `json:"Description,omitempty"`
	// Required for all Campaigns. Details around how a consumer opts-in to their campaign, therefore giving consent to receive their messages. If multiple opt-in methods can be used for the same campaign, they must all be listed. 40 character minimum. 2048 character maximum.
	MessageFlow *string `json:"MessageFlow,omitempty"`
	// An array of sample message strings, min two and max five. Min length for each sample: 20 chars. Max length for each sample: 1024 chars.
	MessageSamples *[]string `json:"MessageSamples,omitempty"`
	// A2P Campaign Use Case. Examples: [ 2FA, EMERGENCY, MARKETING..]
	UsAppToPersonUsecase *string `json:"UsAppToPersonUsecase,omitempty"`
	// Indicates that this SMS campaign will send messages that contain links.
	HasEmbeddedLinks *bool `json:"HasEmbeddedLinks,omitempty"`
	// Indicates that this SMS campaign will send messages that contain phone numbers.
	HasEmbeddedPhone *bool `json:"HasEmbeddedPhone,omitempty"`
	// If end users can text in a keyword to start receiving messages from this campaign, the auto-reply messages sent to the end users must be provided. The opt-in response should include the Brand name, confirmation of opt-in enrollment to a recurring message campaign, how to get help, and clear description of how to opt-out. This field is required if end users can text in a keyword to start receiving messages from this campaign. 20 character minimum. 320 character maximum.
	OptInMessage *string `json:"OptInMessage,omitempty"`
	// Upon receiving the opt-out keywords from the end users, Twilio customers are expected to send back an auto-generated response, which must provide acknowledgment of the opt-out request and confirmation that no further messages will be sent. It is also recommended that these opt-out messages include the brand name. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum.
	OptOutMessage *string `json:"OptOutMessage,omitempty"`
	// When customers receive the help keywords from their end users, Twilio customers are expected to send back an auto-generated response; this may include the brand name and additional support contact information. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). 20 character minimum. 320 character maximum.
	HelpMessage *string `json:"HelpMessage,omitempty"`
	// If end users can text in a keyword to start receiving messages from this campaign, those keywords must be provided. This field is required if end users can text in a keyword to start receiving messages from this campaign. Values must be alphanumeric. 255 character maximum.
	OptInKeywords *[]string `json:"OptInKeywords,omitempty"`
	// End users should be able to text in a keyword to stop receiving messages from this campaign. Those keywords must be provided. This field is required if managing opt out keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum.
	OptOutKeywords *[]string `json:"OptOutKeywords,omitempty"`
	// End users should be able to text in a keyword to receive help. Those keywords must be provided as part of the campaign registration request. This field is required if managing help keywords yourself (i.e. not using Twilio's Default or Advanced Opt Out features). Values must be alphanumeric. 255 character maximum.
	HelpKeywords *[]string `json:"HelpKeywords,omitempty"`
	// A boolean that specifies whether campaign has Subscriber Optin or not.
	SubscriberOptIn *bool `json:"SubscriberOptIn,omitempty"`
	// A boolean that specifies whether campaign is age gated or not.
	AgeGated *bool `json:"AgeGated,omitempty"`
	// A boolean that specifies whether campaign allows direct lending or not.
	DirectLending *bool `json:"DirectLending,omitempty"`
}

func (params *CreateUsAppToPersonParams) SetBrandRegistrationSid(BrandRegistrationSid string) *CreateUsAppToPersonParams {
	params.BrandRegistrationSid = &BrandRegistrationSid
	return params
}
func (params *CreateUsAppToPersonParams) SetDescription(Description string) *CreateUsAppToPersonParams {
	params.Description = &Description
	return params
}
func (params *CreateUsAppToPersonParams) SetMessageFlow(MessageFlow string) *CreateUsAppToPersonParams {
	params.MessageFlow = &MessageFlow
	return params
}
func (params *CreateUsAppToPersonParams) SetMessageSamples(MessageSamples []string) *CreateUsAppToPersonParams {
	params.MessageSamples = &MessageSamples
	return params
}
func (params *CreateUsAppToPersonParams) SetUsAppToPersonUsecase(UsAppToPersonUsecase string) *CreateUsAppToPersonParams {
	params.UsAppToPersonUsecase = &UsAppToPersonUsecase
	return params
}
func (params *CreateUsAppToPersonParams) SetHasEmbeddedLinks(HasEmbeddedLinks bool) *CreateUsAppToPersonParams {
	params.HasEmbeddedLinks = &HasEmbeddedLinks
	return params
}
func (params *CreateUsAppToPersonParams) SetHasEmbeddedPhone(HasEmbeddedPhone bool) *CreateUsAppToPersonParams {
	params.HasEmbeddedPhone = &HasEmbeddedPhone
	return params
}
func (params *CreateUsAppToPersonParams) SetOptInMessage(OptInMessage string) *CreateUsAppToPersonParams {
	params.OptInMessage = &OptInMessage
	return params
}
func (params *CreateUsAppToPersonParams) SetOptOutMessage(OptOutMessage string) *CreateUsAppToPersonParams {
	params.OptOutMessage = &OptOutMessage
	return params
}
func (params *CreateUsAppToPersonParams) SetHelpMessage(HelpMessage string) *CreateUsAppToPersonParams {
	params.HelpMessage = &HelpMessage
	return params
}
func (params *CreateUsAppToPersonParams) SetOptInKeywords(OptInKeywords []string) *CreateUsAppToPersonParams {
	params.OptInKeywords = &OptInKeywords
	return params
}
func (params *CreateUsAppToPersonParams) SetOptOutKeywords(OptOutKeywords []string) *CreateUsAppToPersonParams {
	params.OptOutKeywords = &OptOutKeywords
	return params
}
func (params *CreateUsAppToPersonParams) SetHelpKeywords(HelpKeywords []string) *CreateUsAppToPersonParams {
	params.HelpKeywords = &HelpKeywords
	return params
}
func (params *CreateUsAppToPersonParams) SetSubscriberOptIn(SubscriberOptIn bool) *CreateUsAppToPersonParams {
	params.SubscriberOptIn = &SubscriberOptIn
	return params
}
func (params *CreateUsAppToPersonParams) SetAgeGated(AgeGated bool) *CreateUsAppToPersonParams {
	params.AgeGated = &AgeGated
	return params
}
func (params *CreateUsAppToPersonParams) SetDirectLending(DirectLending bool) *CreateUsAppToPersonParams {
	params.DirectLending = &DirectLending
	return params
}

func (c *ApiService) CreateUsAppToPerson(MessagingServiceSid string, params *CreateUsAppToPersonParams) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.BrandRegistrationSid != nil {
		data.Set("BrandRegistrationSid", *params.BrandRegistrationSid)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.MessageFlow != nil {
		data.Set("MessageFlow", *params.MessageFlow)
	}
	if params != nil && params.MessageSamples != nil {
		for _, item := range *params.MessageSamples {
			data.Add("MessageSamples", item)
		}
	}
	if params != nil && params.UsAppToPersonUsecase != nil {
		data.Set("UsAppToPersonUsecase", *params.UsAppToPersonUsecase)
	}
	if params != nil && params.HasEmbeddedLinks != nil {
		data.Set("HasEmbeddedLinks", fmt.Sprint(*params.HasEmbeddedLinks))
	}
	if params != nil && params.HasEmbeddedPhone != nil {
		data.Set("HasEmbeddedPhone", fmt.Sprint(*params.HasEmbeddedPhone))
	}
	if params != nil && params.OptInMessage != nil {
		data.Set("OptInMessage", *params.OptInMessage)
	}
	if params != nil && params.OptOutMessage != nil {
		data.Set("OptOutMessage", *params.OptOutMessage)
	}
	if params != nil && params.HelpMessage != nil {
		data.Set("HelpMessage", *params.HelpMessage)
	}
	if params != nil && params.OptInKeywords != nil {
		for _, item := range *params.OptInKeywords {
			data.Add("OptInKeywords", item)
		}
	}
	if params != nil && params.OptOutKeywords != nil {
		for _, item := range *params.OptOutKeywords {
			data.Add("OptOutKeywords", item)
		}
	}
	if params != nil && params.HelpKeywords != nil {
		for _, item := range *params.HelpKeywords {
			data.Add("HelpKeywords", item)
		}
	}
	if params != nil && params.SubscriberOptIn != nil {
		data.Set("SubscriberOptIn", fmt.Sprint(*params.SubscriberOptIn))
	}
	if params != nil && params.AgeGated != nil {
		data.Set("AgeGated", fmt.Sprint(*params.AgeGated))
	}
	if params != nil && params.DirectLending != nil {
		data.Set("DirectLending", fmt.Sprint(*params.DirectLending))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteUsAppToPerson(MessagingServiceSid string, Sid string) error {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
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

func (c *ApiService) FetchUsAppToPerson(MessagingServiceSid string, Sid string) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
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

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUsAppToPerson'
type ListUsAppToPersonParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsAppToPersonParams) SetPageSize(PageSize int) *ListUsAppToPersonParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsAppToPersonParams) SetLimit(Limit int) *ListUsAppToPersonParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsAppToPerson records from the API. Request is executed immediately.
func (c *ApiService) PageUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams, pageToken, pageNumber string) (*ListUsAppToPersonResponse, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p"

	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

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

	ps := &ListUsAppToPersonResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsAppToPerson records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams) ([]MessagingV1UsAppToPerson, error) {
	response, errors := c.StreamUsAppToPerson(MessagingServiceSid, params)

	records := make([]MessagingV1UsAppToPerson, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams UsAppToPerson records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsAppToPerson(MessagingServiceSid string, params *ListUsAppToPersonParams) (chan MessagingV1UsAppToPerson, chan error) {
	if params == nil {
		params = &ListUsAppToPersonParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan MessagingV1UsAppToPerson, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUsAppToPerson(MessagingServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUsAppToPerson(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUsAppToPerson(response *ListUsAppToPersonResponse, params *ListUsAppToPersonParams, recordChannel chan MessagingV1UsAppToPerson, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Compliance
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListUsAppToPersonResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUsAppToPersonResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUsAppToPersonResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsAppToPersonResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateUsAppToPerson'
type UpdateUsAppToPersonParams struct {
	// Indicates that this SMS campaign will send messages that contain links.
	HasEmbeddedLinks *bool `json:"HasEmbeddedLinks,omitempty"`
	// Indicates that this SMS campaign will send messages that contain phone numbers.
	HasEmbeddedPhone *bool `json:"HasEmbeddedPhone,omitempty"`
	// An array of sample message strings, min two and max five. Min length for each sample: 20 chars. Max length for each sample: 1024 chars.
	MessageSamples *[]string `json:"MessageSamples,omitempty"`
	// Required for all Campaigns. Details around how a consumer opts-in to their campaign, therefore giving consent to receive their messages. If multiple opt-in methods can be used for the same campaign, they must all be listed. 40 character minimum. 2048 character maximum.
	MessageFlow *string `json:"MessageFlow,omitempty"`
	// A short description of what this SMS campaign does. Min length: 40 characters. Max length: 4096 characters.
	Description *string `json:"Description,omitempty"`
	// A boolean that specifies whether campaign requires age gate for federally legal content.
	AgeGated *bool `json:"AgeGated,omitempty"`
	// A boolean that specifies whether campaign allows direct lending or not.
	DirectLending *bool `json:"DirectLending,omitempty"`
}

func (params *UpdateUsAppToPersonParams) SetHasEmbeddedLinks(HasEmbeddedLinks bool) *UpdateUsAppToPersonParams {
	params.HasEmbeddedLinks = &HasEmbeddedLinks
	return params
}
func (params *UpdateUsAppToPersonParams) SetHasEmbeddedPhone(HasEmbeddedPhone bool) *UpdateUsAppToPersonParams {
	params.HasEmbeddedPhone = &HasEmbeddedPhone
	return params
}
func (params *UpdateUsAppToPersonParams) SetMessageSamples(MessageSamples []string) *UpdateUsAppToPersonParams {
	params.MessageSamples = &MessageSamples
	return params
}
func (params *UpdateUsAppToPersonParams) SetMessageFlow(MessageFlow string) *UpdateUsAppToPersonParams {
	params.MessageFlow = &MessageFlow
	return params
}
func (params *UpdateUsAppToPersonParams) SetDescription(Description string) *UpdateUsAppToPersonParams {
	params.Description = &Description
	return params
}
func (params *UpdateUsAppToPersonParams) SetAgeGated(AgeGated bool) *UpdateUsAppToPersonParams {
	params.AgeGated = &AgeGated
	return params
}
func (params *UpdateUsAppToPersonParams) SetDirectLending(DirectLending bool) *UpdateUsAppToPersonParams {
	params.DirectLending = &DirectLending
	return params
}

func (c *ApiService) UpdateUsAppToPerson(MessagingServiceSid string, Sid string, params *UpdateUsAppToPersonParams) (*MessagingV1UsAppToPerson, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/{Sid}"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.HasEmbeddedLinks != nil {
		data.Set("HasEmbeddedLinks", fmt.Sprint(*params.HasEmbeddedLinks))
	}
	if params != nil && params.HasEmbeddedPhone != nil {
		data.Set("HasEmbeddedPhone", fmt.Sprint(*params.HasEmbeddedPhone))
	}
	if params != nil && params.MessageSamples != nil {
		for _, item := range *params.MessageSamples {
			data.Add("MessageSamples", item)
		}
	}
	if params != nil && params.MessageFlow != nil {
		data.Set("MessageFlow", *params.MessageFlow)
	}
	if params != nil && params.Description != nil {
		data.Set("Description", *params.Description)
	}
	if params != nil && params.AgeGated != nil {
		data.Set("AgeGated", fmt.Sprint(*params.AgeGated))
	}
	if params != nil && params.DirectLending != nil {
		data.Set("DirectLending", fmt.Sprint(*params.DirectLending))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPerson{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
