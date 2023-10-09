package rest

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"
)

type ErrorTestCase struct {
	Input Error
	Code  int
	Str   string
	JSON  string
	Err   error
}

var errorTestCases = []ErrorTestCase{
	// Empty error
	{
		Input: Err,
		Code:  200,
		Str:   "",
		JSON:  `{"message":""}`,
	},

	// Standard errors
	{
		Input: ErrPermanentRedirect,
		Code:  308,
		Str:   "Permanent Redirect",
		JSON:  `{"message":"Permanent Redirect"}`,
	},
	{
		Input: ErrNotFound,
		Code:  404,
		Str:   "Not Found",
		JSON:  `{"message":"Not Found"}`,
	},
	{
		Input: ErrInternalServerError,
		Code:  500,
		Str:   "Internal Server Error",
		JSON:  `{"message":"Internal Server Error"}`,
	},

	// Error with changed message
	{
		Input: ErrBadRequest.WithMessage("Invalid Recipe"),
		Code:  400,
		Str:   "Invalid Recipe",
		JSON:  `{"message":"Invalid Recipe"}`,
	},

	// Error with data
	{
		Input: ErrGatewayTimeout.WithData(map[string]any{"service": "RecipeDatabase"}),
		Code:  504,
		Str:   "Gateway Timeout",
		JSON:  `{"message":"Gateway Timeout","data":{"service":"RecipeDatabase"}}`,
	},

	// Error with value
	{
		Input: ErrGatewayTimeout.WithValue("service", "RecipeDatabase"),
		Code:  504,
		Str:   "Gateway Timeout",
		JSON:  `{"message":"Gateway Timeout","data":{"service":"RecipeDatabase"}}`,
	},

	// Error with error
	{
		Input: ErrInternalServerError.WithError(errors.New("recipe is too delicious")),
		Code:  500,
		Str:   "Internal Server Error",
		JSON:  `{"message":"Internal Server Error","data":{"error":"recipe is too delicious"}}`,
	},
}

func TestErrorWrite(t *testing.T) {
	for i, tc := range errorTestCases {
		t.Logf("(%d) Testing %v", i, tc.Input)

		rec := httptest.NewRecorder()

		_, err := tc.Input.Write(rec)
		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
		if err != nil {
			continue
		}

		res := rec.Result()
		if res.StatusCode != tc.Code {
			t.Errorf("Expected status code %d, got %d", tc.Code, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Unexpected error reading response body: %v", err)
			continue
		}

		if string(body) != tc.Str {
			t.Errorf("Expected body %q, got %q", tc.Str, string(body))
		}
	}
}

func TestErrorWriteJSON(t *testing.T) {
	for i, tc := range errorTestCases {
		t.Logf("(%d) Testing %v", i, tc.Input)

		rec := httptest.NewRecorder()

		err := tc.Input.WriteJSON(rec)
		if err != tc.Err {
			t.Errorf("Expected error %v, got %v", tc.Err, err)
		}
		if err != nil {
			continue
		}

		res := rec.Result()
		if res.StatusCode != tc.Code {
			t.Errorf("Expected status code %d, got %d", tc.Code, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("Unexpected error reading response body: %v", err)
			continue
		}

		if string(body) != tc.JSON {
			t.Errorf("Expected body %q, got %q", tc.JSON, string(body))
		}
	}
}
