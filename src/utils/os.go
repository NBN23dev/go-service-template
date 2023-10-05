package utils

import (
	"fmt"
	"os"
	"strconv"
)

// Value
type Value interface {
	string | int | int8 | int16 | int32 | int64
}

// ErrNotFound
func ErrNotFound(key string) error {
	return fmt.Errorf("libs: env var %s not found", key)
}

// GetEnv
func GetEnv[V Value](key string) (V, error) {
	val, exists := os.LookupEnv(key)

	var value V

	switch any(value).(type) {
	case string:
		if !exists {
			return any(val).(V), ErrNotFound(key)
		}

		return any(val).(V), nil
	case int:
		if !exists {
			return any(0).(V), ErrNotFound(key)
		}

		iVal, err := strconv.Atoi(val)

		return any(iVal).(V), err
	}

	return any(nil).(V), nil
}

// GetEnvOr
func GetEnvOr[V Value](key string, def V) (V, error) {
	val, exists := os.LookupEnv(key)

	var value V

	switch any(value).(type) {
	case string:
		if !exists {
			return any(def).(V), nil
		}

		return any(val).(V), nil
	case int:
		if !exists {
			return any(def).(V), nil
		}

		iVal, err := strconv.Atoi(val)

		if err != nil {
			return any(nil).(V), err
		}

		return any(iVal).(V), err
	}

	return any(nil).(V), nil
}
