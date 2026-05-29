package main

import (
	"fmt"
	"io"
	"net/http"
)

var memoryStorage = make(map[string][]byte)

func handleObject(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")