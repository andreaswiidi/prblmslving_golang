package easy

func MinimumDistance(a []int) int {

	type dataNumber struct {
		indexes      []int
		minimum      int
		validInValid bool
	}

	validInvalidNumber := map[int]*dataNumber{}
	minimum := 99999

	for index, value := range a {
		if _, found := validInvalidNumber[value]; found {
			minim := index - validInvalidNumber[value].indexes[len(validInvalidNumber[value].indexes)-1]
			validInvalidNumber[value].indexes = append(validInvalidNumber[value].indexes, index)
			validInvalidNumber[value].validInValid = true

			if validInvalidNumber[value].minimum == 0 {
				validInvalidNumber[value].minimum = minim
			} else if validInvalidNumber[value].minimum > minim {
				validInvalidNumber[value].minimum = minim
			}

			if minim < minimum {
				minimum = minim
			}

		} else {
			validInvalidNumber[value] = &dataNumber{
				indexes:      []int{index},
				minimum:      0,
				validInValid: false,
			}
		}
	}
	return minimum
}
