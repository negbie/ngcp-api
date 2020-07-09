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

// DeleteaspecificSubscriberRegistrationReader is a Reader for the DeleteaspecificSubscriberRegistration structure.
type DeleteaspecificSubscriberRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificSubscriberRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificSubscriberRegistrationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificSubscriberRegistrationNoContent creates a DeleteaspecificSubscriberRegistrationNoContent with default headers values
func NewDeleteaspecificSubscriberRegistrationNoContent() *DeleteaspecificSubscriberRegistrationNoContent {
	return &DeleteaspecificSubscriberRegistrationNoContent{}
}

/*DeleteaspecificSubscriberRegistrationNoContent handles this case with default header values.

DeleteaspecificSubscriberRegistrationNoContent deleteaspecific subscriber registration no content
*/
type DeleteaspecificSubscriberRegistrationNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificSubscriberRegistrationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /subscriberregistrations/{id}][%d] deleteaspecificSubscriberRegistrationNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificSubscriberRegistrationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificSubscriberRegistrationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}