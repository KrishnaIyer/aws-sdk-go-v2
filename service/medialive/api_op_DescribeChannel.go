// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets details about a channel
func (c *Client) DescribeChannel(ctx context.Context, params *DescribeChannelInput, optFns ...func(*Options)) (*DescribeChannelOutput, error) {
	if params == nil {
		params = &DescribeChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeChannel", params, optFns, addOperationDescribeChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeChannelRequest
type DescribeChannelInput struct {

	// channel ID
	//
	// This member is required.
	ChannelId *string
}

// Placeholder documentation for DescribeChannelResponse
type DescribeChannelOutput struct {

	// The unique arn of the channel.
	Arn *string

	// Specification of CDI inputs for this channel
	CdiInputSpecification *types.CdiInputSpecification

	// The class for this channel. STANDARD for a channel with two pipelines or
	// SINGLE_PIPELINE for a channel with one pipeline.
	ChannelClass types.ChannelClass

	// A list of destinations of the channel. For UDP outputs, there is one destination
	// per output. For other types (HLS, for example), there is one destination per
	// packager.
	Destinations []types.OutputDestination

	// The endpoints where outgoing connections initiate from
	EgressEndpoints []types.ChannelEgressEndpoint

	// Encoder Settings
	EncoderSettings *types.EncoderSettings

	// The unique id of the channel.
	Id *string

	// List of input attachments for channel.
	InputAttachments []types.InputAttachment

	// Specification of network and file inputs for this channel
	InputSpecification *types.InputSpecification

	// The log level being written to CloudWatch Logs.
	LogLevel types.LogLevel

	// The name of the channel. (user-mutable)
	Name *string

	// Runtime details for the pipelines of a running channel.
	PipelineDetails []types.PipelineDetail

	// The number of currently healthy pipelines.
	PipelinesRunningCount int32

	// The Amazon Resource Name (ARN) of the role assumed when running the Channel.
	RoleArn *string

	// Placeholder documentation for ChannelState
	State types.ChannelState

	// A collection of key-value pairs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeChannel{}, middleware.After)
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
	if err = addOpDescribeChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeChannel",
	}
}
