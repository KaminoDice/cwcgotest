/**
 * Singleton
 *
 * @author Chui-Wen Chiu
 */
package main

import (
    "fmt"
)

type Singleton struct{
    value int
}

func (self Singleton) getValue() int{
    return self.value
}

func (self *Singleton) setValue(value int) {
    self.value = value
}

var __instance *Singleton
func GetInstance() *Singleton{
    if __instance == nil {
        __instance = new(Singleton)
        __instance.value = 1
    }
    
    return __instance;
}
     
    
func main(){
    obj := GetInstance()
    fmt.Println( obj.getValue() )
    obj.setValue( 100 )
    fmt.Println( obj.getValue() )
    obj2 := GetInstance()
    fmt.Println( obj2.getValue() )
}

