package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("keyword not found")

func (d Dictionary) Search(key string) (string, error) {
	data, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return data, nil
}

func (d Dictionary) Add(key, data string) {
	d[key] = data
}
