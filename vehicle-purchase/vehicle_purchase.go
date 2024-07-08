package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var bestOption string
	if option1 < option2 {
		bestOption = option1
	} else {
		bestOption = option2
	}

	return fmt.Sprintf("%s %s", bestOption, "is clearly the better choice.")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var discount float64
	if age < 3 {
		discount = 0.8
	} else if age < 10 {
		discount = 0.7
	} else {
		discount = 0.5
	}

	return originalPrice * discount
}