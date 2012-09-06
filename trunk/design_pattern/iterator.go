/*
Composite Patter
*/
package main
import (
    "fmt"
)

type Aggregate interface{
    CreateIterator() *ConcreteIterator
}

type ConcreteAggregate struct{
    items *[]string
}

func NewConcreteAggregate() *ConcreteAggregate{
    ptr := new(ConcreteAggregate)
    ptr.items = new([]string)
    
    return ptr
}

func (self *ConcreteAggregate) Add(item string){
    *self.items = append(*self.items, item)
}
    
    
func (self ConcreteAggregate) CreateIterator() *ConcreteIterator{
    return NewConcreteIterator(self);
}

func (self ConcreteAggregate) Count() int{
    return len(*self.items)
}
    
func (self ConcreteAggregate) GetItem(index int) string{
    return (*self.items)[index]
}
   
    
type Iterator interface{
    First() string
    Next() (string, error)
    IsDone() bool
    CurrentItem() string
}

type OverflowError struct{
}

func (self OverflowError) Error() string{
    return "Over"
}

type ConcreteIterator struct{
    aggregate ConcreteAggregate
    current int
}

func NewConcreteIterator(aggregate ConcreteAggregate)*ConcreteIterator{
    ptr := new(ConcreteIterator)
    ptr.aggregate = aggregate
    ptr.current = 0
    
    return ptr
}

func (self ConcreteIterator) First() string{
    return self.aggregate.GetItem(0)
}

func (self *ConcreteIterator) Next() (string, error){
    self.current += 1
    if self.IsDone(){
        return "", OverflowError{}
    }
    return self.CurrentItem(), nil
}

func (self ConcreteIterator) IsDone() bool{        
    return self.current >= self.aggregate.Count()
}

func (self ConcreteIterator)CurrentItem()string{
    return self.aggregate.GetItem(self.current)
}
    

func main(){
    a := NewConcreteAggregate()
    a.Add( "大鳥" )    
    a.Add( "小菜" )
    a.Add( "行李" )
    a.Add( "老外" )
    a.Add( "巴士內部員工" )
    a.Add( "小偷"  )

    it := a.CreateIterator()
    
    /*item :=*/ it.First()
    
    for  {
        fmt.Printf("%s 請買車票!\n", it.CurrentItem())
        _, err := it.Next()
        if err != nil {
            break
        }
    }
    
    /*for !it.IsDone(){
        fmt.Printf("%s 請買車票!\n", it.CurrentItem())
        it.Next()
    }*/
}   
