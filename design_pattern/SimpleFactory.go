/**
 * Simple Factory
 *
 * @author Chui-Wen chiu
 */
package main

import (
    "fmt"
    //"errors"
)

type IOperation interface{
    GetResult() (int, error)
    SetA(a int)
    SetB(b int)
}

type Operation struct{
    NumberA int
    NumberB int
}

func (self Operation) GetResult() (int, error){
    return -1, nil
}

func (self *Operation) SetA(a int) {
    self.NumberA = a
}

func (self *Operation) SetB(b int) {
    self.NumberB = b
}

type OperationAdd struct{
    Operation
}

func (self OperationAdd) GetResult() (int, error){
    return self.NumberA + self.NumberB, nil
}

func NewOperationAdd() *OperationAdd{
    op := OperationAdd{}
    return &op
}

type OperationSub struct{
    Operation
}

func (self OperationSub) GetResult() (int, error){
    return self.NumberA - self.NumberB, nil
}

func NewOperationSub() *OperationSub{
    op := OperationSub{}
    return &op
}

type OperationMul struct{
    Operation
}

func (self OperationMul) GetResult() (int, error){
    return self.NumberA * self.NumberB, nil
}

func NewOperationMul() *OperationMul{
    op := OperationMul{}
    return &op
}

type OperationError struct{
    Err  error
}

func (self OperationError) Error() string{
    return "op error"
}

type DivZeroError struct{
    Err  error
}

func (self DivZeroError) Error() string{
    return "div zero"
}

type OperationDiv struct{
    Operation
}

func (self OperationDiv) GetResult() (int, error) {
    if self.NumberB == 0 {
        return -1, DivZeroError{}
    }

    return self.NumberA / self.NumberB, nil
}

func NewOperationDiv() *OperationDiv{
    op := OperationDiv{}
    return &op
}

func CreateOperation(operate string) (IOperation, error){
    switch(operate){
    case "+":
        return NewOperationAdd(), nil
    case "-":
        return NewOperationSub(), nil
    case "*":
        return NewOperationMul(), nil
    case "/":
        return NewOperationDiv(), nil
    }
    return nil, OperationError{}
}

func doOp(opc string, a int, b int){
    op, err1 := CreateOperation(opc)
    if err1 != nil {
        fmt.Printf("%s => %s\n", opc, err1)
        return
    }    
    
    op.SetA(a)
    op.SetB(b)
    
    result, err2 := op.GetResult()
    if err2 == nil {
        fmt.Printf("%d %s %d = %d\n", a, opc, b, result)
    }else {
        fmt.Printf("%d %s %d = %s\n", a, opc, b, err2)
    }
}

func main(){
    a, b := 1, 2
    doOp("+", a, b)
    doOp("-", a, b)
    doOp("*", a, b)
    doOp("/", a, b)
    doOp("%", a, b)
    
    doOp("/", a, 0)

}
