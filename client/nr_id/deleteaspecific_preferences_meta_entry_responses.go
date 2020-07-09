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

// DeleteaspecificPreferencesMetaEntryReader is a Reader for the DeleteaspecificPreferencesMetaEntry structure.
type DeleteaspecificPreferencesMetaEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteaspecificPreferencesMetaEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteaspecificPreferencesMetaEntryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteaspecificPreferencesMetaEntryNoContent creates a DeleteaspecificPreferencesMetaEntryNoContent with default headers values
func NewDeleteaspecificPreferencesMetaEntryNoContent() *DeleteaspecificPreferencesMetaEntryNoContent {
	return &DeleteaspecificPreferencesMetaEntryNoContent{}
}

/*DeleteaspecificPreferencesMetaEntryNoContent handles this case with default header values.

DeleteaspecificPreferencesMetaEntryNoContent deleteaspecific preferences meta entry no content
*/
type DeleteaspecificPreferencesMetaEntryNoContent struct {
	Payload interface{}
}

func (o *DeleteaspecificPreferencesMetaEntryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /preferencesmetaentries/{id}][%d] deleteaspecificPreferencesMetaEntryNoContent  %+v", 204, o.Payload)
}

func (o *DeleteaspecificPreferencesMetaEntryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteaspecificPreferencesMetaEntryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
