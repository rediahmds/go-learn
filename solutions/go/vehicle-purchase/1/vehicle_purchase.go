package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
        return true
    }

    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    var car string
	if option1 < option2 {
        car = option1
    } else {
        car = option2
    }

    return car + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// age < 3.0, sell 80%
    if age < 3 {
        return 0.8 * originalPrice
    }
    
    // age > 10, sell 50%
    if age >= 10 {
        return 0.5 * originalPrice
    }
    
    // 3 =< age < 10, sell 70%
    return 0.7 * originalPrice
}
