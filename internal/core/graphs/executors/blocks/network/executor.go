package network

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/smithshelke/flur/internal/core/graphs/executors"
)

type NetworkBlockExecutor interface {
	executors.Executor
}

type networkBlockExecutor struct {
	ctx    context.Context
	config NetworkBlockExecutorConfig
	io     NetworkBlockExecutorIO
}

func (e *networkBlockExecutor) Run() error {
	payload := []byte(e.io.GetBody())

	client := &http.Client{}
	url := e.io.GetHost() + e.io.GetEndpoint()
	req, err := http.NewRequest(e.io.GetMethod(), url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	for k, v := range e.io.GetHeaders() {
		for i := range v {
			req.Header.Add(k, v[i])
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	output := NetworkBlockOutput{
		StatusCode: res.StatusCode,
		Body:       body,
		Headers:    res.Header,
	}
	e.io.GetOutputHandler()(output)
	return nil
}

func (e *networkBlockExecutor) Stop() error {
	return nil
}

func networkCall(method string) {
}
