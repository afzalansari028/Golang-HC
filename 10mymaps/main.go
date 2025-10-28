package main

import "fmt"

func main() {
	fmt.Println("Welcome to the learning of Map in GO")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages is ", languages)

	fmt.Println("JS is shorts for :", languages["JS"])
	//	fmt.Println("PY is shorts for :", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of languages is ", languages)

	//loops are interesting in golang

	for key, value := range languages { //traversing the map
		fmt.Printf("For key %v and for value %v\n", key, value)
	}

	arr := []int{1, 4, 7, 5, 5, 1, 4, 2, 9, 3}
	RemoveDuplicate(arr)
	FrequencyOfElemnts(arr)
}

func FrequencyOfElemnts(arr []int) {
	result := make(map[int]int)

	for _, val := range arr {
		result[val]++
	}
	fmt.Println("result---", result)

}

func RemoveDuplicate(arr []int) {

	visited := make(map[int]bool)
	var result []int

	for _, val := range arr {

		if !visited[val] {
			visited[val] = true

			result = append(result, val)
		}
	}
	fmt.Println("remove duplicates:::", result)
}
