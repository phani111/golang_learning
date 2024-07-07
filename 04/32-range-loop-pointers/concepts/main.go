package concepts

type Foo struct{}

type Store struct {
	m map[string]*Foo
}

func (s Store) Put(id string, foo *Foo) {
	s.m[id] = foo
}

func updateMapValue(mapValue map[string]LargeStruct, id string) {
	value := mapValue[id]
	value.foo = "bar"
	mapValue[id] = value
}

type LargeStruct struct {
	foo string
	_   [1024]int64
}
