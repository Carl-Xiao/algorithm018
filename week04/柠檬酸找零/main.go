package main

// 860. 柠檬水找零

// 题意一开始误解了，应该是挺简单的题目
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _, value := range bills {
		if value == 5 {
			five++
		} else if value == 10 {
			if five >= 1 {
				five--
				ten++
			} else {
				return false
			}
		} else if value == 20 {
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
