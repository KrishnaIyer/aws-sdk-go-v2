// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DeleteCertificateAuthorityRequest
type DeleteCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority.
	// This must have the following form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012.
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// The number of days to make a CA restorable after it has been deleted. This
	// can be anywhere from 7 to 30 days, with 30 being the default.
	PermanentDeletionTimeInDays *int64 `min:"7" type:"integer"`
}

// String returns the string representation
func (s DeleteCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCertificateAuthorityInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}
	if s.PermanentDeletionTimeInDays != nil && *s.PermanentDeletionTimeInDays < 7 {
		invalidParams.Add(aws.NewErrParamMinValue("PermanentDeletionTimeInDays", 7))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DeleteCertificateAuthorityOutput
type DeleteCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCertificateAuthority = "DeleteCertificateAuthority"

// DeleteCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Deletes a private certificate authority (CA). You must provide the ARN (Amazon
// Resource Name) of the private CA that you want to delete. You can find the
// ARN by calling the ListCertificateAuthorities operation. Before you can delete
// a CA, you must disable it. Call the UpdateCertificateAuthority operation
// and set the CertificateAuthorityStatus parameter to DISABLED.
//
// Additionally, you can delete a CA if you are waiting for it to be created
// (the Status field of the CertificateAuthority is CREATING). You can also
// delete it if the CA has been created but you haven't yet imported the signed
// certificate (the Status is PENDING_CERTIFICATE) into ACM PCA.
//
// If the CA is in one of the previously mentioned states and you call DeleteCertificateAuthority,
// the CA's status changes to DELETED. However, the CA won't be permanently
// deleted until the restoration period has passed. By default, if you do not
// set the PermanentDeletionTimeInDays parameter, the CA remains restorable
// for 30 days. You can set the parameter from 7 to 30 days. The DescribeCertificateAuthority
// operation returns the time remaining in the restoration window of a Private
// CA in the DELETED state. To restore an eligible CA, call the RestoreCertificateAuthority
// operation.
//
//    // Example sending a request using DeleteCertificateAuthorityRequest.
//    req := client.DeleteCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DeleteCertificateAuthority
func (c *Client) DeleteCertificateAuthorityRequest(input *DeleteCertificateAuthorityInput) DeleteCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opDeleteCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &DeleteCertificateAuthorityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCertificateAuthorityRequest{Request: req, Input: input, Copy: c.DeleteCertificateAuthorityRequest}
}

// DeleteCertificateAuthorityRequest is the request type for the
// DeleteCertificateAuthority API operation.
type DeleteCertificateAuthorityRequest struct {
	*aws.Request
	Input *DeleteCertificateAuthorityInput
	Copy  func(*DeleteCertificateAuthorityInput) DeleteCertificateAuthorityRequest
}

// Send marshals and sends the DeleteCertificateAuthority API request.
func (r DeleteCertificateAuthorityRequest) Send(ctx context.Context) (*DeleteCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCertificateAuthorityResponse{
		DeleteCertificateAuthorityOutput: r.Request.Data.(*DeleteCertificateAuthorityOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCertificateAuthorityResponse is the response type for the
// DeleteCertificateAuthority API operation.
type DeleteCertificateAuthorityResponse struct {
	*DeleteCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCertificateAuthority request.
func (r *DeleteCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}