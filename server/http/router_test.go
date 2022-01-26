package http_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/injustease/is"
	myhttp "github.com/injustease/test-template/server/http"
)

func TestRouting(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{}

	srv := httptest.NewServer(myhttp.NewServer(""))
	defer srv.Close()

	for _, tt := range tests {
		is := is.New(t)

		res, err := http.Get(fmt.Sprintf("%s%s", srv.URL, tt.path))
		is.NoError(err)
		is.Equal(res.StatusCode, http.StatusOK)
	}
}
