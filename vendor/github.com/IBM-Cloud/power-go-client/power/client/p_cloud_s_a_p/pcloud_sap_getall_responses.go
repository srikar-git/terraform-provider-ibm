// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_s_a_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudSapGetallReader is a Reader for the PcloudSapGetall structure.
type PcloudSapGetallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudSapGetallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudSapGetallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudSapGetallBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudSapGetallInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudSapGetallOK creates a PcloudSapGetallOK with default headers values
func NewPcloudSapGetallOK() *PcloudSapGetallOK {
	return &PcloudSapGetallOK{}
}

/*PcloudSapGetallOK handles this case with default header values.

OK
*/
type PcloudSapGetallOK struct {
	Payload *models.SAPProfiles
}

func (o *PcloudSapGetallOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallOK  %+v", 200, o.Payload)
}

func (o *PcloudSapGetallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SAPProfiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetallBadRequest creates a PcloudSapGetallBadRequest with default headers values
func NewPcloudSapGetallBadRequest() *PcloudSapGetallBadRequest {
	return &PcloudSapGetallBadRequest{}
}

/*PcloudSapGetallBadRequest handles this case with default header values.

Bad Request
*/
type PcloudSapGetallBadRequest struct {
	Payload *models.Error
}

func (o *PcloudSapGetallBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudSapGetallBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudSapGetallInternalServerError creates a PcloudSapGetallInternalServerError with default headers values
func NewPcloudSapGetallInternalServerError() *PcloudSapGetallInternalServerError {
	return &PcloudSapGetallInternalServerError{}
}

/*PcloudSapGetallInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudSapGetallInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudSapGetallInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/sap][%d] pcloudSapGetallInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudSapGetallInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
