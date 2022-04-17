package gross

var unit_score = map[string]int{
	"quarter_of_a_dozen": 3,
	"half_of_a_dozen":    6,
	"dozen":              12,
	"small_gross":        120,
	"gross":              144,
	"great_gross":        1728,
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return unit_score
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	}

	bill[item] += value

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billQuantity, existsBill := bill[item]
	if !existsBill {
		return false
	}

	unitQuantity, existsUnit := units[unit]
	if !existsUnit {
		return false
	}

	finalAmount := billQuantity - unitQuantity
	if finalAmount < 0 {
		return false
	}

	if finalAmount == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = finalAmount
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
