// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSAMLProviderRequest
type GetSAMLProviderInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the SAML provider resource object in IAM
	// to get information about.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// SAMLProviderArn is a required field
	SAMLProviderArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSAMLProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSAMLProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSAMLProviderInput"}

	if s.SAMLProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SAMLProviderArn"))
	}
	if s.SAMLProviderArn != nil && len(*s.SAMLProviderArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("SAMLProviderArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetSAMLProvider request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSAMLProviderResponse
type GetSAMLProviderOutput struct {
	_ struct{} `type:"structure"`

	// The date and time when the SAML provider was created.
	CreateDate *time.Time `type:"timestamp"`

	// The XML metadata document that includes information about an identity provider.
	SAMLMetadataDocument *string `min:"1000" type:"string"`

	// The expiration date and time for the SAML provider.
	ValidUntil *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s GetSAMLProviderOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSAMLProvider = "GetSAMLProvider"

// GetSAMLProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Returns the SAML provider metadocument that was uploaded when the IAM SAML
// provider resource object was created or updated.
//
// This operation requires Signature Version 4 (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
//    // Example sending a request using GetSAMLProviderRequest.
//    req := client.GetSAMLProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSAMLProvider
func (c *Client) GetSAMLProviderRequest(input *GetSAMLProviderInput) GetSAMLProviderRequest {
	op := &aws.Operation{
		Name:       opGetSAMLProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSAMLProviderInput{}
	}

	req := c.newRequest(op, input, &GetSAMLProviderOutput{})
	return GetSAMLProviderRequest{Request: req, Input: input, Copy: c.GetSAMLProviderRequest}
}

// GetSAMLProviderRequest is the request type for the
// GetSAMLProvider API operation.
type GetSAMLProviderRequest struct {
	*aws.Request
	Input *GetSAMLProviderInput
	Copy  func(*GetSAMLProviderInput) GetSAMLProviderRequest
}

// Send marshals and sends the GetSAMLProvider API request.
func (r GetSAMLProviderRequest) Send(ctx context.Context) (*GetSAMLProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSAMLProviderResponse{
		GetSAMLProviderOutput: r.Request.Data.(*GetSAMLProviderOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSAMLProviderResponse is the response type for the
// GetSAMLProvider API operation.
type GetSAMLProviderResponse struct {
	*GetSAMLProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSAMLProvider request.
func (r *GetSAMLProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}