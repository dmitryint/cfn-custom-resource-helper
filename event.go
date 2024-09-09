package cfn_custom_resource_helper

import (
	"context"
	"github.com/aws/aws-lambda-go/cfn"
)

type CrHelperEvent struct {
	RequestType           cfn.RequestType        `json:"RequestType"`
	RequestID             string                 `json:"RequestId"`
	ResponseURL           string                 `json:"ResponseURL"`
	ResourceType          string                 `json:"ResourceType"`
	PhysicalResourceID    string                 `json:"PhysicalResourceId,omitempty"`
	LogicalResourceID     string                 `json:"LogicalResourceId"`
	StackID               string                 `json:"StackId"`
	ResourceProperties    map[string]interface{} `json:"ResourceProperties"`
	OldResourceProperties map[string]interface{} `json:"OldResourceProperties,omitempty"`

	CrHelperPoll       bool                   `json:"CrHelperPoll"`
	CrHelperData       map[string]interface{} `json:"CrHelperData"`
	CrHelperRule       string                 `json:"CrHelperRule"`
	CrHelperPermission string                 `json:"CrHelperPermission"`
}

func (h *Helper) receiveEvent(ctx context.Context, event CrHelperEvent) {
	h.event = event
	h.ctx = ctx
}

func (h *Helper) getLambdaEvent() *cfn.Event {
	return &cfn.Event{
		RequestType:           h.event.RequestType,
		RequestID:             h.event.RequestID,
		ResponseURL:           h.event.ResponseURL,
		ResourceType:          h.event.ResourceType,
		PhysicalResourceID:    h.event.PhysicalResourceID,
		LogicalResourceID:     h.event.LogicalResourceID,
		StackID:               h.event.StackID,
		ResourceProperties:    h.event.ResourceProperties,
		OldResourceProperties: h.event.OldResourceProperties,
	}
}
