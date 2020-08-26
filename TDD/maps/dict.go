package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("the word you were trying to update does not exist")
)

type DictionaryErr string
type dictionaryMethod func(string)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	val, ok := d[word]
	if !ok {
		return val, ErrNotFound
	}
	return d[word], nil
}

func (d Dictionary) checkForExistence(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Add(key, val string) error {
	err := d.checkForExistence(key)
	if err == nil {
		d[key] = val
		return nil
	}
	return err
}

func (d Dictionary) Update(word, definition string) error {
	err := d.checkForExistence(word)
	if err == nil {
		d[word] = definition
		return nil
	}
	return err
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}
