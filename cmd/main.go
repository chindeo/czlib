package main

import "github.com/chindeo/czlib"

func main() {
	gzip, err := czlib.GzipLevel2([]byte("fdsfsdfdsf"))
	if err != nil {
		panic(err)
	}
	gunzip, err := czlib.GunzipLevel2(gzip, -15)
	if err != nil {
		panic(err)
	}
	println(string(gunzip))
}
