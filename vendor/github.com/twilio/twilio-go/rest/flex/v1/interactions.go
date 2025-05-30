/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
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

// Optional parameters for the method 'CreateInteraction'
type CreateInteractionParams struct {
	// The Interaction's channel.
	Channel *interface{} `json:"Channel,omitempty"`
	// The Interaction's routing logic.
	Routing *interface{} `json:"Routing,omitempty"`
	// The Interaction context sid is used for adding a context lookup sid
	InteractionContextSid *string `json:"InteractionContextSid,omitempty"`
}

func (params *CreateInteractionParams) SetChannel(Channel interface{}) *CreateInteractionParams {
	params.Channel = &Channel
	return params
}
func (params *CreateInteractionParams) SetRouting(Routing interface{}) *CreateInteractionParams {
	params.Routing = &Routing
	return params
}
func (params *CreateInteractionParams) SetInteractionContextSid(InteractionContextSid string) *CreateInteractionParams {
	params.InteractionContextSid = &InteractionContextSid
	return params
}

// Create a new Interaction.
func (c *ApiService) CreateInteraction(params *CreateInteractionParams) (*FlexV1Interaction, error) {
	path := "/v1/Interactions"

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.Channel != nil {
		v, err := json.Marshal(params.Channel)

		if err != nil {
			return nil, err
		}

		data.Set("Channel", string(v))
	}
	if params != nil && params.Routing != nil {
		v, err := json.Marshal(params.Routing)

		if err != nil {
			return nil, err
		}

		data.Set("Routing", string(v))
	}
	if params != nil && params.InteractionContextSid != nil {
		data.Set("InteractionContextSid", *params.InteractionContextSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Interaction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) FetchInteraction(Sid string) (*FlexV1Interaction, error) {
	path := "/v1/Interactions/{Sid}"
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

	ps := &FlexV1Interaction{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
