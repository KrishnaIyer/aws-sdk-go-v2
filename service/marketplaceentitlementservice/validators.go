// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplaceentitlementservice

import (
	"context"
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpGetEntitlements struct {
}

func (*validateOpGetEntitlements) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEntitlements) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEntitlementsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEntitlementsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpGetEntitlementsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEntitlements{}, middleware.After)
}

func validateOpGetEntitlementsInput(v *GetEntitlementsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEntitlementsInput"}
	if v.ProductCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProductCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
