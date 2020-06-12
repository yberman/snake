package main

import (
	// "fmt"
	"net/http"
	// "strings"

	// "github.com/gorilla/websocket"
	"github.com/s4y/reserve"
)

type Point struct {
	X, Y int
}

type State struct {
	Apples []Point
	Snake1 []Point
	Snake2 []Point
}

func main() {
	http.Handle("/", reserve.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", nil)
}
