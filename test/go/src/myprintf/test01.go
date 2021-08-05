package main

import (
    "fmt"
)

func main() {
   // Errorf  
   const name, id = "bueller", 17
   err := fmt.Errorf("user %q (id %d) not found", name, id)
   fmt.Println(err.Error())
   
   const a1 = 88

   fmt.Printf("%b", a1)
}
