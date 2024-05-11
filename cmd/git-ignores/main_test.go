package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchGitignoreTemplate(t *testing.T) {
	tests := []struct {
		name           string
		templateName   string
		responseStatus int
		responseBody   string
		wantErr        bool
	}{
		{
			name:           "Success",
			templateName:   "Go",
			responseStatus: http.StatusOK,
			responseBody:   "",
			wantErr:        false,
		},
		{
			name:           "NotFound",
			templateName:   "NonExistentTemplate",
			responseStatus: http.StatusNotFound,
			responseBody:   "",
			wantErr:        true,
		},
		{
			name:           "ServerError",
			templateName:   "SomeTemplate",
			responseStatus: http.StatusInternalServerError,
			responseBody:   "",
			wantErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path == "/github/gitignore/main/"+tt.templateName+".gitignore" {
					w.WriteHeader(tt.responseStatus)
					w.Write([]byte(tt.responseBody))
					return
				}
				w.WriteHeader(http.StatusNotFound)
			}))
			defer server.Close()

			_, err := fetchGitignoreTemplate(tt.templateName)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchGitignoreTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestWriteIgnoreFile(t *testing.T) {
	content := "test content"
	err := writeIgnoreFile(content)
	if err != nil {
		t.Errorf("writeIgnoreFile() error = %v, wantErr %v", err, false)
	}

	fileContent, err := ioutil.ReadFile(".gitignore")
	if err != nil {
		t.Errorf("writeIgnoreFile() error reading file = %v", err)
	}

	if string(fileContent) != content {
		t.Errorf("writeIgnoreFile() content = %v, want %v", string(fileContent), content)
	}

	// Clean up
	os.Remove(".gitignore")
}
