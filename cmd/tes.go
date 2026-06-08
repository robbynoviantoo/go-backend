//go:build ignore

package main

import ("fmt")

func coba(x string) string{
	return "coba" + x
}

func main(){
	fmt.Println(coba("oke"))
}
