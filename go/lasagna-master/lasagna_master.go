package lasagna

// PreparationTime calculates the time needed to prepare a lasagna
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities calculates the ammount of noodles and sauce needed
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds the last ingredient of a friends list to the current list
func AddSecretIngredient(friendsList []string, myList []string) {
	if len(friendsList) > 0 && len(myList) > 0 && myList[len(myList)-1] == "?" {
		myList[len(myList)-1] = friendsList[len(friendsList)-1]
	}
}

// ScaleRecipe scales the ammounts of a recipe to a given number of portions
func ScaleRecipe(ammounts []float64, portions int) []float64 {
	scaledAmmounts := make([]float64, len(ammounts))
	if len(ammounts) > 0 {
		for i, ammount := range ammounts {
			scaledAmmounts[i] = (ammount * float64(portions)) / 2
			}
		}
	return scaledAmmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.