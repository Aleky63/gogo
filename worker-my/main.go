package main

import (
	"errors"
	"fmt"
	"log"
	"yourgo/worker"
)

func main() {

	if err := run(); err != nil {
		var fileNotExists *worker.FileNotExistsError
		if errors.As(err, &fileNotExists) {
			log.Fatalf("ccccccAAAAAcccccccccccccccccc")
		} else {
			log.Fatalf("worker error: %s ", err)
		}

	}
}

func run() error {
	w := worker.New()
	if err := w.DoWork("file-mymy.exe"); err != nil {
		return fmt.Errorf("do work: %w", err)
	}
	return nil
}
