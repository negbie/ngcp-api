// Code generated by go-swagger; DO NOT EDIT.

package callrecordingstreams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api1/models"
)

// GetCallRecordingStreamitemsReader is a Reader for the GetCallRecordingStreamitems structure.
type GetCallRecordingStreamitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCallRecordingStreamitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCallRecordingStreamitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCallRecordingStreamitemsOK creates a GetCallRecordingStreamitemsOK with default headers values
func NewGetCallRecordingStreamitemsOK() *GetCallRecordingStreamitemsOK {
	return &GetCallRecordingStreamitemsOK{}
}

/*GetCallRecordingStreamitemsOK handles this case with default header values.

GetCallRecordingStreamitemsOK get call recording streamitems o k
*/
type GetCallRecordingStreamitemsOK struct {
	Payload []*models.CallRecordingStreams
}

func (o *GetCallRecordingStreamitemsOK) Error() string {
	return fmt.Sprintf("[GET /callrecordingstreams][%d] getCallRecordingStreamitemsOK  %+v", 200, o.Payload)
}

func (o *GetCallRecordingStreamitemsOK) GetPayload() []*models.CallRecordingStreams {
	return o.Payload
}

func (o *GetCallRecordingStreamitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
