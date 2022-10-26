// Package census simulates a system used to collect census data.
package census

import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	newRes := Resident{name, age, address}
	return &newRes
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	if r.Name != "" && r.Age >= 0 && r.Address != nil && r.Address["street"] != "" {
		return true
	}
	return false
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	i := 0
	for j, _ := range residents {
		if residents[i].HasRequiredInfo() {
			fmt.Print(residents[j])
			i = i + 1
			fmt.Println(i)
		}
	}
	return i
}
