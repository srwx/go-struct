package main

import "fmt"

func main() {
	wongBill := newBill("wong")
	wongBill.updateItem("Cake", 100)
	wongBill.addTip(50)
	fmt.Println(wongBill)
}
