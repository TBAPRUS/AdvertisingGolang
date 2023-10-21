package main

import (
	"log"
	"net/http"

	"tbaprus/advertisinggolang/handler"
)

func main() {
	advertisementHadler := &handler.AdvertisementHadler{}

	http.Handle("/advertisements/", advertisementHadler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
