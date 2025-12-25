package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title      string
	Descripion string
}





func (t TaskDTO)ValidateForCreater()error{
if t.Title == "" {
	return  errors.New("title is emply")
}

if t.Descripion=="" {
  return  errors.New("descripion is emply")
}
return  nil
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