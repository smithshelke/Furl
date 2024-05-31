package executors

type OutputHandler func(any)

type input interface {
	ValidateInput() error
}

type output interface {
	ValidateOutput() error
}

type IO interface {
	input
	output
}
