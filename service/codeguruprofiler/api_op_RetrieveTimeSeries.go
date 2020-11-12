// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) RetrieveTimeSeries(ctx context.Context, params *RetrieveTimeSeriesInput, optFns ...func(*Options)) (*RetrieveTimeSeriesOutput, error) {
	if params == nil {
		params = &RetrieveTimeSeriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetrieveTimeSeries", params, optFns, addOperationRetrieveTimeSeriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveTimeSeriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetrieveTimeSeriesInput struct {
	ProfilingGroupName *string

	EndTime *time.Time

	FrameMetrics []*types.FrameMetric

	Period *string

	StartTime *time.Time

	TargetResolution types.AggregationPeriod
}

type RetrieveTimeSeriesOutput struct {
	Data [][]*float64

	EndTime *time.Time

	EndTimes []*time.Time

	FrameMetrics []*types.FrameMetric

	Resolution types.AggregationPeriod

	StartTime *time.Time

	UnprocessedEndTimes map[string][]*time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRetrieveTimeSeriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRetrieveTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRetrieveTimeSeries{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRetrieveTimeSeriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}