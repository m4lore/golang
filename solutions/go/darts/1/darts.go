package darts

import "math"

type radius struct {
    value    float64
	points   int 
}

var r1 = radius {
    value:	1.0,
    points:	10,
}
var r2 = radius {
    value:	5.0,
    points:	5,
}
var r3 = radius {
    value:	10.0,
    points:	1,
}

func Score(x, y float64) int {
	i := math.Sqrt(x * x + y * y)
    
    switch {
    case i <= r1.value:
        return r1.points
    case i <= r2.value:
        return r2.points
    case i <= r3.value:
        return r3.points
    default:
        return 0
    }
}


