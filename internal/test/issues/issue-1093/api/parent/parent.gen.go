// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/CeskyPane/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// Pet defines model for Pet.
type Pet struct {
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /pets)
	GetPets(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetPets operation middleware
func (siw *ServerInterfaceWrapper) GetPets(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetPets(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/pets", wrapper.GetPets)
}

type GetPetsRequestObject struct {
}

type GetPetsResponseObject interface {
	VisitGetPetsResponse(w http.ResponseWriter) error
}

type GetPets200JSONResponse Pet

func (response GetPets200JSONResponse) VisitGetPetsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /pets)
	GetPets(ctx context.Context, request GetPetsRequestObject) (GetPetsResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetPets operation middleware
func (sh *strictHandler) GetPets(ctx *gin.Context) {
	var request GetPetsRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetPets(ctx, request.(GetPetsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetPets")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(GetPetsResponseObject); ok {
		if err := validResponse.VisitGetPetsResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/2xRwa7TMBD8FWvhGDUBbv4BxO0JcXv04Dqb2lViL7sbUBXl39E6rVAlTru2xjPjmQ1i",
	"XagWLCrgN5CYcAltfUO1QVwJWTO2yxIWtKl3QvAgyrlcYe9Aw/U/93sHjL/WzDiCfz9en7snql5uGBV2",
	"g+UyVSMYUSJn0lwLePiRsjgKjEWdEEaXgriljjiL+5NyTC4wurxQZcXRXe5OE7qY8jw2PHSgWWcTO1ig",
	"g9/IcrB/Og1mvRKWQBk8fDkNpwE6oKCpfbcnPHK5HmG8uvuOunIRR6hu4ro0cbmLoq1B23kV5GY7xIgi",
	"TuvPAk2Ug/F8G8HDV9Q3U7K4hGqRI+3Pw2Aj1qLm3W8QiOYc28P+Jubi2ZltHxkn8PCh/1dq/2i0tzpb",
	"0q9/CE7WZmxaZ/cUN6BBBdnSAv++wcozeOgfMe7n/W8AAAD///mNvoM7AgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
