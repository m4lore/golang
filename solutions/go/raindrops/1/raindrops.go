package raindrops

import "strconv"

func Convert(number int) string {
	var raindrops string 
    
	if isDivisibleBy3(number) { raindrops += "Pling" }
    if isDivisibleBy5(number) { raindrops += "Plang" }
    if isDivisibleBy7(number) { raindrops += "Plong" }

	if raindrops == "" {
		return strconv.Itoa(number)
	}
    
    return raindrops
}

func isDivisibleBy3 (number int) bool {
    return number % 3 == 0
}

func isDivisibleBy5 (number int) bool {
    return number % 5 == 0
}

func isDivisibleBy7 (number int) bool {
    return number % 7 == 0
}