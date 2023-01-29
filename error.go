package msgpack

import "fmt"

type TypeError struct {
	want ElementType
	got  ElementType
}

// Error implements error
func (t *TypeError) Error() string {
	return fmt.Sprintf("type error want %s, got %s", t.want, t.got)
}

var _ error = &TypeError{}
