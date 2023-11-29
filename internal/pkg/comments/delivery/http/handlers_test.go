package http

import (
	"net/http"
	"testing"
)

func TestCommentsHandler_CreateComment(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *CommentsHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.CreateComment(tt.args.w, tt.args.r)
		})
	}
}

func TestCommentsHandler_GetProductComments(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *CommentsHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetProductComments(tt.args.w, tt.args.r)
		})
	}
}
