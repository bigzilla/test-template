package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/injustease/is"
	myhttp "github.com/injustease/test-template/server/http"
)

func TestDecodeSchema(t *testing.T) {
	var (
		param1 = "foobar"
		param2 = 99
		url    = fmt.Sprintf("/?param1=%s&param2=%d", param1, param2)
	)

	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		var request = struct {
			Param1 string `schema:"param1"`
			Param2 int    `schema:"param2"`
		}{}

		err := myhttp.DecodeSchema(nil, req, &request)
		is.NoError(err)
		is.Equal(request.Param1, param1)
		is.Equal(request.Param2, param2)
	})

	t.Run("error", func(t *testing.T) {
		is := is.New(t)

		req := httptest.NewRequest(http.MethodGet, url, nil)
		var request = struct {
			Param1 string `schema:"param1"`
			Param2 int    `schema:"param2"`
		}{}

		err := myhttp.DecodeSchema(nil, req, request)
		is.Error(err)
	})
}

func TestDecodeJSON(t *testing.T) {
	var (
		param1 = "foobar"
		param2 = 99
	)

	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		var buff bytes.Buffer
		is.NoError(json.NewEncoder(&buff).Encode(map[string]interface{}{
			"param1": param1,
			"param2": param2,
		}))

		req := httptest.NewRequest(http.MethodGet, "/", &buff)
		var request = struct {
			Param1 string `json:"param1"`
			Param2 int    `json:"param2"`
		}{}

		err := myhttp.DecodeJSON(nil, req, &request)
		is.NoError(err)
		is.Equal(request.Param1, param1)
		is.Equal(request.Param2, param2)
	})

	t.Run("error", func(t *testing.T) {
		is := is.New(t)

		var buff bytes.Buffer
		is.NoError(json.NewEncoder(&buff).Encode(map[string]interface{}{
			"param1": param1,
			"param2": param2,
		}))

		req := httptest.NewRequest(http.MethodGet, "/", &buff)
		var request = struct {
			Param1 string `schema:"param1"`
			Param2 int    `schema:"param2"`
		}{}

		err := myhttp.DecodeJSON(nil, req, request)
		is.Error(err)
	})
}

func TestRespondJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		msg := "success"
		myhttp.RespondJSON(rec, req, http.StatusOK, msg)

		res := rec.Result()
		defer res.Body.Close()

		var body = struct {
			Data string
		}{}
		is.Equal(res.StatusCode, http.StatusOK)
		is.NoError(json.NewDecoder(res.Body).Decode(&body))
		is.Equal(body.Data, msg)
	})

	t.Run("error", func(t *testing.T) {
		is := is.New(t)

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		msg := errors.New("error")
		myhttp.RespondJSON(rec, req, http.StatusBadRequest, msg)

		res := rec.Result()
		defer res.Body.Close()

		var body = struct {
			Errors string
		}{}
		is.Equal(res.StatusCode, http.StatusBadRequest)
		is.NoError(json.NewDecoder(res.Body).Decode(&body))
		is.Equal(body.Errors, msg.Error())
	})
}
