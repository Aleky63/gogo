// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func sortInner(slice []int) {
// 	sort.Slice(slice, func(i, j int) bool {

// 		if slice[i] == 0 {
// 			return true
// 		}
// 		if slice[j] == 0 {
// 			return false
// 		}

// 		isEvenI := slice[i]%2 == 0
// 		isEvenJ := slice[j]%2 == 0
// 		if isEvenI != isEvenJ {
// 			return isEvenI
// 		}

// 		return slice[i] > slice[j]
// 	})
// }

// func magicSort(data [][]int) {

// 	for _, inner := range data {
// 		sortInner(inner)
// 	}

// 	sort.SliceStable(data, func(i, j int) bool {
// 		sum := func(s []int) int {
// 			total := 0
// 			for _, v := range s {
// 				total += v
// 			}
// 			return total
// 		}
// 		return sum(data[i]) < sum(data[j])
// 	})
// }

// func main() {
// 	data := [][]int{
// 		{5, 2, 0, 3, 8, 80 ,11},
// 		{1, 7, 3,22,5,66,1025,33 },
// 		{4, 4, 4, 4},
// 		{9, 0, 2},
// 	}

// 	magicSort(data)

// 	for _, row := range data {
// 		fmt.Println(row)
// 	}
// }

package main

import (
	"fmt"

	"slices"
)

func main35() {
	nums:=[]int{1,2,3,4,5,6,7,8}
	newNums :=slices.Insert(nums,2,99999)
	fmt.Println(nums)
	fmt.Println(newNums)
	fmt.Println(slices.Contains(newNums,5))
	fmt.Println(slices.ContainsFunc(newNums,func(v int) bool {
		return v%3 == 0
	}))
fmt.Println(slices.Compare(newNums,nums))
fmt.Println(slices.CompareFunc(newNums, nums, func(v1, v2 int) int{
 if  v1%2 ==1 && v2%2 == 1 {
return 0

}
return v2 - v1

}))

fmt.Println(slices.Concat(newNums,nums))

fmt.Println(slices.Index(newNums,1))
fmt.Println(slices.IndexFunc(newNums,func(v int) bool {
		return  v%2 ==0
	}))

fmt.Println(slices.Max(newNums))
fmt.Println(slices.Delete(newNums,1,4))

fmt.Println(slices.DeleteFunc(newNums,func(v int) bool {
		return  v%2 ==0
	}))

	numbers := []int{6,5,4,3,2,1 }
    slices.Reverse(numbers)
fmt.Println(numbers )


		fmt.Println(slices.Repeat(numbers,3))
			fmt.Println(slices.Replace(numbers,3 ,5 , 777,888,100 ))

			fmt.Println(slices.Equal(numbers,nums))
			fmt.Println(slices.EqualFunc(newNums, nums, func(v1, v2 int) bool{
 
return v2 == v1
}))
	slice := []int{1,2,3,5,8,9,11,12,15,66,69 }
fmt.Println(slices.IsSorted(slice))


fmt.Println(slices.IsSortedFunc(slice, func(a,b int) int {
		return a-b
	}))


fmt.Println(slice)
	ind,found := slices.BinarySearch(slice, 66)
	fmt.Println(ind,found)
	
type User struct{
	Name string
	Age int
}

users :=[]User{
	{"max",30},
	{"maxx",33},
	{"maxxx",36},
	{"maxyy",100},
}
idx, found := slices.BinarySearchFunc(users, 36, func (v User ,target int) int{


return v.Age - target


})  

fmt.Println(idx,found)
	}


