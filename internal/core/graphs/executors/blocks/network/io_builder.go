package network

import "github.com/smithshelke/flur/internal/core/graphs/executors"

type NetworkBlockExecutorIOBuilder struct {
	headers       map[string][]string
	endpoint      string
	host          string
	method        string
	contentType   string
	body          string
	params        map[string]string
	outputHandler executors.OutputHandler
}

func NewNetworkBlockExecutorIOBuilder() *NetworkBlockExecutorIOBuilder {
	return &NetworkBlockExecutorIOBuilder{}
}

func (b *NetworkBlockExecutorIOBuilder) AddHeaders(headers map[string][]string) *NetworkBlockExecutorIOBuilder {
	b.headers = headers
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddEndpoint(endpoint string) *NetworkBlockExecutorIOBuilder {
	b.endpoint = endpoint
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddHost(host string) *NetworkBlockExecutorIOBuilder {
	b.host = host
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddMethod(method string) *NetworkBlockExecutorIOBuilder {
	b.method = method
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddContentType(contentType string) *NetworkBlockExecutorIOBuilder {
	b.contentType = contentType
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddBody(body string) *NetworkBlockExecutorIOBuilder {
	b.body = body
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddParams(params map[string]string) *NetworkBlockExecutorIOBuilder {
	b.params = params
	return b
}

func (b *NetworkBlockExecutorIOBuilder) AddOutputHandler(handler executors.OutputHandler) *NetworkBlockExecutorIOBuilder {
	b.outputHandler = handler
	return b
}

func (b *NetworkBlockExecutorIOBuilder) Build() (NetworkBlockExecutorIO, error) {
	io := &networkBlockExecutorIO{
		headers:       b.headers,
		endpoint:      b.endpoint,
		host:          b.host,
		method:        b.method,
		contentType:   b.contentType,
		body:          b.body,
		params:        b.params,
		outputHandler: b.outputHandler,
	}
	return io, io.ValidateInput()
}
