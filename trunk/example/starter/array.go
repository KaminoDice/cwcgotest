/**
 * Description: 陣列
 * 
 * author: Chui-Wen Chiu
 */
package main

import "fmt"

func main() {    
    a:= []int{1,2,3,4,5,6,7,8,9,10};   
    for i:=0; i<len(a); i++{
        fmt.Printf("%d", a[i]);
    }
}
