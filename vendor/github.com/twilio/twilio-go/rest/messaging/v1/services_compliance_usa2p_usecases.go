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
	"net/url"
	"strings"
)

// Optional parameters for the method 'FetchUsAppToPersonUsecase'
type FetchUsAppToPersonUsecaseParams struct {
	// The unique string to identify the A2P brand.
	BrandRegistrationSid *string `json:"BrandRegistrationSid,omitempty"`
}

func (params *FetchUsAppToPersonUsecaseParams) SetBrandRegistrationSid(BrandRegistrationSid string) *FetchUsAppToPersonUsecaseParams {
	params.BrandRegistrationSid = &BrandRegistrationSid
	return params
}

func (c *ApiService) FetchUsAppToPersonUsecase(MessagingServiceSid string, params *FetchUsAppToPersonUsecaseParams) (*MessagingV1UsAppToPersonUsecase, error) {
	path := "/v1/Services/{MessagingServiceSid}/Compliance/Usa2p/Usecases"
	path = strings.Replace(path, "{"+"MessagingServiceSid"+"}", MessagingServiceSid, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.BrandRegistrationSid != nil {
		data.Set("BrandRegistrationSid", *params.BrandRegistrationSid)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1UsAppToPersonUsecase{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
