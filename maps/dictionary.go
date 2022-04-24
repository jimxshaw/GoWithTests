package maps

const (
	ErrNotFound         = DictionaryErr("search term cannot be found")
	ErrTermExists       = DictionaryErr("cannot add term because it already exists")
	ErrTermDoesNotExist = DictionaryErr("cannot update term because it does not exist")
)

// Create this custom type that implements the
// error interface to make errors more
// reusable and immutable.
// https://dave.cheney.net/2016/04/07/constant-errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

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

func (d Dictionary) Update(term, newDefinition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		return ErrTermDoesNotExist
	case nil:
		d[term] = newDefinition
	default:
		return err
	}

	return nil
}
