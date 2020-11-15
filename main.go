package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Serve struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func main() {
	//bai 2
	data := ReadFile()
	log.Println(data)
	//bai 3
	resultFilter := filter(data)
	log.Println(resultFilter)
	//bai 4
	var newItem = Serve{"fileCustome", "org.cofax.cds.FileServlet.Custome"}
	data = addItem(newItem, data)
	log.Println("Sclice sau khi them item: ", data)
	file,_:=json.MarshalIndent(data,"","")
	ioutil.WriteFile("exercises1.json",file,0664)
	//bai 5: in ra dia chi cua cac item trong slice tren
	// for i := 0; i < len(data); i++ {
	// 	log.Printf("Dia chi cua Class %d : %p\n", i, &data[i].Class)
	// 	log.Printf("Dia chi cua Name %d : %p\n", i, &data[i].Name)
	// }
	for i := range data {
		log.Println("Dia chi:  ", &data[i].Class)
	}
	//bai 6
	//Sắp xếp slice trên theo thứ tự không giảm.
	slice := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	temp := 0
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				temp = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	log.Println(slice)
	//Tạo 1 slice (copy slice) có index từ 1 đến 7.
	var copySlice []int = slice[1:8]
	log.Println("copySlice: ", copySlice)
	//Tao 1 slice (copy slice) co index tu 1 den 16. =>Loi
	// copySlice=slice[1:15]
	// log.Panicln("copySlide[1:15] ",copySlice)

	//Bai 7
	Bai7()
}
