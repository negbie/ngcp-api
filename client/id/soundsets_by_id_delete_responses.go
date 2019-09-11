// Code generated by go-swagger; DO NOT EDIT.

package id

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SoundsetsByIDDeleteReader is a Reader for the SoundsetsByIDDelete structure.
type SoundsetsByIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoundsetsByIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSoundsetsByIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSoundsetsByIDDeleteNoContent creates a SoundsetsByIDDeleteNoContent with default headers values
func NewSoundsetsByIDDeleteNoContent() *SoundsetsByIDDeleteNoContent {
	return &SoundsetsByIDDeleteNoContent{}
}

/*SoundsetsByIDDeleteNoContent handles this case with default header values.

SoundsetsByIDDeleteNoContent soundsets by Id delete no content
*/
type SoundsetsByIDDeleteNoContent struct {
	Payload interface{}
}

func (o *SoundsetsByIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /soundsets/{id}][%d] soundsetsByIdDeleteNoContent  %+v", 204, o.Payload)
}

func (o *SoundsetsByIDDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *SoundsetsByIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}