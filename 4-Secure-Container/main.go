package main

func main() {
	rangeStart := 146810
	rangeEnd := 612564
	total := 0
	for i := rangeStart; i <= rangeEnd; i++ {
		if validatePasscode(i) {
			total++
		}
	}
	println(total)
}

func columnSplit(in int) []int {
	var columns []int
	for i := in; i % 10 != 0; i /= 10 {
		columns = append(columns, i % 10)
	}

	return columns
}

func validatePasscode(in int) bool {
	columns := columnSplit(in)
	if len(columns) != 6 {
		return false
	}

	last := 99
	repeated := false
	repeatedTwice := false
	validRepeatExists := false
	for _, c := range columns {
		if c > last {
			return false
		}
		if c == last {
			if repeated {
				repeatedTwice = true
			}
			repeated = true
		} else if repeated && !repeatedTwice {
			validRepeatExists = true
			repeated = false
			repeatedTwice = false
		} else {
			repeated = false
			repeatedTwice = false
		}
		last = c
	}
	if repeated && !repeatedTwice {
		validRepeatExists = true
	}

	return validRepeatExists
}
