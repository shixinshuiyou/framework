package array

func InArray(target string, arr []string) bool {
	for _, v := range arr {
		if target == v {
			return true
		}
	}
	return false
}

func IntInArray(target int, arr []int) bool {
	for _, v := range arr {
		if target == v {
			return true
		}
	}
	return false
}

func ArrayUnique(arr []string) (dest []string) {
	unique := make(map[string]bool)
	for _, a := range arr {
		if _, ok := unique[a]; !ok {
			dest = append(dest, a)
			unique[a] = true
		}
	}
	return
}

func ArrayIntUnique(arr []int) (dest []int) {
	unique := make(map[int]bool)
	for _, a := range arr {
		if _, ok := unique[a]; !ok {
			dest = append(dest, a)
			unique[a] = true
		}
	}
	return
}
