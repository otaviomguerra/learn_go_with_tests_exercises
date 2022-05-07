package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

//Search searches a key in a Dictionary.
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}
