package cfn_custom_resource_helper

func (h *Helper) poolEnabled() bool {
	if h.createPoolFn != nil || h.updatePoolFn != nil || h.deletePoolFn != nil {
		return true
	}
	return false
}

func (h *Helper) pollingInit() {

}
