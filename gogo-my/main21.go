package main

import (	
	"fmt"	
)


func isPalindrome (nums [10]int){
isPal := true

 for i := range len(nums)/2 {

	if (nums[i] != nums[9-i]) {
		isPal = false
		break
	}
}
if isPal {
        fmt.Println("Это палиндром!")
    } else {
        fmt.Println("Не палиндром!")
    }
}


func main21() {
    
    nums := [10]int{}
    isPalindrome(nums) 
    
}