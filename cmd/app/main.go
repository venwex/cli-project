package main

import (
	"fmt"
	"log"
	"os"

	"github.com/venwex/cli-project.git/pkg/utils"
	"github.com/venwex/cli-project.git/processor"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("the number of arguments is required")
	}

	input, output := os.Args[1], os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
	words := utils.Tokenize(string(data)) // problem: "word,". solution: "word ," and split then we got "word", ","

	words = utils.NormalizeCommands(words) // problem: "(cap", ",", "6)". solution: "(cap, 6)"

	words = processor.Process(words)

	words = utils.Article(words)

	result := utils.Format(utils.Join(words))
	err = os.WriteFile(output, []byte(result), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
