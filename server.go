package main

import (
	"fmt"
	"net/http"
	"strings"
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
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

const homepage = `
<!DOCTYPE html>
<html>
<head>
<title>SNAAAAAKE!!!!</title>
</head>
<body>
%s
</body>
</html>
`

func makeBoardHTML() string {
	table := []string{}
	n := 10
	for row := 0; row < n; row++ {
		rowCells := []string{}
		for col := 0; col < n; col++ {
			rowCells = append(rowCells, "<td><div style=\"background: #abc6f8;\">hi</div></td>")
		}
		table = append(table,  "<tr>" + strings.Join(rowCells, "") + "</tr>")
	}
	return "<table>" + strings.Join(table, "") + "</table>"
}


func mainHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("CacheControl", "no-cache, no-store, must-revalidate")
	fmt.Fprintf(w, homepage, makeBoardHTML())
}

