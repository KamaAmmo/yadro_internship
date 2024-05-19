package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var size int
	fmt.Fscan(in, &size)
	matrix_balls := make([][]int, size)
	for i := 0; i < size; i++ {
		balls := make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Fscan(in, &balls[j])
		}
		matrix_balls[i] = balls
	}

	sum_balls := make([]int, size)
	cap_container := make([]int, size)
	for i := 0; i < size; i++ {
		balls_by_color := 0
		balls_in_container := 0
		for j := 0; j < size; j++ {
			balls_by_color += matrix_balls[j][i]
			balls_in_container += matrix_balls[i][j]
		}
		sum_balls[i] = balls_by_color
		cap_container[i] = balls_in_container
	}

	sort.Ints(sum_balls)
	sort.Ints(cap_container)
	var flag = true
	for i := 0; i < size; i++ {
		if sum_balls[i] != cap_container[i] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Fprintln(out, "yes")
	} else {
		fmt.Fprintln(out, "no")
	}
}
