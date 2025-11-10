package sendgrid

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestParseErrorDetails(t *testing.T) {
	tests := []struct {
		name          string
		err           error
		wantEnhanced  bool
		wantContains  string
	}{
		{
			name:         "scope error",
			err:          errors.New("invalid or unassignable scopes provided"),
			wantEnhanced: true,
			wantContains: "invalid or unassignable scopes",
		},
		{
			name:         "permission error",
			err:          errors.New("permission denied"),
			wantEnhanced: true,
			wantContains: "permission denied",
		},
		{
			name:         "unauthorized error",
			err:          errors.New("unauthorized access"),
			wantEnhanced: true,
			wantContains: "permission denied",
		},
		{
			name:         "not found error",
			err:          errors.New("resource not found"),
			wantEnhanced: true,
			wantContains: "resource not found",
		},
		{
			name:         "404 error",
			err:          errors.New("404 not found"),
			wantEnhanced: true,
			wantContains: "resource not found",
		},
		{
			name:         "validation error",
			err:          errors.New("validation failed"),
			wantEnhanced: true,
			wantContains: "validation error",
		},
		{
			name:         "invalid error",
			err:          errors.New("invalid request"),
			wantEnhanced: true,
			wantContains: "validation error",
		},
		{
			name:         "generic error",
			err:          errors.New("something went wrong"),
			wantEnhanced: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, enhanced := parseErrorDetails(tt.err)
			if enhanced != tt.wantEnhanced {
				t.Errorf("parseErrorDetails() enhanced = %v, want %v", enhanced, tt.wantEnhanced)
			}
			if tt.wantEnhanced && !strings.Contains(msg, tt.wantContains) {
				t.Errorf("parseErrorDetails() message = %v, want to contain %v", msg, tt.wantContains)
			}
		})
	}
}

func TestEnhanceError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		statusCode   int
		wantContains string
	}{
		{
			name:         "400 bad request",
			err:          errors.New("something went wrong"),
			statusCode:   http.StatusBadRequest,
			wantContains: "bad request (HTTP 400)",
		},
		{
			name:         "401 unauthorized",
			err:          errors.New("auth failed"),
			statusCode:   http.StatusUnauthorized,
			wantContains: "unauthorized (HTTP 401)",
		},
		{
			name:         "403 forbidden",
			err:          errors.New("access denied"),
			statusCode:   http.StatusForbidden,
			wantContains: "forbidden (HTTP 403)",
		},
		{
			name:         "404 not found",
			err:          errors.New("resource missing"),
			statusCode:   http.StatusNotFound,
			wantContains: "resource not found (HTTP 404)",
		},
		{
			name:         "429 rate limit",
			err:          errors.New("too many requests"),
			statusCode:   http.StatusTooManyRequests,
			wantContains: "rate limit exceeded (HTTP 429)",
		},
		{
			name:         "500 server error",
			err:          errors.New("internal error"),
			statusCode:   http.StatusInternalServerError,
			wantContains: "server error (HTTP 500)",
		},
		{
			name:         "502 bad gateway",
			err:          errors.New("gateway error"),
			statusCode:   http.StatusBadGateway,
			wantContains: "server error (HTTP 502)",
		},
		{
			name:         "503 service unavailable",
			err:          errors.New("service down"),
			statusCode:   http.StatusServiceUnavailable,
			wantContains: "server error (HTTP 503)",
		},
		{
			name:         "other status",
			err:          errors.New("some error"),
			statusCode:   418, // I'm a teapot
			wantContains: "request failed with HTTP 418",
		},
		{
			name:         "nil error",
			err:          nil,
			statusCode:   200,
			wantContains: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := enhanceError(tt.err, tt.statusCode)
			if tt.err == nil {
				if result != nil {
					t.Errorf("enhanceError() = %v, want nil", result)
				}
				return
			}
			if result == nil {
				t.Errorf("enhanceError() = nil, want error")
				return
			}
			if !strings.Contains(result.Error(), tt.wantContains) {
				t.Errorf("enhanceError() = %v, want to contain %v", result.Error(), tt.wantContains)
			}
		})
	}
}

func TestRetryOnRateLimit(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		d := &schema.ResourceData{}
		d.SetId("test")
		ctx := context.Background()

		result, err := RetryOnRateLimit(ctx, d, func() (interface{}, RequestError) {
			return "success", RequestError{StatusCode: http.StatusOK, Err: nil}
		})

		if err != nil {
			t.Errorf("RetryOnRateLimit() error = %v, want nil", err)
		}
		if result != "success" {
			t.Errorf("RetryOnRateLimit() result = %v, want 'success'", result)
		}
	})

	t.Run("non-retryable error", func(t *testing.T) {
		d := &schema.ResourceData{}
		d.SetId("test")
		ctx := context.Background()

		_, err := RetryOnRateLimit(ctx, d, func() (interface{}, RequestError) {
			return nil, RequestError{
				StatusCode: http.StatusBadRequest,
				Err:        errors.New("bad request"),
			}
		})

		if err == nil {
			t.Error("RetryOnRateLimit() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "bad request") {
			t.Errorf("RetryOnRateLimit() error = %v, want to contain 'bad request'", err)
		}
	})
}

func TestAPIErrorDetail(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want string
	}{
		{
			name: "map with detail",
			data: map[string]interface{}{"detail": "error detail"},
			want: "error detail",
		},
		{
			name: "map without detail",
			data: map[string]interface{}{"error": "some error", "code": 123},
			want: "map[code:123 error:some error]",
		},
		{
			name: "string data",
			data: "plain string error",
			want: "plain string error",
		},
		{
			name: "nil data",
			data: nil,
			want: "<nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := APIError{f: tt.data}
			got := apiErr.Detail()
			if !strings.Contains(got, tt.want) {
				t.Errorf("APIError.Detail() = %v, want to contain %v", got, tt.want)
			}
		})
	}
}

func TestAPIErrorEmpty(t *testing.T) {
	tests := []struct {
		name string
		data interface{}
		want bool
	}{
		{
			name: "nil data",
			data: nil,
			want: true,
		},
		{
			name: "non-nil data",
			data: "some error",
			want: false,
		},
		{
			name: "empty map",
			data: map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiErr := APIError{f: tt.data}
			if got := apiErr.Empty(); got != tt.want {
				t.Errorf("APIError.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

