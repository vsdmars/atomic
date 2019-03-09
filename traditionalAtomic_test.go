package atomic_test

import (
	"fmt"
	"testing"

	va "github.com/vsdmars/atomic"
)

func TestValue(t *testing.T) {
	v := va.NewValue()

	t.Run("Value test", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {
			t.Parallel()
			for range (*[100000]struct{})(nil) {
				v.Set(42)
				r := v.Get().(int)
				fmt.Printf("Test1 r:%v\n", r)
				if r == 42 {
					continue
				} else {
					t.Errorf("r value: %v", r)
				}
			}
		})

		t.Run("Test2", func(t *testing.T) {
			t.Parallel()
			for range (*[100000]struct{})(nil) {
				v.Set(43)
				r := v.Get().(int)
				fmt.Printf("Test2 r:%v\n", r)
				if r == 43 {
					continue
				} else {
					t.Errorf("r value: %v", r)
				}
			}
		})

		t.Run("Test3", func(t *testing.T) {
			t.Parallel()
			for range (*[100000]struct{})(nil) {
				v.Set(44)
				r := v.Get().(int)
				fmt.Printf("Test3 r:%v\n", r)
				if r == 44 {
					continue
				} else {
					t.Errorf("r value: %v", r)
				}
			}
		})
	})
}

func TestValueChannel(t *testing.T) {
	v := va.NewValue()

	t.Run("Value Channel test", func(t *testing.T) {
		t.Run("Test1", func(t *testing.T) {
			t.Parallel()
			for range (*[100000]struct{})(nil) {
				v.Set(42)
				if r := <-v.ChanGet(); r.(int) != 42 {
					t.Errorf("r value: %v", r)
				}
			}
		})

		t.Run("Test2", func(t *testing.T) {
			t.Parallel()
			for range (*[100000]struct{})(nil) {
				v.Set(43)
				if r := <-v.ChanGet(); r.(int) != 43 {
					t.Errorf("r value: %v", r)
				}
			}
		})

		t.Run("Test3", func(t *testing.T) {
			t.Parallel()
			for range (*[100000]struct{})(nil) {
				v.Set(44)
				if r := <-v.ChanGet(); r.(int) != 44 {
					t.Errorf("r value: %v", r)
				}
			}
		})
	})
}
