// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: fcd9ff140d
// Version Date: 2021-07-14T06:36:40Z

package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"

	"context"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	// This service
	pb "github.com/Presbyter/services/protobuf"
)

const contentType = "application/json; charset=utf-8"

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.NewUserClient
	_ = io.Copy
	_ = errors.Wrap
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(endpoints Endpoints, responseEncoder httptransport.EncodeResponseFunc, options ...httptransport.ServerOption) http.Handler {
	if responseEncoder == nil {
		responseEncoder = EncodeHTTPGenericResponse
	}
	serverOptions := []httptransport.ServerOption{
		httptransport.ServerBefore(headersToContext),
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerAfter(httptransport.SetContentType(contentType)),
	}
	serverOptions = append(serverOptions, options...)
	m := mux.NewRouter()

	m.Methods("GET").Path("/api/v1/user/{user_id}/info").Handler(httptransport.NewServer(
		endpoints.UserInfoEndpoint,
		DecodeHTTPUserInfoZeroRequest,
		responseEncoder,
		serverOptions...,
	))
	return m
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// httpError satisfies the Headerer and StatusCoder interfaces in
// package github.com/go-kit/kit/transport/http.
type httpError struct {
	error
	statusCode int
	headers    map[string][]string
}

func (h httpError) StatusCode() int {
	return h.statusCode
}

func (h httpError) Headers() http.Header {
	return h.headers
}

// Server Decode

// DecodeHTTPUserInfoZeroRequest is a transport/http.DecodeRequestFunc that
// decodes a JSON-encoded userinfo request from the HTTP request
// body. Primarily useful in a server.
func DecodeHTTPUserInfoZeroRequest(_ context.Context, r *http.Request) (interface{}, error) {
	defer r.Body.Close()
	var req pb.UserInfoReq
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read body of http request")
	}
	if len(buf) > 0 {
		// AllowUnknownFields stops the unmarshaler from failing if the JSON contains unknown fields.
		unmarshaller := jsonpb.Unmarshaler{
			AllowUnknownFields: true,
		}
		if err = unmarshaller.Unmarshal(bytes.NewBuffer(buf), &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, httpError{errors.Wrapf(err, "request body '%s': cannot parse non-json request body", buf),
				http.StatusBadRequest,
				nil,
			}
		}
	}

	pathParams := encodePathParams(mux.Vars(r))
	_ = pathParams

	queryParams := r.URL.Query()
	_ = queryParams

	UserIdUserInfoStr := pathParams["user_id"]
	UserIdUserInfo, err := strconv.ParseUint(UserIdUserInfoStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error while extracting UserIdUserInfo from path, pathParams: %v", pathParams))
	}
	req.UserId = UserIdUserInfo

	return &req, err
}

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	marshaller := jsonpb.Marshaler{
		EmitDefaults: false,
		OrigName:     true,
	}

	return marshaller.Marshal(w, response.(proto.Message))
}

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "request-url", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}

// encodePathParams encodes `mux.Vars()` with dot notations into JSON objects
// to be unmarshaled into non-basetype fields.
// e.g. {"book.name": "books/1"} -> {"book": {"name": "books/1"}}
func encodePathParams(vars map[string]string) map[string]string {
	var recur func(path, value string, data map[string]interface{})
	recur = func(path, value string, data map[string]interface{}) {
		parts := strings.SplitN(path, ".", 2)
		key := parts[0]
		if len(parts) == 1 {
			data[key] = value
		} else {
			if _, ok := data[key]; !ok {
				data[key] = make(map[string]interface{})
			}
			recur(parts[1], value, data[key].(map[string]interface{}))
		}
	}

	data := make(map[string]interface{})
	for key, val := range vars {
		recur(key, val, data)
	}

	ret := make(map[string]string)
	for key, val := range data {
		switch val := val.(type) {
		case string:
			ret[key] = val
		case map[string]interface{}:
			m, _ := json.Marshal(val)
			ret[key] = string(m)
		}
	}
	return ret
}
