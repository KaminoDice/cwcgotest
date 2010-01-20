/**
 * Description: 迭代
 * 
 * author: Chui-Wen Chiu
 */
package main

import "fmt"

func main() {    
    for i,v := range []int{1,2,3,4,5} {
        fmt.Printf( "%d, %d\n", i,v);            
    }
   
}
