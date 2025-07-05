package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Massiv so'zlarini kiriting:")
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)

	if len(words) == 0 {
		fmt.Println("Hech qanday so'z kiritilmadi")
		return
	}

	result := longestCommonPrefix(words)

	if result == "" {
		fmt.Println("Kiritilgan so'zlar ichida umumiy prefix mavjud emas.")
	} else {
		fmt.Println("Eng uzun umumiy prefix:", result)
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {
		for len(prefix) > 0 && (len(str) < len(prefix) || str[:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}

	return prefix
}
