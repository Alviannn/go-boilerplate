package helpers

import (
	"bytes"
	"io"
	"net/http"
)

type httpHelper struct{}

func (h *httpHelper) CloneBody(req *http.Request) (clonedBody []byte, err error) {
	defer req.Body.Close()

	clonedBody, err = io.ReadAll(req.Body)
	if err != nil {
		return
	}

	bodyReader := bytes.NewReader(clonedBody)
	req.Body = io.NopCloser(bodyReader)

	return
}

func (h *httpHelper) CloneRequest(req *http.Request) (clonedReq *http.Request, err error) {
	clonedBody, err := h.CloneBody(req)
	if err != nil {
		return
	}

	ctx := req.Context()
	clonedReq = req.Clone(ctx)
	clonedReq.Body = io.NopCloser(bytes.NewReader(clonedBody))

	return
}
