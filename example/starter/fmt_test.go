package main

import (
    "fmt";
)

type T struct {
    i int;
    s string;
    f float;
}

func main() {

    // Boolean 格式
    fmt.Printf("Boolean= True=%t, False=%t\n", true, false );   

    // 數值格式
    fmt.Printf("Integer= 10, %%d=%d, %%b=%b, %%o=%o, %%x=%x, %%X=%X\n", 10, 10, 10, 10, 10 );   

    // 顯示 int 對應的 Unicode 字元
    fmt.Printf("Character= 65, %%c=%c\n", 65 );   

    // 浮點數
    f:=0.123456789;
    fmt.Printf("float= 0.123456789, %%e=%e, %%E=%E, %%f=%f, %%g=%g, %%G=%G\n", f,f,f,f,f);   

    // 顯示型別資訊
    fmt.Printf("Type= int=%T, string=%T, float=%T\n", int(100), string("cwchiu"), float(1.34) );   

    s:="cwchiu";
    fmt.Printf("String= cwchiu, %%s=%s, %%q=%q\n", s,s);   

    t := new(T);
    t.i = 100;
    t.f = 1.34567;
    t.s = "cwchiu";                
    fmt.Printf("T= %%#v=%#v, %%v=%v, %%+v=%+v\n", t,t,t );   

}