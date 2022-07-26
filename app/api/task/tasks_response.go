package task

import (
	"encoding/json"
	"io"
)

func (tasksResponse TasksResponse) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(tasksResponse)
}
