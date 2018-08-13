package ogay

import (
	"strings"
	"unicode"
)

func TranslateWord(englishWord string) (pigLatinWord string) {
	firstVowelIndex := getFirstVowelIndex(englishWord)

	beginning := englishWord[0:firstVowelIndex]
	ending := englishWord[firstVowelIndex:]

	pigLatinWord = ending + strings.ToLower(beginning) + "ay"

	if unicode.IsUpper(rune(englishWord[0])) {
		pigLatinWord = strings.ToUpper(string(pigLatinWord[0])) + pigLatinWord[1:]
	}

	return
}

func getFirstVowelIndex(englishWord string) int {
	for index, character := range englishWord {
		if isVowel(character) {
			return index
		}
	}
	return 0
}

func isVowel(character rune) bool {
	vowels := "AEIOUaeiou"

	for _, vowel := range vowels {
		if character == vowel {
			return true
		}
	}

	return false
}
