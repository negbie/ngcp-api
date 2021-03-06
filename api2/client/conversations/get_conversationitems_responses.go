// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/negbie/ngcp-api/models"
)

// GetConversationitemsReader is a Reader for the GetConversationitems structure.
type GetConversationitemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationitemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationitemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConversationitemsOK creates a GetConversationitemsOK with default headers values
func NewGetConversationitemsOK() *GetConversationitemsOK {
	return &GetConversationitemsOK{}
}

/*GetConversationitemsOK handles this case with default header values.

GetConversationitemsOK get conversationitems o k
*/
type GetConversationitemsOK struct {
	Payload []*models.Conversation
}

func (o *GetConversationitemsOK) Error() string {
	return fmt.Sprintf("[GET /conversations][%d] getConversationitemsOK  %+v", 200, o.Payload)
}

func (o *GetConversationitemsOK) GetPayload() []*models.Conversation {
	return o.Payload
}

func (o *GetConversationitemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
