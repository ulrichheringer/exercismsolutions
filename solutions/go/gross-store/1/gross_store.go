package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := units[unit]; !exists {
		return false
	}
	bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQty, itemOk := bill[item]
	if !itemOk {
		return false
	}

	// Check if the unit is valid.
	qtyToRemove, unitOk := units[unit]
	if !unitOk {
		return false
	}

	// Check if the new quantity would be negative.
	newQty := currentQty - qtyToRemove
	if newQty < 0 {
		return false
	}

	// If the new quantity is 0, remove the item completely.
	if newQty == 0 {
		delete(bill, item)
	} else {
		// Otherwise, update the quantity.
		bill[item] = newQty
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, ok := bill[item]
	return qty, ok
}
