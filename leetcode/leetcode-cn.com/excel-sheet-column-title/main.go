package main

import ("fmt")

func main(){

	fmt.Println(convertToTitle(52))
}

func convertToTitle(columnNumber int) string {
	var str []byte
	
	for columnNumber > 0 {
		columnNumber --
		
		yu := columnNumber % 26
		t := []byte(fmt.Sprintf("%s", string(yu + 65)))
		
		str = append(t, str...)
		
		columnNumber = columnNumber / 26
	}

	return string(str)
}
