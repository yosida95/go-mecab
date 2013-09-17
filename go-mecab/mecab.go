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

	_node := C.mecab_sparse_tonode(t.tagger, argc)
	defer C.free(unsafe.Pointer(_node))

	return NodeFromCmecab_node_t(nil, _node)
}
