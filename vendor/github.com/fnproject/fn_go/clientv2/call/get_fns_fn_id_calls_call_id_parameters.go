// Code generated by go-swagger; DO NOT EDIT.

package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFnsFnIDCallsCallIDParams creates a new GetFnsFnIDCallsCallIDParams object
// with the default values initialized.
func NewGetFnsFnIDCallsCallIDParams() *GetFnsFnIDCallsCallIDParams {
	var ()
	return &GetFnsFnIDCallsCallIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFnsFnIDCallsCallIDParamsWithTimeout creates a new GetFnsFnIDCallsCallIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFnsFnIDCallsCallIDParamsWithTimeout(timeout time.Duration) *GetFnsFnIDCallsCallIDParams {
	var ()
	return &GetFnsFnIDCallsCallIDParams{

		timeout: timeout,
	}
}

// NewGetFnsFnIDCallsCallIDParamsWithContext creates a new GetFnsFnIDCallsCallIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFnsFnIDCallsCallIDParamsWithContext(ctx context.Context) *GetFnsFnIDCallsCallIDParams {
	var ()
	return &GetFnsFnIDCallsCallIDParams{

		Context: ctx,
	}
}

// NewGetFnsFnIDCallsCallIDParamsWithHTTPClient creates a new GetFnsFnIDCallsCallIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFnsFnIDCallsCallIDParamsWithHTTPClient(client *http.Client) *GetFnsFnIDCallsCallIDParams {
	var ()
	return &GetFnsFnIDCallsCallIDParams{
		HTTPClient: client,
	}
}

/*GetFnsFnIDCallsCallIDParams contains all the parameters to send to the API endpoint
for the get fns fn ID calls call ID operation typically these are written to a http.Request
*/
type GetFnsFnIDCallsCallIDParams struct {

	/*CallID
	  Opaque, unique Call ID.

	*/
	CallID string
	/*FnID
	  Opaque, unique Function ID.

	*/
	FnID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) WithTimeout(timeout time.Duration) *GetFnsFnIDCallsCallIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) WithContext(ctx context.Context) *GetFnsFnIDCallsCallIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) WithHTTPClient(client *http.Client) *GetFnsFnIDCallsCallIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallID adds the callID to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) WithCallID(callID string) *GetFnsFnIDCallsCallIDParams {
	o.SetCallID(callID)
	return o
}

// SetCallID adds the callId to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) SetCallID(callID string) {
	o.CallID = callID
}

// WithFnID adds the fnID to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) WithFnID(fnID string) *GetFnsFnIDCallsCallIDParams {
	o.SetFnID(fnID)
	return o
}

// SetFnID adds the fnId to the get fns fn ID calls call ID params
func (o *GetFnsFnIDCallsCallIDParams) SetFnID(fnID string) {
	o.FnID = fnID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFnsFnIDCallsCallIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param callID
	if err := r.SetPathParam("callID", o.CallID); err != nil {
		return err
	}

	// path param fnID
	if err := r.SetPathParam("fnID", o.FnID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
