package lasagna

func PreparationTime(layers []string,
	prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}
	return len(layers) * prepTimePerLayer
}

func Quantities(layers []string) (grams int,
	liters float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			grams += 50
		case "sauce":
			liters += 0.2
		default:
			continue
		}
	}
	return
}

func AddSecretIngredient(friendsList []string,
	myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(ogAmounts []float64,
	scale int) (amounts []float64) {
	amounts = make([]float64, len(ogAmounts))
	for i, v := range ogAmounts {
		amounts[i] = (v / 2.0) * float64(scale)
	}
	return
}
