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

func NewNode(prev, next *Node, surface, feature string, length, rlength, id, rcAttr, lcAttr, posid, char_type, stat, isbest int, alpha, beta, prob float32, wcost, cost int) *Node {
	return &Node{
		Prev: prev,
		Next: next,

		Surface: surface,

		Feature:   feature,
		Length:    length,
		RLength:   rlength,
		Id:        id,
		RcAttr:    rcAttr,
		LcAttr:    lcAttr,
		PosId:     posid,
		Char_type: char_type,
		Stat:      stat,
		Isbest:    isbest,

		Alpha: alpha,
		Beta:  beta,
		Prob:  prob,

		WCost: wcost,
		Cost:  cost}
}

func NodeFromCstruct_mecab_node_t(prev *Node, node *C.struct_mecab_node_t) *Node {
	if node == nil {
		return nil
	}

	n := NewNode(
		prev,
		nil,
		string([]byte(C.GoString(node.surface))[:int(node.length)]),
		C.GoString(node.feature),
		int(node.length),
		int(node.rlength),
		int(node.id),
		int(node.rcAttr),
		int(node.lcAttr),
		int(node.posid),
		int(node.char_type),
		int(node.stat),
		int(node.isbest),
		float32(node.alpha),
		float32(node.beta),
		float32(node.prob),
		int(node.wcost),
		int(node.cost))
	n.Next = NodeFromCstruct_mecab_node_t(n, node.next)
	return n
}

func NodeFromCmecab_node_t(prev *Node, node *C.mecab_node_t) *Node {
	if node == nil {
		return nil
	}

	n := NewNode(
		prev,
		nil,
		string([]byte(C.GoString(node.surface))[:int(node.length)]),
		C.GoString(node.feature),
		int(node.length),
		int(node.rlength),
		int(node.id),
		int(node.rcAttr),
		int(node.lcAttr),
		int(node.posid),
		int(node.char_type),
		int(node.stat),
		int(node.isbest),
		float32(node.alpha),
		float32(node.beta),
		float32(node.prob),
		int(node.wcost),
		int(node.cost))
	n.Next = NodeFromCstruct_mecab_node_t(n, node.next)
	return n
}
