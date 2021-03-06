// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all the tagged or previously tagged resources that are located in the
// specified Region for the AWS account. Depending on what information you want
// returned, you can also specify the following:
//
// * Filters that specify what tags
// and resource types you want returned. The response includes all tags that are
// associated with the requested resources.
//
// * Information about compliance with
// the account's effective tag policy. For more information on tag policies, see
// Tag Policies
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
// in the AWS Organizations User Guide.
//
// You can check the PaginationToken response
// parameter to determine if a query is complete. Queries occasionally return fewer
// results on a page than allowed. The PaginationToken response parameter value is
// null only when there are no more results to display.
func (c *Client) GetResources(ctx context.Context, params *GetResourcesInput, optFns ...func(*Options)) (*GetResourcesOutput, error) {
	if params == nil {
		params = &GetResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResources", params, optFns, addOperationGetResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourcesInput struct {

	// Specifies whether to exclude resources that are compliant with the tag policy.
	// Set this to true if you are interested in retrieving information on noncompliant
	// resources only. You can use this parameter only if the IncludeComplianceDetails
	// parameter is also set to true.
	ExcludeCompliantResources *bool

	// Specifies whether to include details regarding the compliance with the effective
	// tag policy. Set this to true to determine whether resources are compliant with
	// the tag policy and to get details.
	IncludeComplianceDetails *bool

	// A string that indicates that additional data is available. Leave this value
	// empty for your initial request. If the response includes a PaginationToken, use
	// that string for this value to request an additional page of data.
	PaginationToken *string

	// The constraints on the resources that you want returned. The format of each
	// resource type is service[:resourceType]. For example, specifying a resource type
	// of ec2 returns all Amazon EC2 resources (which includes EC2 instances).
	// Specifying a resource type of ec2:instance returns only EC2 instances. The
	// string for each service name and resource type is the same as that embedded in a
	// resource's Amazon Resource Name (ARN). Consult the AWS General Reference for the
	// following:
	//
	// * For a list of service name strings, see AWS Service Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces).
	//
	// *
	// For resource type strings, see Example ARNs
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arns-syntax).
	//
	// *
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces
	// (http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	//
	// You
	// can specify multiple resource types by using an array. The array can include up
	// to 100 items. Note that the length constraint requirement applies to each
	// resource type filter.
	ResourceTypeFilters []string

	// A limit that restricts the number of resources returned by GetResources in
	// paginated output. You can set ResourcesPerPage to a minimum of 1 item and the
	// maximum of 100 items.
	ResourcesPerPage *int32

	// A list of TagFilters (keys and values). Each TagFilter specified must contain a
	// key with values as optional. A request can include up to 50 keys, and each key
	// can include up to 20 values. Note the following when deciding how to use
	// TagFilters:
	//
	// * If you do specify a TagFilter, the response returns only those
	// resources that are currently associated with the specified tag.
	//
	// * If you don't
	// specify a TagFilter, the response includes all resources that were ever
	// associated with tags. Resources that currently don't have associated tags are
	// shown with an empty tag set, like this: "Tags": [].
	//
	// * If you specify more than
	// one filter in a single request, the response returns only those resources that
	// satisfy all specified filters.
	//
	// * If you specify a filter that contains more
	// than one value for a key, the response returns resources that match any of the
	// specified values for that key.
	//
	// * If you don't specify any values for a key, the
	// response returns resources that are tagged with that key irrespective of the
	// value. For example, for filters: filter1 = {key1, {value1}}, filter2 = {key2,
	// {value2,value3,value4}} , filter3 = {key3}:
	//
	// * GetResources( {filter1} ) returns
	// resources tagged with key1=value1
	//
	// * GetResources( {filter2} ) returns resources
	// tagged with key2=value2 or key2=value3 or key2=value4
	//
	// * GetResources( {filter3}
	// ) returns resources tagged with any tag containing key3 as its tag key,
	// irrespective of its value
	//
	// * GetResources( {filter1,filter2,filter3} ) returns
	// resources tagged with ( key1=value1) and ( key2=value2 or key2=value3 or
	// key2=value4) and (key3, irrespective of the value)
	TagFilters []types.TagFilter

	// AWS recommends using ResourcesPerPage instead of this parameter. A limit that
	// restricts the number of tags (key and value pairs) returned by GetResources in
	// paginated output. A resource with no tags is counted as having one tag (one key
	// and value pair). GetResources does not split a resource and its associated tags
	// across pages. If the specified TagsPerPage would cause such a break, a
	// PaginationToken is returned in place of the affected resource and its tags. Use
	// that token in another request to get the remaining data. For example, if you
	// specify a TagsPerPage of 100 and the account has 22 resources with 10 tags each
	// (meaning that each resource has 10 key and value pairs), the output will consist
	// of three pages. The first page displays the first 10 resources, each with its 10
	// tags. The second page displays the next 10 resources, each with its 10 tags. The
	// third page displays the remaining 2 resources, each with its 10 tags. You can
	// set TagsPerPage to a minimum of 100 items and the maximum of 500 items.
	TagsPerPage *int32
}

type GetResourcesOutput struct {

	// A string that indicates that the response contains more data than can be
	// returned in a single response. To receive additional data, specify this string
	// for the PaginationToken value in a subsequent request.
	PaginationToken *string

	// A list of resource ARNs and the tags (keys and values) associated with each.
	ResourceTagMappingList []types.ResourceTagMapping

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResources{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResources(options.Region), middleware.Before); err != nil {
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

// GetResourcesAPIClient is a client that implements the GetResources operation.
type GetResourcesAPIClient interface {
	GetResources(context.Context, *GetResourcesInput, ...func(*Options)) (*GetResourcesOutput, error)
}

var _ GetResourcesAPIClient = (*Client)(nil)

// GetResourcesPaginatorOptions is the paginator options for GetResources
type GetResourcesPaginatorOptions struct {
	// A limit that restricts the number of resources returned by GetResources in
	// paginated output. You can set ResourcesPerPage to a minimum of 1 item and the
	// maximum of 100 items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourcesPaginator is a paginator for GetResources
type GetResourcesPaginator struct {
	options   GetResourcesPaginatorOptions
	client    GetResourcesAPIClient
	params    *GetResourcesInput
	nextToken *string
	firstPage bool
}

// NewGetResourcesPaginator returns a new GetResourcesPaginator
func NewGetResourcesPaginator(client GetResourcesAPIClient, params *GetResourcesInput, optFns ...func(*GetResourcesPaginatorOptions)) *GetResourcesPaginator {
	options := GetResourcesPaginatorOptions{}
	if params.ResourcesPerPage != nil {
		options.Limit = *params.ResourcesPerPage
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &GetResourcesInput{}
	}

	return &GetResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetResources page.
func (p *GetResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.ResourcesPerPage = limit

	result, err := p.client.GetResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "GetResources",
	}
}
