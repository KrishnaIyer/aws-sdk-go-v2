// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListSchemasInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// RegistryName is a required field
	RegistryName *string `location:"uri" locationName:"registryName" type:"string" required:"true"`

	SchemaNamePrefix *string `location:"querystring" locationName:"schemaNamePrefix" type:"string"`
}

// String returns the string representation
func (s ListSchemasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSchemasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSchemasInput"}

	if s.RegistryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistryName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSchemasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.RegistryName != nil {
		v := *s.RegistryName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "registryName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.SchemaNamePrefix != nil {
		v := *s.SchemaNamePrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "schemaNamePrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListSchemasOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `type:"string"`

	Schemas []SchemaSummary `type:"list"`
}

// String returns the string representation
func (s ListSchemasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListSchemasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schemas != nil {
		v := s.Schemas

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Schemas", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListSchemas = "ListSchemas"

// ListSchemasRequest returns a request value for making API operation for
// Schemas.
//
// List the schemas.
//
//    // Example sending a request using ListSchemasRequest.
//    req := client.ListSchemasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/schemas-2019-12-02/ListSchemas
func (c *Client) ListSchemasRequest(input *ListSchemasInput) ListSchemasRequest {
	op := &aws.Operation{
		Name:       opListSchemas,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/registries/name/{registryName}/schemas",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSchemasInput{}
	}

	req := c.newRequest(op, input, &ListSchemasOutput{})
	return ListSchemasRequest{Request: req, Input: input, Copy: c.ListSchemasRequest}
}

// ListSchemasRequest is the request type for the
// ListSchemas API operation.
type ListSchemasRequest struct {
	*aws.Request
	Input *ListSchemasInput
	Copy  func(*ListSchemasInput) ListSchemasRequest
}

// Send marshals and sends the ListSchemas API request.
func (r ListSchemasRequest) Send(ctx context.Context) (*ListSchemasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSchemasResponse{
		ListSchemasOutput: r.Request.Data.(*ListSchemasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSchemasRequestPaginator returns a paginator for ListSchemas.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSchemasRequest(input)
//   p := schemas.NewListSchemasRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSchemasPaginator(req ListSchemasRequest) ListSchemasPaginator {
	return ListSchemasPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSchemasInput
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

// ListSchemasPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSchemasPaginator struct {
	aws.Pager
}

func (p *ListSchemasPaginator) CurrentPage() *ListSchemasOutput {
	return p.Pager.CurrentPage().(*ListSchemasOutput)
}

// ListSchemasResponse is the response type for the
// ListSchemas API operation.
type ListSchemasResponse struct {
	*ListSchemasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSchemas request.
func (r *ListSchemasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}