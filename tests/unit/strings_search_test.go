package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/service"
	"github.com/stretchr/testify/assert"
)

type SearchTestPair struct {
	query    string
	expected int
}

const indexErrorMsg = "Index is not correct."

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
		assert.Equal(t, actual, elem.expected, indexErrorMsg)
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
		assert.Equal(t, actual, elem.expected, indexErrorMsg)
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
		assert.Equal(t, actual, elem.expected, indexErrorMsg)
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
		assert.Equal(t, actual, elem.expected, indexErrorMsg)
	}
}
