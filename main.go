package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	d := kap(6666)
	fmt.Println(d)

}

func ListToNum(c []int) (error, int64) {
	var t string
	for _, z := range c {
		num1 := strconv.Itoa(z)
		t = t + num1
	}
	num, err := strconv.ParseInt(t, 10, 64) // Convert string to integer
	if err != nil {
		fmt.Println(err)
	}
	return err, num
}

func sotrarr(d []int) ([]int, []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(d)))
	s := []int{}
	for _, z := range d {
		s = append(s, z)
	}
	sort.Ints(s)
	return d, s
}

func kap(m int) int {
	//this func take a number and returns number of steps that need to reach 6174(kaprekar number)
	a := intToSlice(m)
	r := isnotsame(a)
	if r {
		fmt.Printf("max and min of %d is same please try other number\n", m)
		return 0
	} else {
		max, min := sotrarr(a)
		_, l1 := ListToNum(max)
		_, l2 := ListToNum(min)
		i := 1
		t := l1 - l2
		v := []int{}
		for z := 0; t != 6174 && z <= 7; z++ {
			fmt.Println(t, "-----i", i, "---z", z)
			i++
			a1 := intToSlice(int(t))
			max1, min1 := sotrarr(a1)
			fmt.Println(max1, min1)
			_, l1 := ListToNum(max1)
			_, l2 := ListToNum(min1)
			v = append(v, int(l2), int(l1))
			fmt.Println(l1, "---", l2)
			mk := len(v)
			t = int64(v[mk-1] - v[mk-2])
			fmt.Println("------->t:", t)
		}
		return i
	}
}

func isnotsame(list []int) bool {
	// this func take slice and if all of didits of it's number be same return true for example [5, 5, 5, 5]
	s := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] != s {
			return false
		}
	}
	return true
}

func intToSlice(num int) []int {
	// this func take number and return slice of it's number for example : 6544 covert to [6, 5, 4, 4]
	strNum := strconv.Itoa(num)
	digits := make([]int, len(strNum))

	// Iterate over each character in the string
	for i, char := range strNum {
		// Convert the character to an integer and store it in the slice
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			//fmt.Println("Error converting character to int:", err)
			return nil
		}
		digits[i] = digit
	}

	return digits
}
