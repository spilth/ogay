package ogay

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWordsStartingWithAVowel(t *testing.T) {
	assert.Equal(t,"appleay", TranslateWord("apple"))
	assert.Equal(t, "elixiray", TranslateWord("elixir"))
	assert.Equal(t,"Appleay", TranslateWord("Apple"))
	assert.Equal(t, "Elixiray", TranslateWord("Elixir"))
}

func TestWordsStartingWithASingleConsonant(t *testing.T) {
	assert.Equal(t, "avajay", TranslateWord("java"))
	assert.Equal(t, "ubyray", TranslateWord("ruby"))
	assert.Equal(t, "Avajay", TranslateWord("Java"))
	assert.Equal(t, "Ubyray", TranslateWord("Ruby"))
}

func TestWordsStartingWithMultipleConsonants(t *testing.T) {
	assert.Equal(t, "ogrammingpray", TranslateWord("programming"))
	assert.Equal(t, "ingstray", TranslateWord("string"))
	assert.Equal(t, "Ogrammingpray", TranslateWord("Programming"))
	assert.Equal(t, "Ingstray", TranslateWord("String"))
}
