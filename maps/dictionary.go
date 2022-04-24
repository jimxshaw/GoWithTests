package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("search term cannot be found")
	ErrTermExists = errors.New("cannot add term because it already exists")
)

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		d[term] = definition
	case nil:
		return ErrTermExists
	default:
		return err
	}

	return nil
}
