package main
//830015
//Write a program that takes a list of integers as input and returns the product of all the numbers in the list.

import "fmt" //for input and output

// My function will take list of integers as input and returns the product of the numbers
func getProductof(numbers []int) int {
    product := 1

    // this loop FOR iterates over the list and multiplies each number with the product
    for _, number := range numbers {
        product *= number
    }

    return product //just returns;)
}

func main() {
    numbers := []int{2, 5, 10}
    fmt.Println("Product of", numbers, "=", getProductof(numbers), ";")
    fmt.Println("Thank you for using my calculator on Golang")
}


