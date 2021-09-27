package main

import (
	"net/http"

	"github.com/pixsil/hdmi-cec-rest/webservice"
)

func main() {
	router := webservice.GetRouter()
	http.ListenAndServe(":5001", router)
}
