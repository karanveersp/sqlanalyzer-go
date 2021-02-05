package analyzers

import (
	"reflect"
	"testing"
)

func TestAnalyzers(t *testing.T) {

	checkArraysEqual := func(t testing.TB, actual, expected []string) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf(`%v != %v`, actual, expected)
		}
	}

	t.Run("able to find single line comments", func(t *testing.T) {
		str := "This has\n" +
			"a\n" +
			"-- comment"
		expected := []string{"-- comment"}
		actual := FindAllSingleLineComments(str)
		checkArraysEqual(t, actual, expected)
	})

	t.Run("able to find many single line comments", func(t *testing.T) {
		str := "This has\n" +
			"a\n" +
			"-- comment\n" +
			"--AnotherLine!\n" +
			"And -- this should also be detected"
		expected := []string{"-- comment\n", "--AnotherLine!\n", "-- this should also be detected"}
		actual := FindAllSingleLineComments(str)
		checkArraysEqual(t, actual, expected)
	})

	t.Run("able to find multi line comments", func(t *testing.T) {
		str := "This has\n" +
			"a\n" +
			"/* multi\n" +
			"line\n" +
			"comment*/"
		expected := []string{"/* multi\nline\ncomment*/"}
		actual := FindAllMultiLineComments(str)
		checkArraysEqual(t, actual, expected)
	})
}
