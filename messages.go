package assertions

const (
	shouldUseContext  = "You must provide context instances as the first argument to this assertion."
	shouldUseDuration = "You must provide duration as arguments to this assertion."

	shouldClosedBefore    = "the context has not been closed for the specified period of time (duration: %v)"
	shouldNotClosedBefore = "the context was closed ahead of time (duration: %v)"
)
