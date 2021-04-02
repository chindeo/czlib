package main

func main() {
	gzip, err := Gzip([]byte("fdsfsdfdsf"))
	if err != nil {
		panic(err)
	}
	gunzip, err := Gunzip(gzip)
	if err != nil {
		panic(err)
	}
	println(string(gunzip))
}
