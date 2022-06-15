package temp

import "fmt"

type GetterFunc func(string) (any, bool)

func (f *GetterFunc) Get(k string) (any, bool) {
	getterFunc := *f
	return getterFunc(k)
}

type Getter interface {
	Get(string) (interface{}, bool)
}

type Cache struct {
	GetterFunc
}

func GGG(g Getter) {
	fmt.Println(g("abc"))
}

func tempmain() {
	c := Cache{}
	GGG(&c)
}
