package atomic

type Value struct {
	c chan interface{}
}
