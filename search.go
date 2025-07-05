package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Massivni kiriting:")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	strNums := strings.Fields(line)

	var nums []int
	for _, s := range strNums {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println("Xato: noto‘g‘ri son kiritildi")
			return
		}
		nums = append(nums, n)
	}

	fmt.Print("Target sonni kiriting: ")
	targetLine, _ := reader.ReadString('\n')
	targetLine = strings.TrimSpace(targetLine)
	target, err := strconv.Atoi(targetLine)
	if err != nil {
		fmt.Println("Xato: target butun son bo‘lishi kerak!")
		return
	}

	index := searchB(nums, target)
	fmt.Println("Natija:", index)
}

func searchB(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
