package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	cases := map[string]int{
		"quarter_of_a_dozen":	3,
		"half_of_a_dozen":		6,
		"dozen":				12,
		"small_gross":			120,
		"gross":				144,
		"great_gross":			1728,
	}
	return cases
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	if _, itemExists := bill[item]; itemExists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQuantity, itemExists := bill[item]
	if !itemExists {
		return false
	}

	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	newQuantity := currentQuantity - unitValue
	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if !exists {
		return 0, false
	}
	return quantity, true
}
