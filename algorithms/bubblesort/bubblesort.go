package bubblesort

func BubbleSort(values []int) {
	flag := true
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}

		if flag == true {
			break
		}
	}
}
