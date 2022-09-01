package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/service"
)

type SearchTestPair struct {
	query    string
	expected int
}

// TestSearchNormal calls service.StringArraySearch
// with an NORMAL array and string query, checking
// for a valid return value.
func TestSearchNormal(t *testing.T) {
	array := []string{
		"kirill",
		"james",
		"chris",
		"charlie",
		"chloe",
	}
	testPairs := []SearchTestPair{
		{"kirill", 0},
		{"chris", 2},
		{"chloe", 4},
	}
	for _, elem := range testPairs {
		actual := service.StringArraySearch(array, elem.query)
		if elem.expected != actual {
			t.Errorf(`Index was incorrect, actual: %d, expected: %d.`, actual, elem.expected)
		}
	}
}

// TestSearchNormal calls service.StringArraySearch
// with an EMPTY array and string query, checking
// for a valid return value.
func TestSearchEmpty(t *testing.T) {
	array := []string{}
	testPairs := []SearchTestPair{
		{"kirill", -1},
		{"chris", -1},
		{"chloe", -1},
	}
	for _, elem := range testPairs {
		actual := service.StringArraySearch(array, elem.query)
		if elem.expected != actual {
			t.Errorf(`Index was incorrect, actual: %d, expected: %d.`, actual, elem.expected)
		}
	}
}

// TestSearchNormal calls service.StringArraySearch
// with an NORMAL array and string query thats have to be not founded,
// checking for a valid return value.
func TestSearchNotFound(t *testing.T) {
	array := []string{
		"kirill",
		"james",
		"chris",
		"charlie",
		"chloe",
	}
	testPairs := []SearchTestPair{
		{"johnson", -1},
		{"", -1},
		{"    ", -1},
	}
	for _, elem := range testPairs {
		actual := service.StringArraySearch(array, elem.query)
		if elem.expected != actual {
			t.Errorf(`Index was incorrect, actual: %d, expected: %d.`, actual, elem.expected)
		}
	}
}

// TestSearchNormal calls service.StringArraySearch
// with an NIL array and string query, checking
// for a valid return value.
func TestSearchNilArray(t *testing.T) {
	testPairs := []SearchTestPair{
		{"johnson", -1},
		{"kirill", -1},
		{"chloe", -1},
	}
	for _, elem := range testPairs {
		actual := service.StringArraySearch(nil, elem.query)
		if elem.expected != actual {
			t.Errorf(`Index was incorrect, actual: %d, expected: %d.`, actual, elem.expected)
		}
	}
}
