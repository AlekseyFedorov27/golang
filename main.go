package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	lengthRow int
	amountRow int
	maxNumber int
}

var myMap = make(map[int]bool)

func main() {
	var initData Data
	initData.lengthRow = 5
	initData.amountRow = 5
	initData.maxNumber = 100

	err := ValidateData(initData.lengthRow, initData.amountRow, initData.maxNumber)
	if err != false {
		fmt.Println(err)
		return
	}

	generate(initData.lengthRow, initData.amountRow, initData.maxNumber)
}

func generate(lengthRow, amountRow, maxNumber int) {
	shuffleArr := generateShuffleArr(maxNumber)
	res := sliceArr(shuffleArr, lengthRow, amountRow, maxNumber)
	printResult(res)
}

func printResult(arr [][]int) {
	for _, row := range arr {
		fmt.Println(row)
	}
}

func generateShuffleArr(maxNumber int) []int {
	a := make([]int, maxNumber)

	for i := 0; i < maxNumber; i++ {
		a[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func sliceArr(arr []int, lengthRow, amountRow, maxNumber int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < amountRow; i++ {
		row := arr[i*lengthRow : i*lengthRow+lengthRow]
		res = append(res, row)
	}
	return res
}

func ValidateData(lengthRow, amountRow, maxNumber int) bool {
	res := true
	switch {
	case lengthRow <= 0 || amountRow <= 0 || maxNumber <= 0:
		fmt.Println("Одно из значений меньше ноля")
		res = true
		break
	case lengthRow*amountRow > maxNumber:
		fmt.Println("максимальное случайное значение меньше допустимого")
		res = true
		break
	default:
		res = false
		break
	}
	return res
}
