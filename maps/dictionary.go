package maps

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	return d[term], nil
}
