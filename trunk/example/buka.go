/**
 * 布馬漫畫 buka 檔案萃取器
 *
 * @environ Go v1
 * @author Chui-Wen Chiu
 */
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {
  if len(os.Args) <= 1 {
    _, fn := filepath.Split( os.Args[0])
    fmt.Printf("Usage: %s a.buka\n", fn)
    os.Exit(-1)
  }
  f, err := os.Open( os.Args[1] )
  defer func(){ f.Close()}()
  if err != nil {
    
    fmt.Println("file open fail")    
  }
  
  dir_buka, fn_buka := filepath.Split( os.Args[1])
  
  //fmt.Println(dir_buka)
  ext:=filepath.Ext( fn_buka )
  var folder_name string
  if len(ext)>0 {
    folder_name = strings.NewReplacer( ext, "").Replace(fn_buka)
  }else{
    folder_name = fn_buka
  }
  target := filepath.Join(dir_buka, folder_name)
  os.MkdirAll( target, 0777)
  
  buf := make([]byte, 1024*1024)
  var fout *os.File
  sn := 1
  for {
    len, err := f.Read(buf)
    if err != nil {
        fmt.Println(err)
        break
    }
    
    for i:=0; i<len; i++{
        if buf[i] == 0xff { 
            if i+1 >= len {
                if fout != nil {
                    fout.Close()
                    fout = nil
                }
                break
            }
            
            if buf[i+1] == 0xd8 {
                fout, err = os.Create( filepath.Join(target, fmt.Sprintf("%03d.jpg", sn)))
                if err != nil{
                    fmt.Println(err)
                    continue
                }
                fout.Write([]byte{buf[i], buf[i+1]})
                fmt.Printf("> %03d.jpg\n", sn)
                sn++
            }else if buf[i+1] == 0xd9 {
                if fout != nil {
                    fout.Write([]byte{buf[i], buf[i+1]})
                    fout.Close()
                    fout = nil
                }
            }else {
                if fout != nil {
                    fout.Write([]byte{buf[i], buf[i+1]})
                }
            }
            i++
        }else{
            if fout != nil {
                fout.Write([]byte{buf[i]})
            }
        }
    }
    
  }
  
  
  
}