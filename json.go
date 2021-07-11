package mpp

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type JSON struct {
}

func NewJSON() *JSON {
	return &JSON{}
}

func (j *JSON) Do(in interface{}) (interface{}, error) {
	buf, err := json.Marshal(in)
	if err != nil {
		return nil, errors.Wrap(err, `failed to marshal source into JSON`)
	}
	return buf, nil
}
