package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    unit := make(map[string]int)
    unit["quarter_of_a_dozen"] = 3
	unit["half_of_a_dozen"] = 6
	unit["dozen"] = 12
	unit["small_gross"] = 120
	unit["gross"] = 144
	unit["great_gross"] = 1728
	return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
    return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if units[unit] != 0 {
		_, found := bill[item]
		if  found == true {
			bill[item] = bill[item] + units[unit]
		} else {
			bill[item] = units[unit]
		}
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, found := bill[item]
	if found == true  && units[unit] != 0{
			if bill[item] - units[unit] > 0 {
				bill[item] = bill[item] - units[unit]
				return true
			} else if bill[item] - units[unit] == 0 {
				delete(bill, item)
				return true
			}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, found := bill[item]
	if  found == true {
		return bill[item], true
	}
	return 0, false
}
