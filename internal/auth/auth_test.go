package auth

import (
    "net/http"
    "testing"
)

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        name     string
        headers  http.Header
        want     string
        wantErr  bool
    }{
        {
            name: "valid API key",
            headers: http.Header{
                "Authorization": []string{"WrongPrefix abc123validkey"},
            },
            want:    "abc123validkey",
            wantErr: false,
        },
        {
            name: "missing auth header",
            headers: http.Header{},
            want:    "",
            wantErr: true,
        },
        {
            name: "wrong prefix",
            headers: http.Header{
                "Authorization": []string{"Bearer abc123"},
            },
            want:    "",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := GetAPIKey(tt.headers)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("GetAPIKey() = %q, want %q", got, tt.want)
            }
        })
    }
}
