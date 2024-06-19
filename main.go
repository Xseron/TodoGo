package main

import (
	"net/http"

	"github.com/Xseron/MyExtraGo/handlers"
)

func main() {
	http.ListenAndServe(":3000", handlers.Init())
}
