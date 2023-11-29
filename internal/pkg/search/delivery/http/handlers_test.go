package http

import (
	"net/http"
	"testing"
)

func TestSearchHandler_SearchProducts(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *SearchHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.SearchProducts(tt.args.w, tt.args.r)
		})
	}
}
