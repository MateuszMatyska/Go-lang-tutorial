package main

import (
   "fmt"
)

func main() {
   
   fmt.Println("Input Word to check if it palindrome")
   var word string
   fmt.Scan(&word)
   length := len(word)
   maxIndex := length - 1 
   var half int = (length/2) 
   
   var it int = 0
   for i:=0; i < half; i++ {	
      if word[i] == word[maxIndex - i] {
      	it++
      } else {
      	break;
      }
   }
   
   if it == half {
      fmt.Printf("%v is palindrome\n", word)   
   } else {
      fmt.Printf("%v is NOT palindrome\n", word)   
   }
}

