/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
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

// Optional parameters for the method 'CreateWorkspace'
type CreateWorkspaceParams struct {
	// A descriptive string that you create to describe the Workspace resource. It can be up to 64 characters long. For example: `Customer Support` or `2014 Election Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// The URL we should call when an event occurs. If provided, the Workspace will publish events to this URL, for example, to collect data for reporting. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. This parameter supports Twilio's [Webhooks (HTTP callbacks) Connection Overrides](https://www.twilio.com/docs/usage/webhooks/webhooks-connection-overrides).
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty"`
	// The list of Workspace events for which to call event_callback_url. For example, if `EventsFilter=task.created, task.canceled, worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
	EventsFilter *string `json:"EventsFilter,omitempty"`
	// Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. However, all workspaces should be created as multi-tasking. The default is `true`. Multi-tasking allows Workers to handle multiple Tasks simultaneously. When enabled (`true`), each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. In single-tasking mode (legacy mode), each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
	MultiTaskEnabled *bool `json:"MultiTaskEnabled,omitempty"`
	// An available template name. Can be: `NONE` or `FIFO` and the default is `NONE`. Pre-configures the Workspace with the Workflow and Activities specified in the template. `NONE` will create a Workspace with only a set of default activities. `FIFO` will configure TaskRouter with a set of default activities and a single TaskQueue for first-in, first-out distribution, which can be useful when you are getting started with TaskRouter.
	Template *string `json:"Template,omitempty"`
	//
	PrioritizeQueueOrder *string `json:"PrioritizeQueueOrder,omitempty"`
}

func (params *CreateWorkspaceParams) SetFriendlyName(FriendlyName string) *CreateWorkspaceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateWorkspaceParams) SetEventCallbackUrl(EventCallbackUrl string) *CreateWorkspaceParams {
	params.EventCallbackUrl = &EventCallbackUrl
	return params
}
func (params *CreateWorkspaceParams) SetEventsFilter(EventsFilter string) *CreateWorkspaceParams {
	params.EventsFilter = &EventsFilter
	return params
}
func (params *CreateWorkspaceParams) SetMultiTaskEnabled(MultiTaskEnabled bool) *CreateWorkspaceParams {
	params.MultiTaskEnabled = &MultiTaskEnabled
	return params
}
func (params *CreateWorkspaceParams) SetTemplate(Template string) *CreateWorkspaceParams {
	params.Template = &Template
	return params
}
func (params *CreateWorkspaceParams) SetPrioritizeQueueOrder(PrioritizeQueueOrder string) *CreateWorkspaceParams {
	params.PrioritizeQueueOrder = &PrioritizeQueueOrder
	return params
}

func (c *ApiService) CreateWorkspace(params *CreateWorkspaceParams) (*TaskrouterV1Workspace, error) {
	path := "/v1/Workspaces"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.EventCallbackUrl != nil {
		data.Set("EventCallbackUrl", *params.EventCallbackUrl)
	}
	if params != nil && params.EventsFilter != nil {
		data.Set("EventsFilter", *params.EventsFilter)
	}
	if params != nil && params.MultiTaskEnabled != nil {
		data.Set("MultiTaskEnabled", fmt.Sprint(*params.MultiTaskEnabled))
	}
	if params != nil && params.Template != nil {
		data.Set("Template", *params.Template)
	}
	if params != nil && params.PrioritizeQueueOrder != nil {
		data.Set("PrioritizeQueueOrder", *params.PrioritizeQueueOrder)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workspace{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteWorkspace(Sid string) error {
	path := "/v1/Workspaces/{Sid}"
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

func (c *ApiService) FetchWorkspace(Sid string) (*TaskrouterV1Workspace, error) {
	path := "/v1/Workspaces/{Sid}"
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

	ps := &TaskrouterV1Workspace{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListWorkspace'
type ListWorkspaceParams struct {
	// The `friendly_name` of the Workspace resources to read. For example `Customer Support` or `2014 Election Campaign`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListWorkspaceParams) SetFriendlyName(FriendlyName string) *ListWorkspaceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *ListWorkspaceParams) SetPageSize(PageSize int) *ListWorkspaceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListWorkspaceParams) SetLimit(Limit int) *ListWorkspaceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Workspace records from the API. Request is executed immediately.
func (c *ApiService) PageWorkspace(params *ListWorkspaceParams, pageToken, pageNumber string) (*ListWorkspaceResponse, error) {
	path := "/v1/Workspaces"

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

	ps := &ListWorkspaceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Workspace records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListWorkspace(params *ListWorkspaceParams) ([]TaskrouterV1Workspace, error) {
	response, errors := c.StreamWorkspace(params)

	records := make([]TaskrouterV1Workspace, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Workspace records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamWorkspace(params *ListWorkspaceParams) (chan TaskrouterV1Workspace, chan error) {
	if params == nil {
		params = &ListWorkspaceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TaskrouterV1Workspace, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageWorkspace(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamWorkspace(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamWorkspace(response *ListWorkspaceResponse, params *ListWorkspaceParams, recordChannel chan TaskrouterV1Workspace, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Workspaces
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListWorkspaceResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListWorkspaceResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListWorkspaceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListWorkspaceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateWorkspace'
type UpdateWorkspaceParams struct {
	// The SID of the Activity that will be used when new Workers are created in the Workspace.
	DefaultActivitySid *string `json:"DefaultActivitySid,omitempty"`
	// The URL we should call when an event occurs. See [Workspace Events](https://www.twilio.com/docs/taskrouter/api/event) for more information. This parameter supports Twilio's [Webhooks (HTTP callbacks) Connection Overrides](https://www.twilio.com/docs/usage/webhooks/webhooks-connection-overrides).
	EventCallbackUrl *string `json:"EventCallbackUrl,omitempty"`
	// The list of Workspace events for which to call event_callback_url. For example if `EventsFilter=task.created,task.canceled,worker.activity.update`, then TaskRouter will call event_callback_url only when a task is created, canceled, or a Worker activity is updated.
	EventsFilter *string `json:"EventsFilter,omitempty"`
	// A descriptive string that you create to describe the Workspace resource. For example: `Sales Call Center` or `Customer Support Team`.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether to enable multi-tasking. Can be: `true` to enable multi-tasking, or `false` to disable it. However, all workspaces should be maintained as multi-tasking. There is no default when omitting this parameter. A multi-tasking Workspace can't be updated to single-tasking unless it is not a Flex Project and another (legacy) single-tasking Workspace exists. Multi-tasking allows Workers to handle multiple Tasks simultaneously. In multi-tasking mode, each Worker can receive parallel reservations up to the per-channel maximums defined in the Workers section. In single-tasking mode (legacy mode), each Worker will only receive a new reservation when the previous task is completed. Learn more at [Multitasking](https://www.twilio.com/docs/taskrouter/multitasking).
	MultiTaskEnabled *bool `json:"MultiTaskEnabled,omitempty"`
	// The SID of the Activity that will be assigned to a Worker when a Task reservation times out without a response.
	TimeoutActivitySid *string `json:"TimeoutActivitySid,omitempty"`
	//
	PrioritizeQueueOrder *string `json:"PrioritizeQueueOrder,omitempty"`
}

func (params *UpdateWorkspaceParams) SetDefaultActivitySid(DefaultActivitySid string) *UpdateWorkspaceParams {
	params.DefaultActivitySid = &DefaultActivitySid
	return params
}
func (params *UpdateWorkspaceParams) SetEventCallbackUrl(EventCallbackUrl string) *UpdateWorkspaceParams {
	params.EventCallbackUrl = &EventCallbackUrl
	return params
}
func (params *UpdateWorkspaceParams) SetEventsFilter(EventsFilter string) *UpdateWorkspaceParams {
	params.EventsFilter = &EventsFilter
	return params
}
func (params *UpdateWorkspaceParams) SetFriendlyName(FriendlyName string) *UpdateWorkspaceParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateWorkspaceParams) SetMultiTaskEnabled(MultiTaskEnabled bool) *UpdateWorkspaceParams {
	params.MultiTaskEnabled = &MultiTaskEnabled
	return params
}
func (params *UpdateWorkspaceParams) SetTimeoutActivitySid(TimeoutActivitySid string) *UpdateWorkspaceParams {
	params.TimeoutActivitySid = &TimeoutActivitySid
	return params
}
func (params *UpdateWorkspaceParams) SetPrioritizeQueueOrder(PrioritizeQueueOrder string) *UpdateWorkspaceParams {
	params.PrioritizeQueueOrder = &PrioritizeQueueOrder
	return params
}

func (c *ApiService) UpdateWorkspace(Sid string, params *UpdateWorkspaceParams) (*TaskrouterV1Workspace, error) {
	path := "/v1/Workspaces/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.DefaultActivitySid != nil {
		data.Set("DefaultActivitySid", *params.DefaultActivitySid)
	}
	if params != nil && params.EventCallbackUrl != nil {
		data.Set("EventCallbackUrl", *params.EventCallbackUrl)
	}
	if params != nil && params.EventsFilter != nil {
		data.Set("EventsFilter", *params.EventsFilter)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.MultiTaskEnabled != nil {
		data.Set("MultiTaskEnabled", fmt.Sprint(*params.MultiTaskEnabled))
	}
	if params != nil && params.TimeoutActivitySid != nil {
		data.Set("TimeoutActivitySid", *params.TimeoutActivitySid)
	}
	if params != nil && params.PrioritizeQueueOrder != nil {
		data.Set("PrioritizeQueueOrder", *params.PrioritizeQueueOrder)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TaskrouterV1Workspace{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
