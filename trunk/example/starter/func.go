/**
 * Description: 函數呼叫
 * 
 * author: Chui-Wen Chiu
 */
package main

import "fmt"
func sum(a []int) int{
    t:=0;
    for i:=0; i<len(a); i++{
        t+=a[i];
    }

    return t;
}

func main() {    
    a:= []int{1,2,3,4,5,6,7,8,9,10};   
    t:=sum( a ) ;

    fmt.Printf("%d", t);

}
