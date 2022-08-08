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

//    initData := Data{
// 	lengthRow= 5,
// 	amountRow= 5,
// 	maxNumber= 100,
//   }
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

	arr := generateArr(initData.lengthRow, initData.amountRow, initData.maxNumber)

	fmt.Println(initData)
	fmt.Println(arr)
}

func generateArr(lengthRow, amountRow, maxNumber int) []int {
	return generateRow(lengthRow, maxNumber)
}

func getRndNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max + 1)
}

func generateRndNumber1(max int) func() int {
	var myMap = make(map[int]bool)
	myMap[0] = true

	return func() int {
		rnd := getRndNumber(max)

		if myMap[rnd] == true {
			rnd = getRndNumber(max)
			fmt.Println("myMap", myMap)
		} else {
			myMap[rnd] = true
			fmt.Println("rnd", rnd)
		}
		fmt.Println("myMap", myMap)
		return rnd
	}
}

func generateRndNumber() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		fmt.Println("ааа", sum)
		return sum
	}
}

func generateRow(lengthRow, maxNumber int) []int {
	res := []int{}
	for i := 0; i <= lengthRow; i++ {
		rnd := generateRndNumber()(maxNumber)

		fmt.Println("иии", rnd)
		res = append(res, rnd)
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
