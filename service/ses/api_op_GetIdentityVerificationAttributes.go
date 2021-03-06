// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Given a list of identities (email addresses and/or domains), returns the
// verification status and (for domain identities) the verification token for each
// identity. The verification status of an email address is "Pending" until the
// email address owner clicks the link within the verification email that Amazon
// SES sent to that address. If the email address owner clicks the link within 24
// hours, the verification status of the email address changes to "Success". If the
// link is not clicked within 24 hours, the verification status changes to
// "Failed." In that case, if you still want to verify the email address, you must
// restart the verification process from the beginning. For domain identities, the
// domain's verification status is "Pending" as Amazon SES searches for the
// required TXT record in the DNS settings of the domain. When Amazon SES detects
// the record, the domain's verification status changes to "Success". If Amazon SES
// is unable to detect the record within 72 hours, the domain's verification status
// changes to "Failed." In that case, if you still want to verify the domain, you
// must restart the verification process from the beginning. This operation is
// throttled at one request per second and can only get verification attributes for
// up to 100 identities at a time.
func (c *Client) GetIdentityVerificationAttributes(ctx context.Context, params *GetIdentityVerificationAttributesInput, optFns ...func(*Options)) (*GetIdentityVerificationAttributesOutput, error) {
	if params == nil {
		params = &GetIdentityVerificationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityVerificationAttributes", params, optFns, addOperationGetIdentityVerificationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityVerificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the Amazon SES verification status of a list of
// identities. For domain identities, this request also returns the verification
// token. For information about verifying identities with Amazon SES, see the
// Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-addresses-and-domains.html).
type GetIdentityVerificationAttributesInput struct {

	// A list of identities.
	//
	// This member is required.
	Identities []string
}

// The Amazon SES verification status of a list of identities. For domain
// identities, this response also contains the verification token.
type GetIdentityVerificationAttributesOutput struct {

	// A map of Identities to IdentityVerificationAttributes objects.
	//
	// This member is required.
	VerificationAttributes map[string]types.IdentityVerificationAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIdentityVerificationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityVerificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityVerificationAttributes{}, middleware.After)
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
	if err = addOpGetIdentityVerificationAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityVerificationAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdentityVerificationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetIdentityVerificationAttributes",
	}
}
