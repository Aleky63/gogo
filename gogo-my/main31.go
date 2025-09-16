package main

import (
	"errors"
	"fmt"
)



func intersectSlices(s1[]int, s2[]int) ([]int, error){


if s1 ==nil || s2 ==nil {
	return nil,errors.New("slices cannot be nil")
}


s3 :=[]int{}
i, j := 0, 0


    for i < len(s1) && j < len(s2)  {
	if s1[i] == s2[j] {

    s3 = append (s3, s1[i])

	
	i++
	j++

} else if s1[i] < s2[j]{
		i++
	}else {
		j++
	}
}

return s3,nil

}


func main31()  {
s1 := []int{1, 2, 3, 4, 5}
s2 := []int{3, 4, 5, 6, 7}


res, err := intersectSlices(s1,s2)

if err != nil{
	fmt.Println(err)
	return
}
fmt.Println(res)
 }