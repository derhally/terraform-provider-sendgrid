package sendgrid

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name       string
		apiKey     string
		host       string
		onBehalfOf string
		wantNil    bool
		wantHost   string
	}{
		{
			name:       "valid client with default host",
			apiKey:     "test-api-key",
			host:       "",
			onBehalfOf: "",
			wantNil:    false,
			wantHost:   "https://api.sendgrid.com/v3/",
		},
		{
			name:       "valid client with custom host",
			apiKey:     "test-api-key",
			host:       "https://custom.sendgrid.com/v3/",
			onBehalfOf: "",
			wantNil:    false,
			wantHost:   "https://custom.sendgrid.com/v3/",
		},
		{
			name:       "with subuser",
			apiKey:     "test-api-key",
			host:       "",
			onBehalfOf: "subuser@example.com",
			wantNil:    false,
			wantHost:   "https://api.sendgrid.com/v3/",
		},
		{
			name:       "empty api key",
			apiKey:     "",
			host:       "",
			onBehalfOf: "",
			wantNil:    false, // Still creates client, validation happens at API call time
			wantHost:   "https://api.sendgrid.com/v3/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.apiKey, tt.host, tt.onBehalfOf)
			if (client == nil) != tt.wantNil {
				t.Errorf("NewClient() = %v, want nil = %v", client, tt.wantNil)
			}
			if client != nil {
				if client.apiKey != tt.apiKey {
					t.Errorf("NewClient().apiKey = %v, want %v", client.apiKey, tt.apiKey)
				}
				if client.OnBehalfOf != tt.onBehalfOf {
					t.Errorf("NewClient().OnBehalfOf = %v, want %v", client.OnBehalfOf, tt.onBehalfOf)
				}
				if client.host != tt.wantHost {
					t.Errorf("NewClient().host = %v, want %v", client.host, tt.wantHost)
				}
			}
		})
	}
}

func TestBodyToJSON(t *testing.T) {
	tests := []struct {
		name    string
		body    interface{}
		wantErr bool
		errType error
	}{
		{
			name:    "valid struct",
			body:    map[string]string{"key": "value"},
			wantErr: false,
		},
		{
			name:    "valid string",
			body:    "test string",
			wantErr: false,
		},
		{
			name:    "nil body",
			body:    nil,
			wantErr: true,
			errType: ErrBodyNotNil,
		},
		{
			name:    "complex struct",
			body:    struct{ Name string }{"test"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := bodyToJSON(tt.body)
			if tt.wantErr {
				if err == nil {
					t.Errorf("bodyToJSON() error = nil, want error")
				}
				if tt.errType != nil && err != tt.errType {
					t.Errorf("bodyToJSON() error = %v, want %v", err, tt.errType)
				}
			} else {
				if err != nil {
					t.Errorf("bodyToJSON() error = %v, want nil", err)
				}
				if result == nil {
					t.Error("bodyToJSON() result = nil, want non-nil")
				}
			}
		})
	}
}

