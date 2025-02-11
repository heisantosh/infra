// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
