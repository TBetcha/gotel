package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var testHandler testHandler
	h := NoSurf(&testHandler)

	switch v := h.(type){
	case http.Handler:
		//do nothing
		default:
			t.Error(fmt.Sprintf("type is not handler %t", v))
	}
}
func TestSessionLoad(t *testing.T) {
	var testHandler testHandler
	h := SessionLoad(&testHandler)

	switch v := h.(type){
	case http.Handler:
	//do nothing
	default:
		t.Error(fmt.Sprintf("type is not handler %t", v))
	}
}
