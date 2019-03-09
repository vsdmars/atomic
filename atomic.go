package atomic

func NewValue() value {
	return value{
		v: make(chan interface{}, 1),
	}
}

func (v *value) Set(val interface{}) {
	if val == nil {
		panic("atomic: store of nil value into value")
	}

	v.v <- val
}

func (v *value) Get() interface{} {
	select {
	case val := <-v.v:
		return val
	default:
		return nil
	}
}

func (v *value) ChanGet() <-chan interface{} {
	return v.v
}
