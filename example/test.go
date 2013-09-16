package main

import (
	"../go-mecab/"
	"fmt"
	"log"
)

func main() {
	tagger, err := mecab.NewTagger("-Owakati")
	if err != nil {
		log.Fatalln(err)
	}

	for node := tagger.ParseToNode("すもももももももものうち"); node != nil; node = node.Next {
		fmt.Printf("%s\t%s\n", node.Surface, node.Feature)
	}
}
