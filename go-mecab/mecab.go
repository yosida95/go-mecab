package mecab

/*
#cgo darwin CFLAGS: -I/usr/local/Cellar/mecab/0.996/include
#cgo darwin LDFLAGS: -L/usr/local/Cellar/mecab/0.996/lib -lmecab -lstdc++

#include <stdlib.h>
#include <mecab.h>
*/
import "C"
import "runtime"
import "unsafe"

type Tagger struct {
	tagger *C.mecab_t
}

func init() {
	runtime.LockOSThread()
}

func NewTagger(arg string) (*Tagger, error) {
	argc := C.CString(arg)
	defer C.free(unsafe.Pointer(argc))

	tagger, err := C.mecab_new2(argc)
	if err != nil {
		return nil, err
	}

	return &Tagger{tagger}, nil
}

func (t *Tagger) Parse(str string) string {
	p := C.CString(str)
	defer C.free(unsafe.Pointer(p))

	r := C.mecab_sparse_tostr(t.tagger, p)
	defer C.free(unsafe.Pointer(r))

	result := C.GoString(r)
	return result
}

func (t *Tagger) ParseToNode(arg string) *Node {
	argc := C.CString(arg)
	defer C.free(unsafe.Pointer(argc))

	var convert func(prev *Node, node *C.struct_mecab_node_t) *Node
	convert = func(prev *Node, node *C.struct_mecab_node_t) *Node {
		if node == nil {
			return nil
		}

		self := new(Node)
		self.Prev = prev
		self.Next = convert(self, node.next)

		self.Length = int(node.length)
		self.RLength = int(node.rlength)
		self.Id = int(node.id)
		self.RcAttr = int(node.rcAttr)
		self.LcAttr = int(node.lcAttr)
		self.PosId = int(node.posid)
		self.Char_type = int(node.char_type)
		self.Stat = int(node.stat)
		self.Isbest = int(node.isbest)

		self.Alpha = float32(node.alpha)
		self.Beta = float32(node.beta)
		self.Prob = float32(node.prob)

		self.WCost = int(node.wcost)
		self.Cost = int(node.cost)
		self.Surface = string([]byte(C.GoString(node.surface))[:self.Length])
		self.Feature = C.GoString(node.feature)
		return self
	}

	_node := C.mecab_sparse_tonode(t.tagger, argc)
	defer C.free(unsafe.Pointer(_node))

	node := new(Node)
	node.Prev = nil
	node.Next = convert(node, _node.next)

	node.Length = int(_node.length)
	node.RLength = int(_node.rlength)
	node.Id = int(_node.id)
	node.RcAttr = int(_node.rcAttr)
	node.LcAttr = int(_node.lcAttr)
	node.PosId = int(_node.posid)
	node.Char_type = int(_node.char_type)
	node.Stat = int(_node.stat)
	node.Isbest = int(_node.isbest)

	node.Alpha = float32(_node.alpha)
	node.Beta = float32(_node.beta)
	node.Prob = float32(_node.prob)

	node.WCost = int(_node.wcost)
	node.Cost = int(_node.cost)
	node.Surface = string([]byte(C.GoString(_node.surface))[:node.Length])
	node.Feature = C.GoString(_node.feature)

	return node
}
