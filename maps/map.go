package maps

const (
	// ErrNotFound when word has no def
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists when word exists
	ErrWordExists = DictionaryErr("word exists")
	// ErrWordDoesNotExist when word exists
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search func
func Search(dict map[string]string, word string) string {
	return dict[word]
}

// Dictionary map
type Dictionary map[string]string

// Search func
func (d Dictionary) Search(word string) (string, error) {
	result, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}

// Add func
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update func
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete func
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
