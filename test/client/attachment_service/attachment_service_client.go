// Code generated by go-swagger; DO NOT EDIT.

package attachment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new attachment service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for attachment service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAttachment gets attachment
*/
func (a *Client) GetAttachment(params *GetAttachmentParams, authInfo runtime.ClientAuthInfoWriter) (*GetAttachmentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAttachment",
		Method:             "GET",
		PathPattern:        "/v1/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAttachmentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAttachmentOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
