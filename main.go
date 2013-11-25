package main

import (
	"fmt"
	"net/http"
	"strconv"
	"testgouga/math"
)

func main() {
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		units, _ := strconv.ParseFloat(r.URL.Query()["units"][0], 64)
		price, _ := strconv.ParseFloat(r.URL.Query()["price"][0], 64)
		val := math.Calc(units, price)
		fmt.Println(strconv.FormatFloat(val, 'f', 2, 64))
		fmt.Fprintf(w, strconv.FormatFloat(val, 'f', 2, 64))
	})
	http.ListenAndServe(":8081", nil)
}
