package main

import (
	"fmt"
	"github.com/miaoyin/go-command/http"
)

func main(){
	if err := http.Cmd.Execute();err!=nil{
		fmt.Println(err)
	}
}