package model

import (
	"encoding/json"
	"io"
)

type Todo struct {
	ID     int    `json:"id"`
	Value  string `json:"value"`
	Marked int    `json:"marked"`
}

type Todos []*Todo

func (todo *Todo) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(todo)
}
