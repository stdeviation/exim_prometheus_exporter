package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Run(port int, path string, handler http.Handler) error {
	http.Handle(path, handler)
	return http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func createTempFile(prefix string, data []byte) (string, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), prefix)
	if err != nil {
		return "", fmt.Errorf("Failed to create temporary file: %v", err.Error())
	}
	_, err = tempFile.Write(data)
	if err != nil {
		return "", fmt.Errorf("Failed to write temporary file: %v", err.Error())
	}
	err = tempFile.Close()
	if err != nil {
		return "", fmt.Errorf("Failed to close temporary file: %v", err.Error())
	}
	return tempFile.Name(), nil
}
