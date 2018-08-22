package httputil

import (
	"testing"
)

func TestValidateSignupCredentials(t *testing.T) {
	if UnmarshalRequestBody(nil, nil) == nil {
		t.Errorf("expected an error - received nil")
	}
}
