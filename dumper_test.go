package dumper

import (
	"log"
	"testing"
)

func TestD(t *testing.T) {
	t.Run("dump map", func(t *testing.T) {
		m := map[string]interface{}{
			"boo": "string",
		}
		D(&m)
		D(m)
	})

	t.Run("dump struct", func(t *testing.T) {
		type S struct {
			String    string
			Int       int
			Structure struct{}
		}
		s := S{Int: 1, String: "test", Structure: struct{}{}}
		D(s)
		D(&s)
	})

	t.Run("dump chan", func(t *testing.T) {
		var c = make(chan string)
		D(&c)
		D(c)
	})

	t.Run("dump func", func(t *testing.T) {
		var f = func() {
			log.Println("test")
		}
		D(&f)
		D(f)
	})

	t.Run("dump int", func(t *testing.T) {
		i := 1
		D(&i)
		D(i)
	})

	t.Run("dump string", func(t *testing.T) {
		s := "hello"
		D(&s)
		D(s)
	})

	t.Run("dump boolean", func(t *testing.T) {
		b := true
		D(&b)
		D(b)
	})

	t.Run("dump nil", func(t *testing.T) {
		D(nil)
	})

	t.Run("dump few values", func(t *testing.T) {
		D(1, "string", true, struct{}{})
	})

	t.Run("dump few values", func(t *testing.T) {
		defer func() {
			_ = recover()
		}()
		DD(1, "string", true, struct{}{})
	})
}
