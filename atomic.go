package atomic

func (v *Value) Set(val interface{}) {
	if v.c == nil {
		v.c = make(chan struct{})
	}

	v.c <- val
}

func (v *Value) Get() {
	return <-v.c
}
