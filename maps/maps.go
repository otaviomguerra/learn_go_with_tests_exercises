package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound       = DictionaryErr("could not find the word you were looking for.")
	ErrWordExists     = DictionaryErr("cannot add word because it already exists.")
	ErrNotFoundUpdate = DictionaryErr("Could not find word to update.")
)

//Search searches a key in a Dictionary.
func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

//Add adds a key and value to Dictionary.
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

//Update updates a value of a given Dictionary key
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrNotFoundUpdate
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}
}

//Delete deletes a key/value of a given Dictionary.
func (d Dictionary) Delete(key string) {
	delete(d, key)
}
