package main

import "fmt"

type node struct{
	data int
	next *node
}

//we cannot use typedef in Go... Go doesnot support typedef....

var head *node

func insertAtBeginning(val int){

	newNode := new(node) 
	//Allocating memory to new node using new function. The new function will return the pointer of the allocated new node.

	if newNode == nil{ //If any error in memory allocation new node will return nil instead of the address
		fmt.Println("Not Enough Memory")
		return
	}
	//In Go even if we use pointers for accessing values we need to use the dot operator only newNode.next ... but it is as same as newNode->next
	newNode.data = val
	if head == nil {
		newNode.next = nil //Here we need to use nil instead of NULL
		head = newNode
	} else {
		newNode.next = head
		head = newNode
	}
}

func display() {
	if head == nil {
		fmt.Println("The List is empty")
		return
	}
	fmt.Println("The contents of the Linked list are")
	var temp *node
	temp=head
	for temp != nil {
		fmt.Printf("%d ",temp.data)
		temp = temp.next
	}
	fmt.Printf("\n")
}

func insertAtPosition(pos,val int) {
	newNode := new(node)
	if newNode == nil{
		fmt.Println("Not Enough Memory")
		return
	}
	newNode.data = val
	var temp *node = head
	for i:=0;i<pos-1;i++ {
		temp = temp.next
	}
	newNode.next = temp.next
	temp.next = newNode
}

func deleteAtPos(pos int) {
	if head == nil {
		fmt.Println("The list is empty ...")
		return
	}
	if(pos == 0){
		head = head.next
		return
	}
	var temp *node = head
	for i:=0;i<pos-1;i++{
		if temp.next == nil {
			fmt.Println("Position out of bounds")
			return
		}
		temp = temp.next
	}
	temp.next = temp.next.next
}

func insertAtEnd(val int) {
	newNode := new(node)
	if newNode == nil {
		fmt.Println("Not Enough Memory")
	}
	newNode.data = val
	var temp *node = head

	for temp.next != nil{
		temp = temp.next
	}
	temp.next = newNode
	newNode.next =nil

}

func deleteAtBeginning(){
	if head == nil {
		fmt.Println("The List is empty")
		return
	}
	if head.next == nil {
		head = nil
		return
	}
	head = head.next
}

func main(){

	head =nil

	var choice,val,pos int
	for{
		fmt.Println("---------------- Linked List Menu ------------------")
		fmt.Println("1.Insertion at Beginning")
		fmt.Println("2.Dislpaying the contents of the Linked List")
		fmt.Println("3.Insertion at the specified position")
		fmt.Println("4.Deleting the node at the specified position")
		fmt.Println("5.Insertion at end")
		fmt.Println("6.Deletion at beginning")
		fmt.Println("7.Deletion at end")
		fmt.Println("8.Exit")
		fmt.Println("----------------------------------------------------")

		fmt.Printf("Enter the choice : \n")
		fmt.Scanf("%d",&choice)

		switch choice {
		case 1 : fmt.Println("Enter the data to be inserted")
				 fmt.Scanf("%d",&val)
				 insertAtBeginning(val)
		case 2 : display()
		case 3 : fmt.Println("Enter the position... Position starts from ZERO :")
				 fmt.Scanf("%d",&pos)
				 fmt.Println("Enter the data to be inserted")
				 fmt.Scanf("%d",&val)
				 if pos ==0 {
					insertAtBeginning(val)
				 } else {
					insertAtPosition(pos,val)
				 }
		case 4 : fmt.Println("Enter the position... Position starts from ZERO :")
				 fmt.Scanf("%d",&pos)
				 deleteAtPos(pos)
		case 5 : fmt.Printf("Enter the data to be inserted")
				 fmt.Scanf("%d",&val)
				 insertAtEnd(val)
		case 6 : deleteAtBeginning()
		case 7 : deleteAtEnd()

		default : fmt.Println("Invalid choice")
		}

		if choice == 8 {
			break;
		}

	}

}

func deleteAtEnd(){
	if head == nil {
		fmt.Println("The List is empty")
		return
	}
	if head.next == nil {
		head = nil
		return
	}
	var temp *node = head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next = nil
}