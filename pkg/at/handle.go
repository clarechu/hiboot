package at

import "github.com/gorilla/mux"

type Handle func(router *mux.Router) error
