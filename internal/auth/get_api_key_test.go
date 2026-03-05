package auth

import (
	"testing"
	"reflect"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name	string
		headers	http.Header
		want	string
		wantErr	bool
	}{
		{
			name:	 "valid header",
			headers: http.Header{"Authorization": []string{"ApiKey abc123"}},
			want:	 "abc123",
			wantErr: false,
		},
		{
			name:	 "empty header",
			headers: http.Header{"Authorization": []string{""}},
			want:	 "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
	    	got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error in testing: %v", err)
			}
	    	if !reflect.DeepEqual(tt.want, got) {
	    	     t.Fatalf("expected: %v, got: %v", tt.want, got)
	    	}
		})
	}
}
