// package main

// import (
//     "errors"
//     "fmt"
//     "strings"
// )

// func UserProfileToString(name string, age int) (string, error) {
//     trimmedName := strings.TrimSpace(name)

//     if name == "" {
//         return "", errors.New("empty name")
//     }
//     if trimmedName == "" {
//         return "", errors.New("name cannot contain only spaces")
//     }
//     if age < 0 {
//         return "", errors.New("negative age")
//     }

//     return fmt.Sprintf("Имя человека: %s, возраст: %d.", trimmedName, age), nil
// }

// func main6() {
//     str, err := UserProfileToString("Алексей", 25)
//     if err != nil {
//         fmt.Println("Ошибка:", err)
//     } else {
//         fmt.Println(str)
//     }
// }


func calculate(x1 float64, x2 float64, res string)(float64,error){
    switch res{

    case "add":
        return x1+x2,nil
    case "subtract":
        return x1-x2,nil
    case "multiply":
        return x1*x2,nil
    case "divide":

        if x2 == 0 {
            return 0, errors.New("division by zero")
        }
        return x1/x2,nil
    
default:
		return 0, errors.New("unknown operation")
  }
}


