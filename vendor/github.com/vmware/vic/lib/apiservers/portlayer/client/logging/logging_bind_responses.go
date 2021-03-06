package logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// LoggingBindReader is a Reader for the LoggingBind structure.
type LoggingBindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoggingBindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoggingBindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewLoggingBindNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewLoggingBindInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoggingBindOK creates a LoggingBindOK with default headers values
func NewLoggingBindOK() *LoggingBindOK {
	return &LoggingBindOK{}
}

/*LoggingBindOK handles this case with default header values.

OK
*/
type LoggingBindOK struct {
	Payload *models.LoggingBindResponse
}

func (o *LoggingBindOK) Error() string {
	return fmt.Sprintf("[POST /logging/binding][%d] loggingBindOK  %+v", 200, o.Payload)
}

func (o *LoggingBindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoggingBindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoggingBindNotFound creates a LoggingBindNotFound with default headers values
func NewLoggingBindNotFound() *LoggingBindNotFound {
	return &LoggingBindNotFound{}
}

/*LoggingBindNotFound handles this case with default header values.

VirtualDevice not found
*/
type LoggingBindNotFound struct {
	Payload *models.Error
}

func (o *LoggingBindNotFound) Error() string {
	return fmt.Sprintf("[POST /logging/binding][%d] loggingBindNotFound  %+v", 404, o.Payload)
}

func (o *LoggingBindNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoggingBindInternalServerError creates a LoggingBindInternalServerError with default headers values
func NewLoggingBindInternalServerError() *LoggingBindInternalServerError {
	return &LoggingBindInternalServerError{}
}

/*LoggingBindInternalServerError handles this case with default header values.

Connecting VirtualDevice failed
*/
type LoggingBindInternalServerError struct {
	Payload *models.Error
}

func (o *LoggingBindInternalServerError) Error() string {
	return fmt.Sprintf("[POST /logging/binding][%d] loggingBindInternalServerError  %+v", 500, o.Payload)
}

func (o *LoggingBindInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
