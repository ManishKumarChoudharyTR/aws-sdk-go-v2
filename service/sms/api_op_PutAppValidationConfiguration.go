// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a validation configuration for the specified application.
func (c *Client) PutAppValidationConfiguration(ctx context.Context, params *PutAppValidationConfigurationInput, optFns ...func(*Options)) (*PutAppValidationConfigurationOutput, error) {
	if params == nil {
		params = &PutAppValidationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAppValidationConfiguration", params, optFns, addOperationPutAppValidationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAppValidationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppValidationConfigurationInput struct {

	// The ID of the application.
	//
	// This member is required.
	AppId *string

	// The configuration for application validation.
	AppValidationConfigurations []*types.AppValidationConfiguration

	// The configuration for instance validation.
	ServerGroupValidationConfigurations []*types.ServerGroupValidationConfiguration
}

type PutAppValidationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutAppValidationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutAppValidationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAppValidationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutAppValidationConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppValidationConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutAppValidationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "PutAppValidationConfiguration",
	}
}