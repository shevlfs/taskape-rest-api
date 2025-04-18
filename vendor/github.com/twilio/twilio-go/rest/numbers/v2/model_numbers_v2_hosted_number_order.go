/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Numbers
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// NumbersV2HostedNumberOrder struct for NumbersV2HostedNumberOrder
type NumbersV2HostedNumberOrder struct {
	// A 34 character string that uniquely identifies this HostedNumberOrder.
	Sid *string `json:"sid,omitempty"`
	// A 34 character string that uniquely identifies the account.
	AccountSid *string `json:"account_sid,omitempty"`
	// A 34 character string that uniquely identifies the [IncomingPhoneNumber](https://www.twilio.com/docs/phone-numbers/api/incomingphonenumber-resource) resource that represents the phone number being hosted.
	IncomingPhoneNumberSid *string `json:"incoming_phone_number_sid,omitempty"`
	// A 34 character string that uniquely identifies the Address resource that represents the address of the owner of this phone number.
	AddressSid *string `json:"address_sid,omitempty"`
	// A 34 character string that uniquely identifies the [Authorization Document](https://www.twilio.com/docs/phone-numbers/hosted-numbers/hosted-numbers-api/authorization-document-resource) the user needs to sign.
	SigningDocumentSid *string `json:"signing_document_sid,omitempty"`
	// Phone number to be hosted. This must be in [E.164](https://en.wikipedia.org/wiki/E.164) format, e.g., +16175551212
	PhoneNumber  *string                                 `json:"phone_number,omitempty"`
	Capabilities *NumbersV2HostedNumberOrderCapabilities `json:"capabilities,omitempty"`
	// A 128 character string that is a human-readable text that describes this resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	Status       *string `json:"status,omitempty"`
	// A message that explains why a hosted_number_order went to status \"action-required\"
	FailureReason *string `json:"failure_reason,omitempty"`
	// The date this resource was created, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this resource was updated, given as [GMT RFC 2822](http://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Email of the owner of this phone number that is being hosted.
	Email *string `json:"email,omitempty"`
	// A list of emails that LOA document for this HostedNumberOrder will be carbon copied to.
	CcEmails *[]string `json:"cc_emails,omitempty"`
	// The URL of this HostedNumberOrder.
	Url *string `json:"url,omitempty"`
	// The title of the person authorized to sign the Authorization Document for this phone number.
	ContactTitle *string `json:"contact_title,omitempty"`
	// The contact phone number of the person authorized to sign the Authorization Document.
	ContactPhoneNumber *string `json:"contact_phone_number,omitempty"`
	// A 34 character string that uniquely identifies the bulk hosting request associated with this HostedNumberOrder.
	BulkHostingRequestSid *string `json:"bulk_hosting_request_sid,omitempty"`
	// The next step you need to take to complete the hosted number order and request it successfully.
	NextStep *string `json:"next_step,omitempty"`
	// The number of attempts made to verify ownership via a call for the hosted phone number.
	VerificationAttempts int `json:"verification_attempts,omitempty"`
	// The Call SIDs that identify the calls placed to verify ownership.
	VerificationCallSids *[]string `json:"verification_call_sids,omitempty"`
	// The number of seconds to wait before initiating the ownership verification call. Can be a value between 0 and 60, inclusive.
	VerificationCallDelay int `json:"verification_call_delay,omitempty"`
	// The numerical extension to dial when making the ownership verification call.
	VerificationCallExtension *string `json:"verification_call_extension,omitempty"`
	// The digits the user must pass in the ownership verification call.
	VerificationCode *string `json:"verification_code,omitempty"`
	VerificationType *string `json:"verification_type,omitempty"`
}
