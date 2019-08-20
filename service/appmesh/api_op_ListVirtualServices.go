// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualServicesInput
type ListVirtualServicesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results returned by ListVirtualServices in paginated
	// output. When you use this parameter, ListVirtualServices returns only limit
	// results in a single page along with a nextToken response element. You can
	// see the remaining results of the initial request by sending another ListVirtualServices
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListVirtualServices returns up to 100
	// results and a nextToken value if applicable.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The name of the service mesh to list virtual services in.
	//
	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// The nextToken value returned from a previous paginated ListVirtualServices
	// request where limit was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListVirtualServicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVirtualServicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVirtualServicesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualServicesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualServicesOutput
type ListVirtualServicesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListVirtualServices request. When
	// the results of a ListVirtualServices request exceed limit, you can use this
	// value to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of existing virtual services for the specified service mesh.
	//
	// VirtualServices is a required field
	VirtualServices []VirtualServiceRef `locationName:"virtualServices" type:"list" required:"true"`
}

// String returns the string representation
func (s ListVirtualServicesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualServicesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualServices != nil {
		v := s.VirtualServices

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "virtualServices", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVirtualServices = "ListVirtualServices"

// ListVirtualServicesRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing virtual services in a service mesh.
//
//    // Example sending a request using ListVirtualServicesRequest.
//    req := client.ListVirtualServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualServices
func (c *Client) ListVirtualServicesRequest(input *ListVirtualServicesInput) ListVirtualServicesRequest {
	op := &aws.Operation{
		Name:       opListVirtualServices,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualServices",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVirtualServicesInput{}
	}

	req := c.newRequest(op, input, &ListVirtualServicesOutput{})
	return ListVirtualServicesRequest{Request: req, Input: input, Copy: c.ListVirtualServicesRequest}
}

// ListVirtualServicesRequest is the request type for the
// ListVirtualServices API operation.
type ListVirtualServicesRequest struct {
	*aws.Request
	Input *ListVirtualServicesInput
	Copy  func(*ListVirtualServicesInput) ListVirtualServicesRequest
}

// Send marshals and sends the ListVirtualServices API request.
func (r ListVirtualServicesRequest) Send(ctx context.Context) (*ListVirtualServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVirtualServicesResponse{
		ListVirtualServicesOutput: r.Request.Data.(*ListVirtualServicesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVirtualServicesRequestPaginator returns a paginator for ListVirtualServices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVirtualServicesRequest(input)
//   p := appmesh.NewListVirtualServicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVirtualServicesPaginator(req ListVirtualServicesRequest) ListVirtualServicesPaginator {
	return ListVirtualServicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVirtualServicesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListVirtualServicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVirtualServicesPaginator struct {
	aws.Pager
}

func (p *ListVirtualServicesPaginator) CurrentPage() *ListVirtualServicesOutput {
	return p.Pager.CurrentPage().(*ListVirtualServicesOutput)
}

// ListVirtualServicesResponse is the response type for the
// ListVirtualServices API operation.
type ListVirtualServicesResponse struct {
	*ListVirtualServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVirtualServices request.
func (r *ListVirtualServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}