package cfn_custom_resource_helper

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func Start(createFn, updateFn, deleteFn, createPoolFn, updatePoolFn, deletePoolFn CustomResourceHandlerFunc) {
	crHelper := NewHelper(
		HelperConfig{},
		createFn, updateFn, deleteFn, createPoolFn, updatePoolFn, deletePoolFn,
	)
	lambda.Start(crHelper.HandleCustomResource)
}
