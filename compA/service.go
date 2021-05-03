package compA

import "example.com/database"

func Sum(x int, y int) int {
	return x + y
}

// Example with DB usage
func GetById(x int) CompA {
	db := database.GetDB()

	var result CompA
	db.Model(&CompA{}).First(&result)

	return result
}
