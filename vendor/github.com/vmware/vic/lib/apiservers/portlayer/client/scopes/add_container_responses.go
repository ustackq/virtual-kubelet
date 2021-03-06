package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// AddContainerReader is a Reader for the AddContainer structure.
type AddContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewAddContainerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewAddContainerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddContainerOK creates a AddContainerOK with default headers values
func NewAddContainerOK() *AddContainerOK {
	return &AddContainerOK{}
}

/*AddContainerOK handles this case with default header values.

OK
*/
type AddContainerOK struct {
	Payload string
}

func (o *AddContainerOK) Error() string {
	return fmt.Sprintf("[POST /scopes/{scope}/containers][%d] addContainerOK  %+v", 200, o.Payload)
}

func (o *AddContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddContainerNotFound creates a AddContainerNotFound with default headers values
func NewAddContainerNotFound() *AddContainerNotFound {
	return &AddContainerNotFound{}
}

/*AddContainerNotFound handles this case with default header values.

Not found
*/
type AddContainerNotFound struct {
	Payload *models.Error
}

func (o *AddContainerNotFound) Error() string {
	return fmt.Sprintf("[POST /scopes/{scope}/containers][%d] addContainerNotFound  %+v", 404, o.Payload)
}

func (o *AddContainerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddContainerInternalServerError creates a AddContainerInternalServerError with default headers values
func NewAddContainerInternalServerError() *AddContainerInternalServerError {
	return &AddContainerInternalServerError{}
}

/*AddContainerInternalServerError handles this case with default header values.

error
*/
type AddContainerInternalServerError struct {
	Payload *models.Error
}

func (o *AddContainerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /scopes/{scope}/containers][%d] addContainerInternalServerError  %+v", 500, o.Payload)
}

func (o *AddContainerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
