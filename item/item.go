package item

type Type string

const (
	TypeA Type = "A"
	TypeB Type = "B"
)

type Request struct {
	Types *[]Type
}
