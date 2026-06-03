package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name  string
		input http.Header
		want  string
		err   error
	}{
		{
			name:  "no auth header returns error",
			input: http.Header{},
			want:  "",
			err:   ErrNoAuthHeaderIncluded,
		},
		{
			name:  "Valid key",
			input: http.Header{"Authorization": []string{"ApiKey mysectetkey"}},
			want:  "mysectetkey",
			err:   nil,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if !reflect.DeepEqual(tc.want, got) || err != tc.err {
			t.Fatalf("Test %s failed: got %s, want %s, got error %v, want error %v", tc.name, got, tc.want, err, tc.err)
		}
	}
}
