package main

import "github.com/snowlyg/czlib"

func main() {
	gzip, err := czlib.Gzip([]byte("fdsfsdfdsf"))
	if err != nil {
		panic(err)
	}
	gunzip, err := czlib.Gunzip(gzip)
	if err != nil {
		panic(err)
	}
	println(string(gunzip))
}
