package cfn_custom_resource_helper

import (
	"context"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambdacontext"
	log "github.com/go-pkgz/lgr"
)

func (h *Helper) HandleCustomResource(ctx context.Context, event CrHelperEvent) (physicalResourceID string, err error) {
	log.Printf("[DEBUG] Start")
	if h.poolEnabled() {
		log.Printf("[DEBUG] pool enabled")
	}

	h.receiveEvent(ctx, event)
	log.Printf("[DEBUG] RequestType: %s", h.event.RequestType)
	switch h.event.RequestType {
	case cfn.RequestCreate:
		return h.wrapCrh(h.createFn)
	case cfn.RequestUpdate:
		return h.wrapCrh(h.updateFn)
	case cfn.RequestDelete:
		return h.wrapCrh(h.deleteFn)
	}

	return
}

func (h *Helper) wrapCrh(f CustomResourceHandlerFunc) (reason string, err error) {
	r := cfn.NewResponse(h.getLambdaEvent())

	funcDidPanic := true
	defer func() {
		if funcDidPanic {
			r.Status = cfn.StatusFailed
			r.Reason = "Function panicked, see log stream for details"
			// FIXME: something should be done if an error is returned here
			_ = r.Send()
		}
	}()

	r.PhysicalResourceID, r.Data, err = f(h.ctx, h.event)
	funcDidPanic = false

	if err != nil {
		r.Status = cfn.StatusFailed
		r.Reason = err.Error()
		log.Printf("[ERROR] sending status failed: %s", r.Reason)
	} else {
		r.Status = cfn.StatusSuccess

		if r.PhysicalResourceID == "" {
			log.Printf("[INFO] PhysicalResourceID must exist on creation, copying Log Stream name")
			r.PhysicalResourceID = lambdacontext.LogStreamName
		}
	}

	err = r.Send()
	if err != nil {
		reason = err.Error()
	}

	return
}
