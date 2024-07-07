package lasagna

const (
	defaultPreparationTime = 2
	noodlesAmmount         = 50
	sauceAmmount           = 0.2
	defaultPortions        = 2
)


// PreparationTime calculates the time needed to prepare a lasagna
func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime <= 0 {
		preparationTime = defaultPreparationTime
	}
	return len(layers) * preparationTime
}

// Quantities calculates the ammount of noodles and sauce needed
func Quantities(layers []string) (int, float64) {
	var noodleLayers int
	var sauceLayers float64
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodleLayers += noodlesAmmount
		case "sauce":
			sauceLayers += sauceAmmount
		}
	}
	return noodleLayers, sauceLayers
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
			scaledAmmounts[i] = (ammount * float64(portions)) / defaultPortions
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