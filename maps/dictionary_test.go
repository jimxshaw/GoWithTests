package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("term exists", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("term doesn't exist", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new term", func(t *testing.T) {
		dictionary := Dictionary{}
		term := "test"
		definition := "this is a test"

		err := dictionary.Add(term, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, term, definition)
	})

	t.Run("existing term", func(t *testing.T) {
		term := "test"
		definition := "this is a test"
		dictionary := Dictionary{term: definition}

		err := dictionary.Add(term, "new test")

		assertError(t, err, ErrTermExists)
		assertDefinition(t, dictionary, term, definition)
	})

}

func TestUpdate(t *testing.T) {
	term := "test"
	definition := "this is a test"
	dictionary := Dictionary{term: definition}
	newDefinition := "new definition"

	dictionary.Update(term, newDefinition)

	assertDefinition(t, dictionary, term, newDefinition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, term, definition string) {
	t.Helper()

	got, err := dictionary.Search(term)
	if err != nil {
		t.Fatal("should find added term: ", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
