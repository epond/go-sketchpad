package channel

type Baz struct {
	Y string
}

func (b Baz) X() string {
	return "burp"
}

type Foo interface {
	X() string
}

func DoIt(stuff <-chan Foo) string {
	thing := <-stuff
	return thing.X()
}
