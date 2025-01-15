package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	cases := []struct {
		name         string
		header       string
		expectApiKey string
		expectError  bool
	}{
		{
			name:         "Valid",
			header:       "ApiKey 123",
			expectApiKey: "123",
			expectError:  true,
		},
		{
			name:         "Empty",
			header:       "",
			expectApiKey: "",
			expectError:  true,
		},
		{
			name:         "Wrong prefix",
			header:       "WordyWordy 123",
			expectApiKey: "",
			expectError:  true,
		},
		{
			name:         "No ApiKey value",
			header:       "ApiKey",
			expectApiKey: "",
			expectError:  true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			header := http.Header{}
			header.Add("Authorization", tc.header)
			apiKey, err := GetAPIKey(header)
			if (err != nil) != tc.expectError {
				t.Errorf("expectError=%v but err=%v", tc.expectError, err)
			}
			if apiKey != tc.expectApiKey {
				t.Errorf("expectApiKey=%v but got %v", tc.expectApiKey, apiKey)
			}
		})
	}
}
