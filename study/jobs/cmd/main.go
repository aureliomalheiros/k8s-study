package main

import (
	"fmt"
	"os"
	"math/rand"
	"time"
	"strconv"
	"service-sort/pkg/err"
	
)

func main() {

	numberStr := os.Getenv("NUMBER")
	if numberStr == "" {
		fmt.Println("Error: NUMBER environemnt variable is required")
	}
	number, err := strconv.Atoi(numberStr)
	errorHandler.HandlerError(err)
	
	sort := rand.Intn(10) + 1
	
	delayStr := os.Getenv("DELAY")

	if delayStr == "" {
		fmt.Println("You don't set the variable DELAY in your environment, so we will use the default 3 seconds")
		delayStr = "3"
	}

	delay, err := strconv.Atoi(delayStr)
	errorHandler.HandlerError(err)
	
	fmt.Println("Sorting your number ...")
	
	time.Sleep(time.Duration(delay) * time.Second)

	if number == sort {
		fmt.Println("YOU WIN")
		fmt.Println("YOUR NUMBER: ", number)
		fmt.Println("SORT NUMBER: ", sort)
	}else{
		fmt.Println("YOU LOSE")
		fmt.Println("YOUR NUMBER: ", number)
		fmt.Println("SORT NUMBER: ", sort)
		os.Exit(1)
	}

}
