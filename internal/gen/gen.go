package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/stellaraf/go-utils"
)

const AUVIK_OPENAPI_SPEC_URL string = "https://auvikapi.us1.my.auvik.com/spec"

func getSpecPath(base string) (string, error) {
	specPath := filepath.Join(base, "auvik-api-spec.json")
	specPath, err := filepath.Abs(specPath)
	if err != nil {
		return "", err
	}
	_, err = os.Stat(specPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(specPath)
			if err != nil {
				return "", err
			}
			log.Printf("created %s", specPath)
			defer f.Close()
			return specPath, nil
		}
		return "", err
	}
	return specPath, nil
}

func downloadSpec(specPath string) error {
	res, err := http.Get(AUVIK_OPENAPI_SPEC_URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("failed to retrieve Auvik API spec: %d %s", res.StatusCode, res.Status)
		return err
	}
	specFile, err := os.OpenFile(specPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer specFile.Close()
	specData, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	specReplaced := bytes.ReplaceAll(specData, []byte(`"type" : "double"`), []byte(`"type" : "number"`))
	_, err = specFile.Write(specReplaced)
	if err != nil {
		return err
	}
	log.Printf("wrote spec to %s", specFile.Name())
	return nil
}

func main() {
	root, err := utils.FindProjectRoot(4)
	if err != nil {
		panic(err)
	}
	specPath, err := getSpecPath(root)
	if err != nil {
		panic(err)
	}
	err = downloadSpec(specPath)
	if err != nil {
		panic(err)
	}
}
