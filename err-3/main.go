package main

import (
	"err-3/networkerr"
	"err-3/worker"
	"errors"
	"log"
)

func main() {
	w := worker.New()

	if err := w.DoWork(); err != nil {
		var networkerr *networkerr.NetworkError
		switch {
		case errors.Is(err, worker.ErrNoInternet):
			log.Fatalf("check internet ")

		case errors.As(err, &networkerr):
			log.Fatalf("network worker error with code %d: %s", networkerr.Code, networkerr)
		default:
			log.Fatalf("unknown worker error: %s", err)
		}
	}

}
