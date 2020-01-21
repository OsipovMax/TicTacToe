package main

func checkMap(moveCombination map[string]int) (bool, string) {
	for key, val := range moveCombination {
		if val == 3 {
			return true, key
		}
	}
	return false, ""
}

func checkWin(field [9]string) string {
	m := make(map[string]int)
	//Horizontal check
	for i := 0; i < len(field); i += 3 {
		for j := 0; j < len(field)/3; j++ {
			m[field[i+j]]++
		}
		isSuccess, playerSign := checkMap(m)
		if isSuccess {
			return playerSign
		}
		m = make(map[string]int)
	}
	//Vertical check
	for j := 0; j < len(field)/3; j++ {
		for i := 0; i < len(field); i += 3 {
			m[field[j+i]]++
		}
		isSuccess, playerSign := checkMap(m)
		if isSuccess {
			return playerSign
		}
		m = make(map[string]int)
	}
	// Diagonal check
	for i := 0; i < len(field)/3; i++ {
		m[field[i*3+i]]++
	}
	isSuccess, playerSign := checkMap(m)
	if isSuccess {
		return playerSign
	}
	m = make(map[string]int)

	for i := 0; i < len(field)/3; i++ {
		m[field[(i+1)*3-(i+1)]]++
	}
	isSuccess, playerSign = checkMap(m)
	if isSuccess {
		return playerSign
	}

	return ""
}

func checkDrawGame(field [9]string) bool {
	m := make(map[string]int)
	for _, elem := range field {
		m[elem]++
	}
	if len(m) == 2 {
		return true
	}
	return false
}

func gameOver(field [9]string) (bool, string) {
	result := checkWin(field)
	if result == "x" {
		return true, "x"
	}
	if result == "o" {
		return true, "o"
	}

	if checkDrawGame(field) == true {
		return true, "-"
	}

	return false, ""
}
