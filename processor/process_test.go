package processor

import (
	"testing"

	"github.com/venwex/cli-project.git/pkg/utils"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test 1",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			words := utils.Tokenize(test.input)
			words = utils.NormalizeCommands(words)

			words = Process(words)

			words = utils.Article(words)

			result := utils.Format(utils.Join(words))

			if result != test.expected {
				t.Errorf("expected:\n%s\n\ngot:\n%s", test.expected, result)
			}
		})
	}
}
