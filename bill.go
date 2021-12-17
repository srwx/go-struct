package main

import "fmt"

// Create struct "bill"
type bill struct {
	name  string
	items map[string]int
	tip   int
}

// Function that return bill object with initial value
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]int{},
		tip:   0,
	}
	return b
}

// Receiver function for bill object, This function used to summary total items
func (b bill) summary() string {
	total := 0
	str := "Summary:\n"

	// forEach bill.items
	for key, value := range b.items {
		str += fmt.Sprintf("- %-20s ...$%d\n", key+":", value)
		total += value
	}
	str += fmt.Sprintf("\n%-22s ...$%d", "Total:", total)

	return str
}

// Add item to bill.items
/*
   b.items คือตัวแปรชนิด map ซึ่ง pass by reference อยู่แล้ว ทำให้ตอนเขียน (b bill) ไม่ต้อง pass เป็น pointer
*/
func (b bill) updateItem(name string, value int) {
	b.items[name] = value
}

// Add tip to bill.tip
/*
   b.tip คือตัวแปรชนิด int ซึ่ง pass by value จึงต้อง pass เป็น pointer (b *bill) เพื่อให้ค่าไปอัพเดทใน object ด้วย
*/
func (b *bill) addTip(tip int) {
	b.tip = tip
}
