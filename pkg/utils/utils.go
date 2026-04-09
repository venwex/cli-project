package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	punctuations := []string{",", ".", "!", "?", ":", ";"}

	for _, p := range punctuations {
		text = strings.ReplaceAll(text, p, " "+p+" ")
	}

	return strings.Fields(text)
}

func NormalizeCommands(tokens []string) []string {
	var result []string

	for i := 0; i < len(tokens); i++ {
		// проверка на: (up , 3) / (low , 3) / (cap , 3)
		if i+2 < len(tokens) && tokens[i+1] == "," && strings.HasSuffix(tokens[i+2], ")") && (tokens[i] == "(up" || tokens[i] == "(low" || tokens[i] == "(cap") {
			cmd := tokens[i] + ", " + strings.TrimSuffix(tokens[i+2], ")") + ")"
			result = append(result, cmd)
			i += 2
			continue
		}

		result = append(result, tokens[i])
	}

	return result
}

func Remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func Parse(s string) (int, error) {
	s = strings.TrimPrefix(s, "(up, ")
	s = strings.TrimPrefix(s, "(low, ")
	s = strings.TrimPrefix(s, "(cap, ")

	s = strings.TrimSuffix(s, ")")

	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func Hex(s string) (string, error) {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(n, 10), nil
}

func Bin(s string) (string, error) {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(n, 10), nil
}

func Up(s string) string {
	return strings.ToUpper(s)
}

func Low(s string) string {
	return strings.ToLower(s)
}

func Cap(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func Format(text string) string {
	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")
	return text
}

func Article(words []string) []string {

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			next := strings.Trim(words[i+1], ".,!?:;'")
			next = strings.ToLower(next)

			if strings.HasPrefix(next, "a") ||
				strings.HasPrefix(next, "e") ||
				strings.HasPrefix(next, "i") ||
				strings.HasPrefix(next, "o") ||
				strings.HasPrefix(next, "u") ||
				strings.HasPrefix(next, "h") {

				words[i] = "an"
			}

		}
	}

	return words
}

func Join(words []string) string {
	var b strings.Builder
	inQuote := false

	for i, word := range words {

		if word == "'" {
			if !inQuote && i > 0 {
				b.WriteByte(' ')
			}

			b.WriteString("'")
			inQuote = !inQuote
			continue
		}

		if i > 0 && words[i-1] != "'" {
			b.WriteByte(' ')
		}

		if i > 0 && words[i-1] == "'" && !inQuote {
			b.WriteByte(' ')
		}

		b.WriteString(word)
	}

	return b.String()
}
