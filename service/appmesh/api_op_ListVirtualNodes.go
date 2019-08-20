// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualNodesInput
type ListVirtualNodesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results returned by ListVirtualNodes in paginated output.
	// When you use this parameter, ListVirtualNodes returns only limit results
	// in a single page along with a nextToken response element. You can see the
	// remaining results of the initial request by sending another ListVirtualNodes
	// request with the returned nextToken value. This value can be between 1 and
	// 100. If you don't use this parameter, ListVirtualNodes returns up to 100
	// results and a nextToken value if applicable.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The name of the service mesh to list virtual nodes in.
	//
	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// The nextToken value returned from a previous paginated ListVirtualNodes request
	// where limit was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListVirtualNodesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVirtualNodesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVirtualNodesInput"}
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
func (s ListVirtualNodesInput) MarshalFields(e protocol.FieldEncoder) error {
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualNodesOutput
type ListVirtualNodesOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListVirtualNodes request. When
	// the results of a ListVirtualNodes request exceed limit, you can use this
	// value to retrieve the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of existing virtual nodes for the specified service mesh.
	//
	// VirtualNodes is a required field
	VirtualNodes []VirtualNodeRef `locationName:"virtualNodes" type:"list" required:"true"`
}

// String returns the string representation
func (s ListVirtualNodesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualNodesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualNodes != nil {
		v := s.VirtualNodes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "virtualNodes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVirtualNodes = "ListVirtualNodes"

// ListVirtualNodesRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing virtual nodes.
//
//    // Example sending a request using ListVirtualNodesRequest.
//    req := client.ListVirtualNodesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualNodes
func (c *Client) ListVirtualNodesRequest(input *ListVirtualNodesInput) ListVirtualNodesRequest {
	op := &aws.Operation{
		Name:       opListVirtualNodes,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualNodes",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVirtualNodesInput{}
	}

	req := c.newRequest(op, input, &ListVirtualNodesOutput{})
	return ListVirtualNodesRequest{Request: req, Input: input, Copy: c.ListVirtualNodesRequest}
}

// ListVirtualNodesRequest is the request type for the
// ListVirtualNodes API operation.
type ListVirtualNodesRequest struct {
	*aws.Request
	Input *ListVirtualNodesInput
	Copy  func(*ListVirtualNodesInput) ListVirtualNodesRequest
}

// Send marshals and sends the ListVirtualNodes API request.
func (r ListVirtualNodesRequest) Send(ctx context.Context) (*ListVirtualNodesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVirtualNodesResponse{
		ListVirtualNodesOutput: r.Request.Data.(*ListVirtualNodesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVirtualNodesRequestPaginator returns a paginator for ListVirtualNodes.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVirtualNodesRequest(input)
//   p := appmesh.NewListVirtualNodesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVirtualNodesPaginator(req ListVirtualNodesRequest) ListVirtualNodesPaginator {
	return ListVirtualNodesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVirtualNodesInput
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

// ListVirtualNodesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVirtualNodesPaginator struct {
	aws.Pager
}

func (p *ListVirtualNodesPaginator) CurrentPage() *ListVirtualNodesOutput {
	return p.Pager.CurrentPage().(*ListVirtualNodesOutput)
}

// ListVirtualNodesResponse is the response type for the
// ListVirtualNodes API operation.
type ListVirtualNodesResponse struct {
	*ListVirtualNodesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVirtualNodes request.
func (r *ListVirtualNodesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}