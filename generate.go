package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deepmap/oapi-codegen/pkg/codegen"
	codegen_util "github.com/deepmap/oapi-codegen/pkg/util"
)

const AUVIK_OPENAPI_SPEC_URL string = "https://auvikapi.us1.my.auvik.com/spec"

func main() {
	res, err := http.Get(AUVIK_OPENAPI_SPEC_URL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("failed to retrieve Auvik API spec: %d %s", res.StatusCode, res.Status)
		panic(err)
	}
	specFile, err := os.CreateTemp("", "auvik-api-spec-*.json")
	if err != nil {
		panic(err)
	}
	specData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	specReplaced := bytes.ReplaceAll(specData, []byte(`"type" : "double"`), []byte(`"type" : "number"`))
	_, err = specFile.Write(specReplaced)
	if err != nil {
		panic(err)
	}
	spec, err := codegen_util.LoadSwagger(specFile.Name())
	if err != nil {
		panic(err)
	}

	clientResult, err := codegen.Generate(spec, codegen.Configuration{
		PackageName: "auvik",
		Generate: codegen.GenerateOptions{
			Client: true,
		},
		OutputOptions: codegen.OutputOptions{},
		Compatibility: codegen.CompatibilityOptions{},
	})
	if err != nil {
		panic(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	clientFile := filepath.Join(cwd, "client", "client.go")
	clientOutFile, err := os.Create(clientFile)
	if err != nil {
		panic(err)
	}
	clientOutFile.WriteString(clientResult)
	log.Printf("wrote generated client to %s", clientFile)

	typesResult, err := codegen.Generate(spec, codegen.Configuration{
		PackageName: "auvik",
		Generate: codegen.GenerateOptions{
			Models: true,
		},
		OutputOptions: codegen.OutputOptions{},
	})
	if err != nil {
		panic(err)
	}
	typesFile := filepath.Join(cwd, "client", "types.go")
	typesOutFile, err := os.Create(typesFile)
	if err != nil {
		panic(err)
	}
	typesOutFile.WriteString(typesResult)
	log.Printf("wrote generated types to %s", typesFile)
}
