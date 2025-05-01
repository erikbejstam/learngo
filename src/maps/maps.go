package maps

type Dictionary map[string]string

const (
	ErrWordDoesNotExist = DictionaryError("word does not exist in dictionary")
	ErrNotFound         = DictionaryError("could not find word")
	ErrWordExists       = DictionaryError("word already exists in dictionary")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(find string) (string, error) {
	value, ok := d[find]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = newDefinition
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}
