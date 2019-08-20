// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateVolumeRequest
type UpdateVolumeInput struct {
	_ struct{} `type:"structure"`

	// The new mount point.
	MountPoint *string `type:"string"`

	// The new name.
	Name *string `type:"string"`

	// The volume ID.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateVolumeOutput
type UpdateVolumeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVolume = "UpdateVolume"

// UpdateVolumeRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Updates an Amazon EBS volume's name or mount point. For more information,
// see Resource Management (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using UpdateVolumeRequest.
//    req := client.UpdateVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateVolume
func (c *Client) UpdateVolumeRequest(input *UpdateVolumeInput) UpdateVolumeRequest {
	op := &aws.Operation{
		Name:       opUpdateVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVolumeInput{}
	}

	req := c.newRequest(op, input, &UpdateVolumeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateVolumeRequest{Request: req, Input: input, Copy: c.UpdateVolumeRequest}
}

// UpdateVolumeRequest is the request type for the
// UpdateVolume API operation.
type UpdateVolumeRequest struct {
	*aws.Request
	Input *UpdateVolumeInput
	Copy  func(*UpdateVolumeInput) UpdateVolumeRequest
}

// Send marshals and sends the UpdateVolume API request.
func (r UpdateVolumeRequest) Send(ctx context.Context) (*UpdateVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVolumeResponse{
		UpdateVolumeOutput: r.Request.Data.(*UpdateVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVolumeResponse is the response type for the
// UpdateVolume API operation.
type UpdateVolumeResponse struct {
	*UpdateVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVolume request.
func (r *UpdateVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}