package main

import "fmt"

func main() {
	var A []int
	A = []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(A))

	A = []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII2(A))

	A = []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII3(A))

	A = []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII4(A))
}

/**
选择排序
*/
func sortArrayByParityII(A []int) []int {
	for i := 0; i < len(A); i++ {
		if i&1 == A[i]&1 {
			continue
		}
		index := i
		for j := index + 1; j < len(A); j++ {
			if i&1 == A[j]&1 {
				A[i], A[j] = A[j], A[i]
				break
			}
		}
	}
	return A
}

/**
两次遍历
*/
func sortArrayByParityII2(A []int) []int {
	A1, A2 := []int{}, []int{}
	for i := 0; i < len(A); i++ {
		if A[i]&1 == 0 { //偶数
			A2 = append(A2, A[i])
		} else {
			A1 = append(A1, A[i])
		}
	}
	for i := 0; i < len(A); i++ {
		if i&1 == 0 { //偶数
			A[i] = A2[0]
			A2 = A2[1:]
		} else {
			A[i] = A1[0]
			A1 = A1[1:]
		}
	}
	return A
}

/**
双指针
	i 偶数指针，j 奇数指针
*/
func sortArrayByParityII3(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		if A[i]&1 != 0 { //偶数位出现了奇数
			for A[j]&1 == 1 {
				j += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}

/**
他人提交中最快的做法
*/
func sortArrayByParityII4(A []int) []int {
    l,r:=0,1
    for{        
        for l<len(A) && A[l]%2==0{
            l+=2
        }
        for r<len(A) && A[r]%2==1{
            r+=2
        }
        if l>len(A)||r>len(A){
            break
        }
        A[l],A[r]=A[r],A[l]
    }
    return A
}