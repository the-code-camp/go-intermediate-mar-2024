package model

type Category int

const (
	FOOD Category = iota
	BOOKS
	BEVERAGE
	TOYS
)

func (c Category) String() string {
	switch c {
	case FOOD:
		return "FOOD"
	case BEVERAGE:
		return "BABY"
	case BOOKS:
		return "BOOKS"
	case TOYS:
		return "TOYS"
	}
	return "unknown"
}
