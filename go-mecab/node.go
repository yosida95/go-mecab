package mecab

/*
#cgo darwin CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
#cgo darwin LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++

#include <stdlib.h>
#include <mecab.h>
*/
import "C"

type Node struct {
	Prev *Node
	Next *Node

	Enext *Node
	Bnext *Node

	Surface string

	Feature   string
	Length    int
	RLength   int
	Id        int
	RcAttr    int
	LcAttr    int
	PosId     int
	Char_type int
	Stat      int
	Isbest    int

	Alpha float32
	Beta  float32
	Prob  float32

	WCost int
	Cost  int
}
