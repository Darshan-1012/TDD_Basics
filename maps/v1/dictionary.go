package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("keyword not found")

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
