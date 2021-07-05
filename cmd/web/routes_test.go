package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/tbetcha/gotel/internal/config"
	"testing"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig
	mux:= routes(&app)
	switch v := mux.(type){
	case *chi.Mux:
		// do nada
		default:
			t.Error(fmt.Sprintf("Type is not chi.mux type is %t", v))
	}
}
