// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon Rekognition stream processor that you can use to detect and
// recognize faces in a streaming video. Amazon Rekognition Video is a consumer of
// live video from Amazon Kinesis Video Streams. Amazon Rekognition Video sends
// analysis results to Amazon Kinesis Data Streams. You provide as input a Kinesis
// video stream (Input) and a Kinesis data stream (Output) stream. You also specify
// the face recognition criteria in Settings. For example, the collection
// containing faces that you want to recognize. Use Name to assign an identifier
// for the stream processor. You use Name to manage the stream processor. For
// example, you can start processing the source video by calling
// StartStreamProcessor with the Name field. After you have finished analyzing a
// streaming video, use StopStreamProcessor to stop processing. You can delete the
// stream processor by calling DeleteStreamProcessor.
func (c *Client) CreateStreamProcessor(ctx context.Context, params *CreateStreamProcessorInput, optFns ...func(*Options)) (*CreateStreamProcessorOutput, error) {
	if params == nil {
		params = &CreateStreamProcessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamProcessor", params, optFns, addOperationCreateStreamProcessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamProcessorInput struct {

	// Kinesis video stream stream that provides the source streaming video. If you are
	// using the AWS CLI, the parameter name is StreamProcessorInput.
	//
	// This member is required.
	Input *types.StreamProcessorInput

	// An identifier you assign to the stream processor. You can use Name to manage the
	// stream processor. For example, you can get the current status of the stream
	// processor by calling DescribeStreamProcessor. Name is idempotent.
	//
	// This member is required.
	Name *string

	// Kinesis data stream stream to which Amazon Rekognition Video puts the analysis
	// results. If you are using the AWS CLI, the parameter name is
	// StreamProcessorOutput.
	//
	// This member is required.
	Output *types.StreamProcessorOutput

	// ARN of the IAM role that allows access to the stream processor.
	//
	// This member is required.
	RoleArn *string

	// Face recognition input parameters to be used by the stream processor. Includes
	// the collection to use for face recognition and the face attributes to detect.
	//
	// This member is required.
	Settings *types.StreamProcessorSettings
}

type CreateStreamProcessorOutput struct {

	// ARN for the newly create stream processor.
	StreamProcessorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStreamProcessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStreamProcessor{}, middleware.After)
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
	if err = addOpCreateStreamProcessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamProcessor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStreamProcessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CreateStreamProcessor",
	}
}