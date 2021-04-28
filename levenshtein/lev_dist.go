package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
)

var (
	matrix [][]int
)

func worker(id int, workerNum int, row int, col int, index int, bufferSize int, s []byte, t []byte, wg *sync.WaitGroup, w *[]sync.WaitGroup) {

	var r, c int
	totalTimes := row + col + 1

	for index < totalTimes {

		// calculate this thread's cycletimes in this round
		var cycleTimes, remain int = 0, bufferSize
		for remain > workerNum {
			remain = remain - workerNum
			cycleTimes++
		}
		if id <= remain {
			cycleTimes++
		}
		// fmt.Println("I am No.", id, "thread, now pos, buffersize and cycleTimes are:", index, bufferSize, cycleTimes)

		for time := 0; time < cycleTimes; time++ {
			// calculate upper half	part of matrix
			if index < col+1 {
				r = id + workerNum*time
				c = index - r

				if s[r-1] == t[c-1] {
					matrix[r][c] = matrix[r-1][c-1]
				} else {
					matrix[r][c] = minimum(matrix[r-1][c]+1, matrix[r][c-1]+1, matrix[r-1][c-1]+1)
				}
			}
			// calculate other part of matrix
			if index >= col+1 {
				c = col - id - workerNum*time + 1
				r = index - c

				if s[r-1] == t[c-1] {
					matrix[r][c] = matrix[r-1][c-1]
				} else {
					matrix[r][c] = minimum(matrix[r-1][c]+1, matrix[r][c-1]+1, matrix[r-1][c-1]+1)
				}
			}
		}

		// the job of this round is done and wait for other partners
		(*w)[index].Done()
		(*w)[index].Wait()

		// update buffersize and index for next round
		index++
		if index <= row {
			bufferSize++
		} else if (index <= col+1) && (index > row) {
			bufferSize = row
		} else {
			bufferSize--
		}
	}
	wg.Done()
}

func minimum(num1 int, num2 int, num3 int) int {

	var tmp int
	if num1 <= num2 {
		tmp = num1
	} else {
		tmp = num2
	}

	if tmp <= num3 {
		return tmp
	}
	return num3
}

func main() {

	var sourceStr string
	var targetStr string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	targetStr = scanner.Text()
	scanner.Scan()
	sourceStr = scanner.Text()

	// source是较短字符串，target是较长字符串
	row, column := len(sourceStr), len(targetStr)

	// get number of CPUS
	var cpus int = runtime.NumCPU()
	if row < cpus {
		cpus = row
	}
	runtime.GOMAXPROCS(cpus)

	// create and init the two-dimensional matrix
	matrix = make([][]int, row+1)
	for i := range matrix {
		matrix[i] = make([]int, column+1)
	}
	for i := range matrix {
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}

	// init the position and number of items
	var calculateItems int = 1
	var pos int = 2

	// worker goroutine : sync mutex
	var wg sync.WaitGroup
	var w []sync.WaitGroup
	for i := 0; i < (row + column + 1); i++ {
		var tmpWg sync.WaitGroup
		tmpWg.Add(cpus)
		w = append(w, tmpWg)
	}

	wg.Add(cpus)
	for i := 1; i <= cpus; i++ {
		go worker(i, cpus, row, column, pos, calculateItems, []byte(sourceStr), []byte(targetStr), &wg, &w)
	}
	wg.Wait()
	fmt.Println(matrix[row][column])
}
