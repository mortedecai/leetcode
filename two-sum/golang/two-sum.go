package leetcode

type valToIdxMap map[int][]int

func createValToIdxMap(nums []int) (result valToIdxMap) {
	result = make(valToIdxMap)

	for i, v := range nums {
		if indices, ok := result[v]; !ok {
			result[v] = []int{i}
		} else {
			indices = append(indices, i)
			result[v] = indices
		}
	}
	return
}

func twoSum(nums []int, target int) []int {
	result := []int{-1, -1}
	mapOfValues := createValToIdxMap(nums)

	for key, indices := range mapOfValues {
		diff := target - key
		if diff == key {
			if len(indices) > 1 {
				result[0] = indices[0]
				result[1] = indices[1]
				break
			}
			continue
		}
		if diffLocations, ok := mapOfValues[diff]; ok {
			result[0] = indices[0]
			result[1] = diffLocations[0]
			break
		}
	}

	return result
}
