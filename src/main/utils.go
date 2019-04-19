package main

import (
	"encoding/json"
	"errors"
)

func sliceToString(slice []string) (string, error) {
	if slice != nil {
		str, err := json.Marshal(slice)
		if err != nil {
			return string(str), errors.New("utils: Failed to marshal slice")
		}
		return string(str), nil
	}
	return "", nil
}

func stringToSlice(str string) ([]string, error) {
	if str != "" {
		var slice []string
		err := json.Unmarshal([]byte(str), &slice)
		if err != nil {
			return slice, errors.New("utils: Failed to unmarshal slice")
		}
		return slice, nil
	}
	return []string{}, nil
}
