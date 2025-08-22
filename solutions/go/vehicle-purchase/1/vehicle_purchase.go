package purchase

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
        return option2+" is clearly the better choice."
    }
    return option1+" is clearly the better choice."
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3.00 {
        return originalPrice*0.8
    } else if age >= 3.00 && age < 10.00 {
        return originalPrice*0.7
    }
    return originalPrice*0.5
}
