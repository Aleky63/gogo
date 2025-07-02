// func PetBattle(cats,dogs int){
//  switch {
//     case cats > dogs:
//         fmt.Printf("Котики победили со счетом %d:%d!", cats, dogs)
//     case dogs > cats:
//         fmt.Printf("Собачки победили со счетом %d:%d!", dogs, cats)
//     default:
//         fmt.Println("Ничья! Все дружат!")
//     }
// }
// func  printNumberInfo (num int){

//     if num < 0 {
//         fmt.Printf("Число %d отрицательное.\n", num)
//     } else if num > 0 {
//         fmt.Printf("Число %d положительное.\n", num)
//     } else {
//         fmt.Println("Число равно 0.")
//     }


//     if num%2 == 0 {
// 		fmt.Printf("Число %d четное.\n", num)
// 	} else {
// 		fmt.Printf("Число %d нечетное.\n", num)
// 	}

// if num > 0{
//     root := math.Sqrt(float64(num))
//     if root == float64 (int(root)){
//        fmt.Printf("Квадратный корень числа %d является целым числом и равен %.0f.\n", num, root) 
//     }else{
//                fmt.Printf("Квадратный корень числа %d не является целым числом и равен %.5f.\n", num, root) 
//     }
//   }
// }


func generateCompliment (name string) (string){
    number := rand.IntN(3) 
var res string
switch number {
        case 0:
           res = "Ты великолепен,"
        case 1:
           res = "У тебя потрясающая улыбка,"
        case 2:
           res = "Ты вдохновляешь,"
    }  
 return fmt.Sprintf("%s %s!",res, name)
}
// func generateCompliment (name string) (string){
//      rand.Seed(time.Now().UnixNano())

// str := []string {
//     "Ты великолепен, " + name + "!",
//         "У тебя потрясающая улыбка, " + name + "!",
//         "Ты вдохновляешь, " + name + "!",
// }
// return str[rand.Intn(len(str))]
// }