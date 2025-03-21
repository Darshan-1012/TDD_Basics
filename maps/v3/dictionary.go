package maps

type Dictionary map[string]string

const (
	ErrNotFound          = DictionaryErr("keyword not found")
	ErrWordExists        = DictionaryErr("keyword already exists")
	ErrWordDoesNotExists = DictionaryErr("keyword doesn't exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	data, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return data, nil
}

func (d Dictionary) Add(key, data string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = data
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, data string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[key] = data
	default:
		return err
	}
	return nil
}
