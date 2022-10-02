package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

func jumpingOnClouds(c []int32) int32 {
	var count int32
	i := 0
	count = -1
	for i < len(c) {
		fmt.Println(i)
		if i < len(c)-2 && c[i+2] == 0 {
			i += 2
			count++
		} else {
			i++
			count++
		}
	}
	return count

}

func main() {
	c := []int32{1, 0, 0, 0, 0, 1, 0}

	result := jumpingOnClouds(c)

	fmt.Printf("%d\n", result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
