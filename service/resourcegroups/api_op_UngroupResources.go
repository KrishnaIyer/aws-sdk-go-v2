// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroups

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the specified resources from the specified group.
func (c *Client) UngroupResources(ctx context.Context, params *UngroupResourcesInput, optFns ...func(*Options)) (*UngroupResourcesOutput, error) {
	if params == nil {
		params = &UngroupResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UngroupResources", params, optFns, addOperationUngroupResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UngroupResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UngroupResourcesInput struct {

	// The name or the ARN of the resource group from which to remove the resources.
	//
	// This member is required.
	Group *string

	// The ARNs of the resources to be removed from the group.
	//
	// This member is required.
	ResourceArns []string
}

type UngroupResourcesOutput struct {

	// The resources that failed to be removed from the group.
	Failed []types.FailedResource

	// The ARNs of the resources that were successfully removed from the group.
	Succeeded []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUngroupResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUngroupResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUngroupResources{}, middleware.After)
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
	if err = addOpUngroupResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUngroupResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUngroupResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-groups",
		OperationName: "UngroupResources",
	}
}
