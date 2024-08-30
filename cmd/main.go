package main

import "fmt"

func main() {
	//fmt.Println("hi hi ha ha")
	//fmt.Println("len : ", len("https://developers.google.com/speed/webp/gallery?hl=vi"))
//	fmt.Println(`{\"t\": \"https://developers.google.com/speed/webp/gallery?hl=vi ssdsdsd\", \"lk\": [{\"e\": 54, \"s\": 0}]}`)
	url := "https://developers.google.com/speed/webp/gallery?hl=vi ssdsdsd"
	e := 54
	s := 0

	// Format the JSON string
	jsonStr := fmt.Sprintf(`{\"t\": \"%s\", \"lk\": [{\"e\": %d, \"s\": %d}]}`, url, e, s)
	
	// Print the JSON stringc    cc
	fmt.Println(jsonStr)csa  dscdvsdvcd
	//kk 55 uuu 44  dsvd 999 dd fff  ddd   ddd
	//40
}
