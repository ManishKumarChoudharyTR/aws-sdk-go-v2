// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a proxy session on the specified Amazon Chime Voice Connector for the
// specified participant phone numbers.
func (c *Client) CreateProxySession(ctx context.Context, params *CreateProxySessionInput, optFns ...func(*Options)) (*CreateProxySessionOutput, error) {
	if params == nil {
		params = &CreateProxySessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProxySession", params, optFns, addOperationCreateProxySessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProxySessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProxySessionInput struct {

	// The proxy session capabilities.
	//
	// This member is required.
	Capabilities []types.Capability

	// The participant phone numbers.
	//
	// This member is required.
	ParticipantPhoneNumbers []*string

	// The Amazon Chime voice connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The number of minutes allowed for the proxy session.
	ExpiryMinutes *int32

	// The preference for matching the country or area code of the proxy phone number
	// with that of the first participant.
	GeoMatchLevel types.GeoMatchLevel

	// The country and area code for the proxy phone number.
	GeoMatchParams *types.GeoMatchParams

	// The name of the proxy session.
	Name *string

	// The preference for proxy phone number reuse, or stickiness, between the same
	// participants across sessions.
	NumberSelectionBehavior types.NumberSelectionBehavior
}

type CreateProxySessionOutput struct {

	// The proxy session details.
	ProxySession *types.ProxySession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateProxySessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProxySession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProxySession{}, middleware.After)
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
	if err = addOpCreateProxySessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProxySession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProxySession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateProxySession",
	}
}