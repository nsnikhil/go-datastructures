package trie

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateNewTrie(t *testing.T) {
	actualResult := NewTrie()
	expectedResult := &Trie{newNode()}

	assert.Equal(t, actualResult, expectedResult)
}

func TestTrieInsert(t *testing.T) {
	testCases := []struct {
		name          string
		actualResult  func() error
		expectedError error
	}{
		{
			name: "insert one word into trie",
			actualResult: func() error {
				tr := NewTrie()
				return tr.Insert("test")
			},
		},
		{
			name: "insert two word into trie",
			actualResult: func() error {
				tr := NewTrie()
				err := tr.Insert("test")
				if err != nil {
					return err
				}
				return tr.Insert("other")
			},
		},
		{
			name: "insert two unicode into trie",
			actualResult: func() error {
				tr := NewTrie()
				return tr.Insert("%ìǗΨԹ")
			},
		},
		{
			name: "insert failed when root is nil",
			actualResult: func() error {
				tr := &Trie{}
				return tr.Insert("test")
			},
			expectedError: errors.New("root is nil"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedError, testCase.actualResult())
		})
	}
}

func TestTrieSearchWord(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when whole word is present",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchWord("test")
			},
			expectedResult: true,
		},
		{
			name: "test return true when whole word ends",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))
				require.NoError(t, tr.Insert("tester"))

				return tr.SearchWord("test")
			},
			expectedResult: true,
		},
		{
			name: "test return true all the words are present",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))
				require.NoError(t, tr.Insert("other"))

				return tr.SearchWord("test") && tr.SearchWord("other")
			},
			expectedResult: true,
		},
		{
			name: "test return true for unicode characters",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("%ìǗΨԹ"))

				return tr.SearchWord("%ìǗΨԹ")
			},
			expectedResult: true,
		},
		{
			name: "test return false when word is not present",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchWord("other")
			},
			expectedResult: false,
		},
		{
			name: "test return false when searching input prefix",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchWord("tes")
			},
			expectedResult: false,
		},
		{
			name: "test return false search term is longer",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchWord("testt")
			},
			expectedResult: false,
		},
		{
			name: "test return when root is nil",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))
				tr.root = nil

				return tr.SearchWord("test")
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestTrieSearchPrefix(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() bool
		expectedResult bool
	}{
		{
			name: "test return true when prefix is present",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchPrefix("test")
			},
			expectedResult: true,
		},
		{
			name: "test return true all the words prefix are present",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))
				require.NoError(t, tr.Insert("other"))

				return tr.SearchPrefix("test") && tr.SearchPrefix("other")
			},
			expectedResult: true,
		},
		{
			name: "test return true when prefix is also the full word",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchPrefix("test")
			},
			expectedResult: true,
		},
		{
			name: "test return true when searching for unicode character prefix",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("%ìǗΨԹ"))

				return tr.SearchPrefix("%ì")
			},
			expectedResult: true,
		},
		{
			name: "test return false when prefix is not present",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchPrefix("other")
			},
			expectedResult: false,
		},
		{
			name: "test return false search term is longer",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))

				return tr.SearchPrefix("testt")
			},
			expectedResult: false,
		},
		{
			name: "test return false when root is nil",
			actualResult: func() bool {
				tr := NewTrie()
				require.NoError(t, tr.Insert("test"))
				tr.root = nil

				return tr.SearchPrefix("tes")
			},
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expectedResult, testCase.actualResult())
		})
	}
}

func TestTrieGet(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() []string
		expectedResult map[string]bool
		expectedCount  int
	}{
		{
			name: "test get all words against the prefix m",
			actualResult: func() []string {
				tr := NewTrie()
				data := []string{"mobile", "mouse", "mousepad", "money", "monitor", "matter"}
				for _, d := range data {
					require.NoError(t, tr.Insert(d))
				}

				return tr.Get("m")
			},
			expectedResult: map[string]bool{
				"mobile": true, "mouse": true, "mousepad": true, "money": true, "monitor": true, "matter": true,
			},
			expectedCount: 6,
		},
		{
			name: "test get all words against the prefix mo",
			actualResult: func() []string {
				tr := NewTrie()
				data := []string{"mobile", "mouse", "mousepad", "money", "monitor", "matter"}
				for _, d := range data {
					require.NoError(t, tr.Insert(d))
				}

				return tr.Get("mo")
			},
			expectedResult: map[string]bool{
				"mobile": true, "mouse": true, "mousepad": true, "money": true, "monitor": true, "matter": false,
			},
			expectedCount: 5,
		},
		{
			name: "test get all words against the prefix mou",
			actualResult: func() []string {
				tr := NewTrie()
				data := []string{"mobile", "mouse", "mousepad", "money", "monitor", "matter"}
				for _, d := range data {
					require.NoError(t, tr.Insert(d))
				}

				return tr.Get("mou")
			},
			expectedResult: map[string]bool{
				"mobile": false, "mouse": true, "mousepad": true, "money": false, "monitor": false, "matter": false,
			},
			expectedCount: 2,
		},
		{
			name: "test get all words against the prefix mous",
			actualResult: func() []string {
				tr := NewTrie()
				data := []string{"mobile", "mouse", "mousepad", "money", "monitor", "matter"}
				for _, d := range data {
					require.NoError(t, tr.Insert(d))
				}

				return tr.Get("mous")
			},
			expectedResult: map[string]bool{
				"mobile": false, "mouse": true, "mousepad": true, "money": false, "monitor": false, "matter": false,
			},
			expectedCount: 2,
		},
		{
			name: "test get all words against the prefix mouse",
			actualResult: func() []string {
				tr := NewTrie()
				data := []string{"mobile", "mouse", "mousepad", "money", "monitor", "matter"}
				for _, d := range data {
					require.NoError(t, tr.Insert(d))
				}

				return tr.Get("mouse")
			},
			expectedResult: map[string]bool{
				"mobile": false, "mouse": true, "mousepad": true, "money": false, "monitor": false, "matter": false,
			},
			expectedCount: 2,
		},
		{
			name: "test get return empty list for prefix oth",
			actualResult: func() []string {
				tr := NewTrie()
				data := []string{"mobile", "mouse", "mousepad", "money", "monitor", "matter"}
				for _, d := range data {
					require.NoError(t, tr.Insert(d))
				}

				return tr.Get("oth")
			},
			expectedResult: map[string]bool{},
			expectedCount:  0,
		},
		{
			name: "test return empty list when root is null",
			actualResult: func() []string {
				tr := NewTrie()
				require.NoError(t, tr.Insert("data"))

				tr.root = nil

				return tr.Get("oth")
			},
			expectedResult: map[string]bool{},
			expectedCount:  0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, testCase.expectedCount, len(res))

			if testCase.expectedCount != 0 {
				for _, r := range res {
					require.True(t, testCase.expectedResult[r])
				}
			}
		})
	}
}
