/**
 * Bridge
 *
 */
package main
import (
    "fmt"
)


type Abstraction struct{
    implementor IImplementor
}

func (self *Abstraction) SetImplementor(implementor IImplementor) {
    self.implementor = implementor
}

func (self Abstraction) Operation(){
    self.implementor.OperatorImp()
}


type RefinedAbstraction struct{
    Abstraction
}

type IImplementor interface{
    OperatorImp()
}

type Implementor struct{
}

func (self Implementor) OperatorImp(){
    fmt.Println("op imp")
}

type ConcreteImplementorA struct{
    Implementor
}

func (self ConcreteImplementorA) OperatorImp(){
    fmt.Println("op imp #1")
}

type ConcreteImplementorB struct{
    Implementor
}

func (self ConcreteImplementorB) OperatorImp(){
    fmt.Println("op imp #2")
}

func RunDP(){
    //標準 Design Pattern 測試程式
    r := RefinedAbstraction{}

    imp := ConcreteImplementorA{}    
    r.SetImplementor(imp)
    r.Operation()
    
    imp2 := ConcreteImplementorB{}   
    r.SetImplementor(imp2)
    r.Operation()    
}
// ------- example ---------

type IHandsetSoft interface{
    Run()
}

/**
 * 手機軟體抽象類別
 *
 */
type HandsetSoft struct{
}

func (self HandsetSoft) Run(){
}

/** 
 * 手機遊戲
 *
 */
type HandsetGame struct{
    HandsetSoft
}

func (self HandsetGame) Run(){
    fmt.Println("執行手機遊戲")
}

/**
 * 手機通訊錄
 *
 */
type HandsetAddressList struct{
    HandsetSoft
}

func (self HandsetAddressList) Run(){
    fmt.Println("執行手機通訊錄")
}
/**
 * 手機品牌
 *
 */
type HandsetBrand struct{
    soft IHandsetSoft
}

func (self *HandsetBrand) SetHandsetSoft(soft IHandsetSoft){
    self.soft = soft
}

func (self HandsetBrand) Run(){

}

/**
 * 手機品牌N
 *
 */
type HandsetBrandN struct{
    HandsetBrand
}

func (self HandsetBrandN) Run(){
    fmt.Println("手機品牌N")
    self.soft.Run()
}
        
/**
 * 手機品牌M
 *
 */
type HandsetBrandM struct{
    HandsetBrand
}

func (self HandsetBrandM) Run(){
    fmt.Println("手機品牌M")
    self.soft.Run()
}
        
func main(){
    RunDP()

    
    ab := HandsetBrandN{}
    ab.SetHandsetSoft( HandsetGame{} )
    ab.Run()
    ab.SetHandsetSoft( HandsetAddressList{} )
    ab.Run()
    
    ab2 := HandsetBrandM{}
    ab2.SetHandsetSoft( HandsetGame{} )
    ab2.Run()
    ab2.SetHandsetSoft( HandsetAddressList{} )
    ab2.Run()
}