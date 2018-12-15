package main

import "fmt"

func CompressString(str string) string {
	length := len(str)
	strByte := []byte(str)
	strByteCom := make([]byte, length+1)

	if length <= 2 {
		return str
	}

	idx, freq := 0, 1
	for i := 1; i < length; i++ {
		if idx >= length {
			return str
		} else if strByte[i] != strByte[i-1] {
			strByteCom[idx] = strByte[i-1]
			strByteCom[idx+1] = byte(freq+'0')
			idx += 2
			freq = 1
		} else {
			freq++
			if i == length-1 {
				strByteCom[idx] = strByte[i]
				strByteCom[idx+1] = byte(freq+'0')
				idx += 2
			}
		}
	}

	if idx >= length {
		return str
	} else {
		return string(strByteCom)
	}
}

func main() {
	str1 := "appplllee"
	fmt.Printf("compressString(%v) = %v\n", str1, CompressString(str1))
	str2 := "apple"
	fmt.Printf("compressString(%v) = %v\n", str2, CompressString(str2))
}
