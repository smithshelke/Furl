package network

import (
	"github.com/smithshelke/flur/internal/core/graphs/executors"
)

type NetworkBlockOutput struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
}

type NetworkBlockExecutorIO interface {
	executors.IO
	GetHeaders() map[string][]string
	GetEndpoint() string
	GetHost() string
	GetMethod() string
	GetContentType() string
	GetBody() string
	GetParams() map[string]string
	GetOutputHandler() executors.OutputHandler
}

type networkBlockExecutorIO struct {
	headers       map[string][]string
	endpoint      string
	host          string
	method        string
	contentType   string
	body          string
	params        map[string]string
	outputHandler executors.OutputHandler
}

func (i *networkBlockExecutorIO) GetHeaders() map[string][]string {
	return i.headers
}

func (i *networkBlockExecutorIO) GetEndpoint() string {
	return i.endpoint
}

func (i *networkBlockExecutorIO) GetHost() string {
	return i.host
}

func (i *networkBlockExecutorIO) GetMethod() string {
	return i.method
}

func (i *networkBlockExecutorIO) GetContentType() string {
	return i.contentType
}

func (i *networkBlockExecutorIO) GetBody() string {
	return i.body
}

func (i *networkBlockExecutorIO) GetParams() map[string]string {
	return i.params
}

func (i *networkBlockExecutorIO) GetOutputHandler() executors.OutputHandler {
	return i.outputHandler
}

func (i *networkBlockExecutorIO) ValidateInput() error {
	return nil
}

func (i *networkBlockExecutorIO) ValidateOutput() error {
	return nil
}
