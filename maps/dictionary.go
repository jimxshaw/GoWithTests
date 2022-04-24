package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("search term cannot be found")

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term, definition string) {
	d[term] = definition
}
