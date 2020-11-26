package main

import "fmt"

func main()  {
	salary := []int{4000,3000,1000,2000}
	fmt.Println(average(salary))
}

func average(salary []int) float64 {
	max, min, total := salary[0], salary[0], salary[0]
	for i:=1; i<len(salary); i++ {
		total += salary[i]
		if salary[i] > max {
			max = salary[i]
		}
		if salary[i] < min {
			min = salary[i]
		}
	}
	return float64(total - min - max)/float64(len(salary) - 2)
}