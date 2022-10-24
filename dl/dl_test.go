package dl

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

// Custom type that allows setting the func that our Mock Do func will run instead
type mockDoType func(req *http.Request) (*http.Response, error)

// MockClient is the mock client
type mockClient struct {
	mockDo mockDoType
}

// Overriding what the Do function should "do" in our MockClient
func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.mockDo(req)
}

func TestDownload(t *testing.T) {
	// build our response JSON
	jsonResponse := `[{"full_name": "mock-repo"}]`
	// create a new reader with that JSON
	r := io.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	httpClient = func() HTTPClient {
		return &mockClient{
			mockDo: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 200,
					Body:       r,
				}, nil
			},
		}
	}

	_, err := Download("test", "token")
	if err != nil {
		t.Errorf("failed success response test: %v", err)
	}

	httpClient = func() HTTPClient {
		return &mockClient{
			mockDo: func(*http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 404,
					Body:       io.NopCloser(bytes.NewReader([]byte("Error 404"))),
				}, nil
			},
		}
	}

	b, err := Download("test", "token")
	if err == nil {
		t.Errorf("expected to fail due to 404 error but didn't: %v", string(b))
	}

	httpClient = func() HTTPClient {
		return &mockClient{
			mockDo: func(*http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("failed to execute test request")
			},
		}
	}

	_, err = Download("test", "token")
	if err == nil {
		t.Errorf("expected to return error, but didn't")
	}

	httpClient = createClient

	_, err = Download("http://thisisnotavalid url", "token")
	if err == nil {
		t.Errorf("expected to fail in request creation")
	}

	_, err = Download("http://thisisnotavalidurl", "token")
	if err == nil {
		t.Errorf("expected to fail due to invalid url")
	}
}
