// Package jsonrpc provides JSON RPC utilities for serialisation of AWS
// requests and responses.
package jsonrpc

//go:generate go run ../../fixtures/protocol/generate.go ../../fixtures/protocol/input/json.json build_test.go
//go:generate go run ../../fixtures/protocol/generate.go ../../fixtures/protocol/output/json.json unmarshal_test.go

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/internal/protocol/json/jsonutil"
)

var emptyJSON = []byte("{}")

// Build builds a JSON payload for a JSON RPC request.
func Build(req *service.Request) {
	var buf []byte
	var err error
	if req.ParamsFilled() {
		buf, err = jsonutil.BuildJSON(req.Params)
		if err != nil {
			req.Error = awserr.New("SerializationError", "failed encoding JSON RPC request", err)
			return
		}
	} else {
		buf = emptyJSON
	}

	if req.Service.TargetPrefix != "" || string(buf) != "{}" {
		req.SetBufferBody(buf)
	}

	if req.Service.TargetPrefix != "" {
		target := req.Service.TargetPrefix + "." + req.Operation.Name
		req.HTTPRequest.Header.Add("X-Amz-Target", target)
	}
	if req.Service.JSONVersion != "" {
		jsonVersion := req.Service.JSONVersion
		req.HTTPRequest.Header.Add("Content-Type", "application/x-amz-json-"+jsonVersion)
	}
}

// Unmarshal unmarshals a response for a JSON RPC service.
func Unmarshal(req *service.Request) {
	defer req.HTTPResponse.Body.Close()
	if req.DataFilled() {
		err := jsonutil.UnmarshalJSON(req.Data, req.HTTPResponse.Body)
		if err != nil {
			req.Error = awserr.New("SerializationError", "failed decoding JSON RPC response", err)
		}
	}
	return
}

// UnmarshalMeta unmarshals headers from a response for a JSON RPC service.
func UnmarshalMeta(req *service.Request) {
	req.RequestID = req.HTTPResponse.Header.Get("x-amzn-requestid")
}

// UnmarshalError unmarshals an error response for a JSON RPC service.
func UnmarshalError(req *service.Request) {
	defer req.HTTPResponse.Body.Close()
	bodyBytes, err := ioutil.ReadAll(req.HTTPResponse.Body)
	if err != nil {
		req.Error = awserr.New("SerializationError", "failed reading JSON RPC error response", err)
		return
	}
	if len(bodyBytes) == 0 {
		req.Error = awserr.NewRequestFailure(
			awserr.New("SerializationError", req.HTTPResponse.Status, nil),
			req.HTTPResponse.StatusCode,
			"",
		)
		return
	}
	var jsonErr jsonErrorResponse
	if err := json.Unmarshal(bodyBytes, &jsonErr); err != nil {
		req.Error = awserr.New("SerializationError", "failed decoding JSON RPC error response", err)
		return
	}

	codes := strings.SplitN(jsonErr.Code, "#", 2)
	req.Error = awserr.NewRequestFailure(
		awserr.New(codes[len(codes)-1], jsonErr.Message, nil),
		req.HTTPResponse.StatusCode,
		"",
	)
}

type jsonErrorResponse struct {
	Code    string `json:"__type"`
	Message string `json:"message"`
}
