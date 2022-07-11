package json

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON[T any] struct {
	inner T
}

func (j *JSON[T]) Get() T {
	return j.inner
}
func (j *JSON[T]) Set(v T) {
	j.inner = v
}

func (j JSON[T]) GormDataType() string {
	return "json"
}

func (j JSON[T]) Value() (driver.Value, error) {
	return json.Marshal(j.Get())
}

func (j *JSON[T]) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("[]byte assertion failed")
	}

	var v T
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	j.Set(v)

	return nil
}
