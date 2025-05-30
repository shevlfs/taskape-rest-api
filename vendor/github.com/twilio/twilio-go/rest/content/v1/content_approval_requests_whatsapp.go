/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Content
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateApprovalCreate'
type CreateApprovalCreateParams struct {
	//
	ContentApprovalRequest *ContentApprovalRequest `json:"ContentApprovalRequest,omitempty"`
}

func (params *CreateApprovalCreateParams) SetContentApprovalRequest(ContentApprovalRequest ContentApprovalRequest) *CreateApprovalCreateParams {
	params.ContentApprovalRequest = &ContentApprovalRequest
	return params
}

func (c *ApiService) CreateApprovalCreate(ContentSid string, params *CreateApprovalCreateParams) (*ContentV1ApprovalCreate, error) {
	path := "/v1/Content/{ContentSid}/ApprovalRequests/whatsapp"
	path = strings.Replace(path, "{"+"ContentSid"+"}", ContentSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/json",
	}

	body := []byte{}
	if params != nil && params.ContentApprovalRequest != nil {
		b, err := json.Marshal(*params.ContentApprovalRequest)
		if err != nil {
			return nil, err
		}
		body = b
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers, body...)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ContentV1ApprovalCreate{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
