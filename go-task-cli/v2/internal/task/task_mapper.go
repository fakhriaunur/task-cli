package task

import (
	"encoding/json"
	"os"
)

var _ taskFileMapperPort = (*taskMapper)(nil)

type taskMapper struct{}

// decode implements taskFileMapperPort.
func (t *taskMapper) decode(file *os.File) (taskStructure, error) {
	var ts taskStructure
	// ts := taskStructure{
	// 	Tasks: map[int]task{},
	// }
	if err := json.NewDecoder(file).Decode(&ts); err != nil && err.Error() != "EOF" {
		return taskStructure{}, err
	}

	return ts, nil
}

// encode implements taskFileMapperPort.
func (t *taskMapper) encode(file *os.File, ts taskStructure) error {
	if err := json.NewEncoder(file).Encode(ts); err != nil {
		return err
	}

	return nil
}

func NewTaskMapper() taskFileMapperPort {
	return &taskMapper{}
}

// // decode implements taskMapperPort.
// func (t *taskMapper) decode(dat []byte) (taskStructure, error) {
// 	var ts taskStructure
// 	if err := json.Unmarshal(dat, &ts); err != nil {
// 		return taskStructure{}, err
// 	}

// 	return ts, nil
// }

// // encode implements taskMapperPort.
// func (t *taskMapper) encode(ts taskStructure) (dat []byte, err error) {
// 	return json.MarshalIndent(ts, "", "\t")
// }
