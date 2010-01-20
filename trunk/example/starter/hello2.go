/**
 * Description: 使用 os package 字串輸出
 * 
 * author: Chui-Wen Chiu
 */
package main

import "os"

func main() {
    os.Stdout.WriteString("Hello, world; or Καλημέρα κόσμε; or こんにちは 世界\n")
}
