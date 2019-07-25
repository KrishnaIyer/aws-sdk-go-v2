// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorTargetRequest
type DeleteTrafficMirrorTargetInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the Traffic Mirror target.
	//
	// TrafficMirrorTargetId is a required field
	TrafficMirrorTargetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrafficMirrorTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrafficMirrorTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrafficMirrorTargetInput"}

	if s.TrafficMirrorTargetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorTargetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorTargetResult
type DeleteTrafficMirrorTargetOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted Traffic Mirror target.
	TrafficMirrorTargetId *string `locationName:"trafficMirrorTargetId" type:"string"`
}

// String returns the string representation
func (s DeleteTrafficMirrorTargetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTrafficMirrorTarget = "DeleteTrafficMirrorTarget"

// DeleteTrafficMirrorTargetRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified Traffic Mirror target.
//
// You cannot delete a Traffic Mirror target that is in use by a Traffic Mirror
// session.
//
//    // Example sending a request using DeleteTrafficMirrorTargetRequest.
//    req := client.DeleteTrafficMirrorTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteTrafficMirrorTarget
func (c *Client) DeleteTrafficMirrorTargetRequest(input *DeleteTrafficMirrorTargetInput) DeleteTrafficMirrorTargetRequest {
	op := &aws.Operation{
		Name:       opDeleteTrafficMirrorTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTrafficMirrorTargetInput{}
	}

	req := c.newRequest(op, input, &DeleteTrafficMirrorTargetOutput{})
	return DeleteTrafficMirrorTargetRequest{Request: req, Input: input, Copy: c.DeleteTrafficMirrorTargetRequest}
}

// DeleteTrafficMirrorTargetRequest is the request type for the
// DeleteTrafficMirrorTarget API operation.
type DeleteTrafficMirrorTargetRequest struct {
	*aws.Request
	Input *DeleteTrafficMirrorTargetInput
	Copy  func(*DeleteTrafficMirrorTargetInput) DeleteTrafficMirrorTargetRequest
}

// Send marshals and sends the DeleteTrafficMirrorTarget API request.
func (r DeleteTrafficMirrorTargetRequest) Send(ctx context.Context) (*DeleteTrafficMirrorTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTrafficMirrorTargetResponse{
		DeleteTrafficMirrorTargetOutput: r.Request.Data.(*DeleteTrafficMirrorTargetOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTrafficMirrorTargetResponse is the response type for the
// DeleteTrafficMirrorTarget API operation.
type DeleteTrafficMirrorTargetResponse struct {
	*DeleteTrafficMirrorTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTrafficMirrorTarget request.
func (r *DeleteTrafficMirrorTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}