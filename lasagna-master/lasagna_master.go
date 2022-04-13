package lasagna

func PreparationTime(layers []string, averageTime int) int {
	if averageTime == 0 {
		averageTime = 2
	}

	return len(layers) * averageTime
}

func Quantities(slices []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0

	for _, slice := range slices {
		if slice == "noodles" {
			noodles += 50
		}

		if slice == "sauce" {
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	frindListLastElement := len(friendsList) - 1
	myListLastElement := len(myList) - 1
	myList[myListLastElement] = friendsList[frindListLastElement]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	currentPortions := 2

	newQuantities := make([]float64, len(quantities))
	for i, v := range quantities {
		newQuantities[i] = v / float64(currentPortions) * float64(portions)
	}

	return newQuantities
}
