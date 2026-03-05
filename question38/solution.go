package question38

import (
	"strconv"
	"strings"
)

func CountAndSay(n int) string {

	var result string

	var innerCountAndSay func(n int)
	innerCountAndSay = func(currentN int) {
		if currentN == 1 {
			result = "1"
			return
		}

		var builder strings.Builder
		var count int = 1

		innerCountAndSay(currentN - 1)

		for i := 1; i <= len(result)-1; i++ {
			if result[i] == result[i-1] {
				count++
			} else {
				builder.WriteString(strconv.Itoa(count))
				builder.WriteByte(result[i-1])
				count = 1
			}
		}

		builder.WriteString(strconv.Itoa(count))
		builder.WriteByte(result[len(result)-1])

		result = builder.String()
	}

	innerCountAndSay(n)

	return result
}
