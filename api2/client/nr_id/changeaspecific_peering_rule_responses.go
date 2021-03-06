// Code generated by go-swagger; DO NOT EDIT.

package nr_id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeaspecificPeeringRuleReader is a Reader for the ChangeaspecificPeeringRule structure.
type ChangeaspecificPeeringRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeaspecificPeeringRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewChangeaspecificPeeringRuleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeaspecificPeeringRuleNoContent creates a ChangeaspecificPeeringRuleNoContent with default headers values
func NewChangeaspecificPeeringRuleNoContent() *ChangeaspecificPeeringRuleNoContent {
	return &ChangeaspecificPeeringRuleNoContent{}
}

/*ChangeaspecificPeeringRuleNoContent handles this case with default header values.

ChangeaspecificPeeringRuleNoContent changeaspecific peering rule no content
*/
type ChangeaspecificPeeringRuleNoContent struct {
	Payload interface{}
}

func (o *ChangeaspecificPeeringRuleNoContent) Error() string {
	return fmt.Sprintf("[PATCH /peeringrules/{id}][%d] changeaspecificPeeringRuleNoContent  %+v", 204, o.Payload)
}

func (o *ChangeaspecificPeeringRuleNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeaspecificPeeringRuleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
