package http

import (
	"context"
	"reflect"
	"testing"
)

var h HttpClient

func TestHttpClient_do(t *testing.T) {
	prefix := "/api/users/login"
	ctx := context.TODO()
	h, _ := NewHttpClient(&ctx, prefix)

	type args struct {
		method  string
		url     string
		body    []byte
		headers map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantBuf []byte
		wantErr bool
	}{
		{
			name:    "it should fail with notify other fail",
			wantBuf: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBuf, err := h.do(tt.args.method, tt.args.url, tt.args.body, tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpClient.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBuf, tt.wantBuf) {
				t.Errorf("HttpClient.do() = %v, want %v", gotBuf, tt.wantBuf)
			}
		})
	}
}
