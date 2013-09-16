package main

import (
	"../go-mecab/"
	"fmt"
	"log"
)

func main() {
	tagger, err := mecab.NewTagger("-d /usr/local/Cellar/mecab/0.996/lib/mecab/dic/ipadic")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tagger.Parse("すもももももももものうち"))
}
