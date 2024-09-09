package cfn_custom_resource_helper

import (
	"context"
)

type HelperConfig struct {
	jsonLogging     bool
	logLevel        string
	botoLevel       string
	pollingInterval int
	sleepOnDelete   int
	sslVerify       bool
}

type CustomResourceHandlerFunc func(
	ctx context.Context, event CrHelperEvent) (physicalResourceID string, data map[string]interface{}, err error)

type Helper struct {
	config                                                                 HelperConfig
	event                                                                  CrHelperEvent
	ctx                                                                    context.Context
	createFn, updateFn, deleteFn, createPoolFn, updatePoolFn, deletePoolFn CustomResourceHandlerFunc
}

func NewHelper(config HelperConfig, createFn, updateFn, deleteFn, createPoolFn, updatePoolFn,
	deletePoolFn CustomResourceHandlerFunc) *Helper {
	return &Helper{
		config:       config,
		createFn:     createFn,
		updateFn:     updateFn,
		deleteFn:     deleteFn,
		createPoolFn: createPoolFn,
		updatePoolFn: updatePoolFn,
		deletePoolFn: deletePoolFn,
	}
}
