/**
 * Abstrract Factory Pattern
 *
 * @author Chui-Wen chiu
 */
package main

import (
    "fmt"
)

type IFactory interface{
    CreateUser() IUser
    CreateDepartment() IDepartment
}

var SqlserverFactoryInstance *SqlserverFactory

func GetSqlserverFactoryInstance() *SqlserverFactory{
    if SqlserverFactoryInstance == nil {
        SqlserverFactoryInstance = new(SqlserverFactory)
    }
    
    return SqlserverFactoryInstance
}

type SqlserverFactory struct{    
}

func (self SqlserverFactory) CreateUser() IUser{
    return new(SqlserverUser)
}

func (self SqlserverFactory) CreateDepartment() IDepartment{
    return new(SqlserverDepartment)
}

var AccessFactoryInstace *AccessFactory
func getAccessFactoryInstance() *AccessFactory{
    if AccessFactoryInstace == nil {
        AccessFactoryInstace = new(AccessFactory)
    }
    
    return AccessFactoryInstace
}


type AccessFactory struct{
}

func (self AccessFactory) CreateUser() IUser{
    return new(AccessUser)
}

func (self AccessFactory) CreateDepartment() IDepartment{
    return new(AccessDepartment)
}

type IUser interface{
    Insert(item User)
    GetUser(idx int)
}

type SqlserverUser struct{
}

func (self SqlserverUser) Insert(item User) {
    fmt.Printf( "SqlServer User 表新增紀錄 %s\n", item.getName())
}

func (self SqlserverUser) GetUser(idx int) {
    fmt.Println( "SqlServer User 表回傳一筆紀錄")
}

    
type AccessUser struct{
}

func (self AccessUser) Insert(item User) {
    fmt.Printf( "Access User 表新增紀錄 %s\n", item.getName())
}

func (self AccessUser) GetUser(idx int) {
    fmt.Println( "Access User 表回傳一筆紀錄")
}
        
        
type IDepartment interface{
    Insert(item Department)    
    GetDepartment(idx int)
}

type SqlserverDepartment struct{
}

func (self SqlserverDepartment) Insert(item Department){
        fmt.Printf( "SqlServer Department 表新增紀錄 %s\n", item.getName())
}
    
func (self SqlserverDepartment) GetDepartment(idx int){
        fmt.Println( "SqlServer Department 表回傳一筆紀錄")
}

type AccessDepartment struct{
}

func (self AccessDepartment) Insert(item Department){
        fmt.Printf( "Access Department 表新增紀錄 %s\n", item.getName())
}
    
func (self AccessDepartment) GetDepartment(idx int){
        fmt.Println( "Access Department 表回傳一筆紀錄")
}

type User struct{
    name string
}

func (self User) getName() string{
    return self.name
}

func NewUser() *User{
    user := new(User)
    user.name = "大雄"
    return user
}


type Department struct{
    name string
}

func (self Department) getName() string{
    return self.name
}

func NewDepartment() *Department{
    dep := new(Department)
    dep.name = "客服"
    return dep
}
        
func main(){
    user := NewUser()
    dept := NewDepartment()
    
    iu := GetSqlserverFactoryInstance().CreateUser()
    iu.Insert(*user)
    iu.GetUser(1)

    idept := GetSqlserverFactoryInstance().CreateDepartment()
    idept.Insert(*dept)
    idept.GetDepartment(1)
    
    iu2 := getAccessFactoryInstance().CreateUser()
    iu2.Insert(*user)
    iu2.GetUser(1)

    idept2 := getAccessFactoryInstance().CreateDepartment()
    idept2.Insert(*dept)
    idept2.GetDepartment(1)
}