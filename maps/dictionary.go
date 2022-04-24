package maps

type Dictionary map[string]string

func (d Dictionary) Search(dictonary map[string]string, term string) string {
	return dictonary[term]
}
