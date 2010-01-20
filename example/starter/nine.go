/**
 * Description: 九九乘法表
 * 
 * author: Chui-Wen Chiu
 */
package main

import "fmt"

func main() {    
    for i := 1; i < 9; i++ {
        for j:=1; j<=9; j++{
            fmt.Printf( "%d\t", i*j);            
        }
        fmt.Printf( "\n");            
    }
   
}
