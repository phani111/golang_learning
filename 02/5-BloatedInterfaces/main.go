package main

// Bad interface, mixed all the methods
type GenericInterface interface {
	A() int
	B() string
	C() ([]byte, error)
}

// Better usage, smaller the interface better the readability

type DoerA interface {
	A() int
}

type DoerB interface {
	B() string
}

type DoerC interface {
	c() ([]byte, error)
}

func main() {

}
