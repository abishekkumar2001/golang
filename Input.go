package main
import "fmt"

func main(){

	//Getting Integer as Input from user
	var i,j int
	fmt.Print("Enter the value of i : ")
	fmt.Scan(&i)
	fmt.Println("The value of i is ",i)
	fmt.Print("Enter the value of j : ")
	fmt.Scanln(&j)
	fmt.Println("The value of j is ",j)

	//Getting Float as Input from user
	var a,b float32
	fmt.Print("Enter the value of a and b : ")
	fmt.Scanf("%v %v",&a,&b)
	fmt.Println("The value of a is",a,"and b is",b)
	fmt.Printf("The value of a is %v and b is %v\n",a,b)

	//Getting String as Input from user
	var name1,name2 string
	fmt.Print("Enter the name of first person :")
	fmt.Scan(&name1)
	fmt.Println("Name of First person is",name1)
	fmt.Print("Enter the name of second person : ")
	fmt.Scan(&name2)
	fmt.Println("Name of Second person is",name2)

	//Getting Array as Input from user
	var n int
	fmt.Println("Enter the size of the array : ")
	fmt.Scan(&n)
	
}