package processor

import (
	"log"
	"strings"

	"github.com/venwex/cli-project.git/pkg/utils"
)

func Process(words []string) []string {
	var err error

	for i := 0; i < len(words); i++ {
		w := words[i]
		switch {

		case w == "(hex)":
			if i > 0 {
				words[i-1], err = utils.Hex(words[i-1])
				if err != nil {
					log.Fatal(err)
				}
			}
			words = utils.Remove(words, i)
			i--

		case w == "(bin)":
			if i > 0 {
				words[i-1], err = utils.Bin(words[i-1])
				if err != nil {
					log.Fatal(err)
				}
			}
			words = utils.Remove(words, i)
			i--

		case w == "(up)":
			if i > 0 {
				words[i-1] = utils.Up(words[i-1])
			}
			words = utils.Remove(words, i)
			i--

		case w == "(low)":
			if i > 0 {
				words[i-1] = utils.Low(words[i-1])
			}
			words = utils.Remove(words, i)
			i--

		case w == "(cap)":
			if i > 0 {
				words[i-1] = utils.Cap(words[i-1])
			}
			words = utils.Remove(words, i)
			i--

		case strings.HasPrefix(w, "(up,") || strings.HasPrefix(w, "(low,") || strings.HasPrefix(w, "(cap,"):

			n, err := utils.Parse(w)
			if err != nil {
				log.Fatal(err)
			}

			start := i - n
			if start < 0 {
				start = 0
			}

			for j := i - 1; j >= start; j-- {
				if strings.HasPrefix(w, "(up") {
					words[j] = utils.Up(words[j])
				} else if strings.HasPrefix(w, "(low") {
					words[j] = utils.Low(words[j])
				} else {
					words[j] = utils.Cap(words[j])
				}
			}

			words = utils.Remove(words, i)
			i--
		}
	}

	return words
}
