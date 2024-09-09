// Package spec_base provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/CeskyPane/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package spec_base

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	externalRef0 "github.com/CeskyPane/oapi-codegen/v2/internal/test/issues/issue-removed-external-ref/gen/spec_ext"
	"github.com/go-chi/chi/v5"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// DirectBar defines model for DirectBar.
type DirectBar = externalRef0.Foo

// PackedBar defines model for PackedBar.
type PackedBar struct {
	Core    *externalRef0.Foo `json:"core,omitempty"`
	Directd *DirectBar        `json:"directd,omitempty"`
	Id      *string           `json:"id,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /invalidExtRefTrouble)
	PostInvalidExtRefTrouble(w http.ResponseWriter, r *http.Request)

	// (POST /noTrouble)
	PostNoTrouble(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// (POST /invalidExtRefTrouble)
func (_ Unimplemented) PostInvalidExtRefTrouble(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// (POST /noTrouble)
func (_ Unimplemented) PostNoTrouble(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostInvalidExtRefTrouble operation middleware
func (siw *ServerInterfaceWrapper) PostInvalidExtRefTrouble(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostInvalidExtRefTrouble(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostNoTrouble operation middleware
func (siw *ServerInterfaceWrapper) PostNoTrouble(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNoTrouble(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/invalidExtRefTrouble", wrapper.PostInvalidExtRefTrouble)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/noTrouble", wrapper.PostNoTrouble)
	})

	return r
}

type PostInvalidExtRefTroubleRequestObject struct {
}

type PostInvalidExtRefTroubleResponseObject interface {
	VisitPostInvalidExtRefTroubleResponse(w http.ResponseWriter) error
}

type PostInvalidExtRefTrouble300JSONResponse struct{ externalRef0.Pascal }

func (response PostInvalidExtRefTrouble300JSONResponse) VisitPostInvalidExtRefTroubleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(300)

	return json.NewEncoder(w).Encode(response)
}

type PostNoTroubleRequestObject struct {
}

type PostNoTroubleResponseObject interface {
	VisitPostNoTroubleResponse(w http.ResponseWriter) error
}

type PostNoTrouble200JSONResponse struct {
	DirectBar   *DirectBar        `json:"directBar,omitempty"`
	DirectFoo   *externalRef0.Foo `json:"directFoo,omitempty"`
	IndirectFoo *PackedBar        `json:"indirectFoo,omitempty"`
	Name        *string           `json:"name,omitempty"`
}

func (response PostNoTrouble200JSONResponse) VisitPostNoTroubleResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (POST /invalidExtRefTrouble)
	PostInvalidExtRefTrouble(ctx context.Context, request PostInvalidExtRefTroubleRequestObject) (PostInvalidExtRefTroubleResponseObject, error)

	// (POST /noTrouble)
	PostNoTrouble(ctx context.Context, request PostNoTroubleRequestObject) (PostNoTroubleResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostInvalidExtRefTrouble operation middleware
func (sh *strictHandler) PostInvalidExtRefTrouble(w http.ResponseWriter, r *http.Request) {
	var request PostInvalidExtRefTroubleRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostInvalidExtRefTrouble(ctx, request.(PostInvalidExtRefTroubleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostInvalidExtRefTrouble")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostInvalidExtRefTroubleResponseObject); ok {
		if err := validResponse.VisitPostInvalidExtRefTroubleResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostNoTrouble operation middleware
func (sh *strictHandler) PostNoTrouble(w http.ResponseWriter, r *http.Request) {
	var request PostNoTroubleRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostNoTrouble(ctx, request.(PostNoTroubleRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostNoTrouble")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostNoTroubleResponseObject); ok {
		if err := validResponse.VisitPostNoTroubleResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
