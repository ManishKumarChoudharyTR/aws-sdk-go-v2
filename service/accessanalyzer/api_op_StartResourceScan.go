// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Starts a scan of the policies applied to the specified resource.
type StartResourceScanInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the analyzer to use to scan the policies applied to the specified
	// resource.
	//
	// AnalyzerArn is a required field
	AnalyzerArn *string `locationName:"analyzerArn" type:"string" required:"true"`

	// The ARN of the resource to scan.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" type:"string" required:"true"`
}

// String returns the string representation
func (s StartResourceScanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartResourceScanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartResourceScanInput"}

	if s.AnalyzerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AnalyzerArn"))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartResourceScanInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AnalyzerArn != nil {
		v := *s.AnalyzerArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "analyzerArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type StartResourceScanOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartResourceScanOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StartResourceScanOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStartResourceScan = "StartResourceScan"

// StartResourceScanRequest returns a request value for making API operation for
// Access Analyzer.
//
// Immediately starts a scan of the policies applied to the specified resource.
//
//    // Example sending a request using StartResourceScanRequest.
//    req := client.StartResourceScanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/accessanalyzer-2019-11-01/StartResourceScan
func (c *Client) StartResourceScanRequest(input *StartResourceScanInput) StartResourceScanRequest {
	op := &aws.Operation{
		Name:       opStartResourceScan,
		HTTPMethod: "POST",
		HTTPPath:   "/resource/scan",
	}

	if input == nil {
		input = &StartResourceScanInput{}
	}

	req := c.newRequest(op, input, &StartResourceScanOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartResourceScanRequest{Request: req, Input: input, Copy: c.StartResourceScanRequest}
}

// StartResourceScanRequest is the request type for the
// StartResourceScan API operation.
type StartResourceScanRequest struct {
	*aws.Request
	Input *StartResourceScanInput
	Copy  func(*StartResourceScanInput) StartResourceScanRequest
}

// Send marshals and sends the StartResourceScan API request.
func (r StartResourceScanRequest) Send(ctx context.Context) (*StartResourceScanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartResourceScanResponse{
		StartResourceScanOutput: r.Request.Data.(*StartResourceScanOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartResourceScanResponse is the response type for the
// StartResourceScan API operation.
type StartResourceScanResponse struct {
	*StartResourceScanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartResourceScan request.
func (r *StartResourceScanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}