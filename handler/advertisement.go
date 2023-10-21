package handler

import (
	"fmt"
	"net/http"
)

type AdvertisementHadler struct {
}

func (h *AdvertisementHadler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	segments := GetSegments(r.URL.Path)

	switch r.Method {
	case "GET":
		switch len(segments) {
		case 1:
			fmt.Fprintf(w, "All ads")
		case 2:
			fmt.Fprintf(w, "Single ads")
		}
	}
}
