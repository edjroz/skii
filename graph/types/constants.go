package types

const (
	blue  float64 = 1
	red   float64 = 2
	black float64 = 3
)

// DiffcultyConverter - converts colors to floats that can be used for weights
func DifficultyConverter(s string) float64 {
	var difficulty float64
	switch s {
	case "blue":
		difficulty = blue
	case "red":
		difficulty = red
	case "black":
		difficulty = black
	}
	return difficulty
}
