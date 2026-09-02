package darts

import "math"

type ring struct {
	radius float64
	points int
}

var rings = []ring{
	{radius: 1, points: 10},
	{radius: 5, points: 5},
	{radius: 10, points: 1},
}

func Score(x, y float64) int {
	d := math.Hypot(x, y)
	for _, r := range rings {
		if d <= r.radius {
			return r.points
		}
	}
	return 0
}