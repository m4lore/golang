package purchase

import "slices"

func NeedsLicense(kind string) bool {
    lk := []string{"car", "truck"}
	if slices.Contains(lk, kind) { return true }
    return false
}

func ChooseVehicle(option1, option2 string) string {
    const def string = " is clearly the better choice."
    if option1 < option2 { return option1 + def }
    return option2 + def
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
    case age >= 3 && age < 10:
        return originalPrice * 0.7
    case age >= 10:
        return originalPrice * 0.5
    default: 
    	return originalPrice * 0.8
    }
}
