package http

import (
	"encoding/json"
	"time"
)

type TaskDTO struct {
	Title      string
	Descripion string
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO)ToString()string{
	  b, err:= json.MarshalIndent(e,"","    ")
	  if err !=nil{
		panic(err)
	  }
	  return string(b)
}