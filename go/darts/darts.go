package darts

const (
	InnerR  = 1
	MiddleR = 5 * 5
	OuterR  = 10 * 10
)

func Score(x, y float64) int {
	distance := x*x + y*y
	switch {
	case distance <= InnerR:
		return 10
	case distance <= MiddleR:
		return 5
	case distance <= OuterR:
		return 1
	default:
		return 0
	}
}
