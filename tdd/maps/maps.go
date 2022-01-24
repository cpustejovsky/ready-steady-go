package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist to be update")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(k string) (string, error) {
	v, ok := d[k]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)
	switch err {
	case ErrNotFound:
		d[k] = v
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[k] = v
	default:
		return err
	}
	d[k] = v
	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}