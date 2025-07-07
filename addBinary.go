package main

import (
	_ "fmt"
)

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}
	res := []byte(b)
	carryInt := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || carryInt > 0 {
		bitA := 0
		//to create ascii 48 is 0 and 49 is 1
		if i >= 0 {
			bitA = int(a[i] - '0')
		}
		bitB := 0
		if j >= 0 {
			bitB = int(res[j] - '0')
		}
		sum := carryInt + bitA + bitB
		bit := sum % 2
		carryInt = sum / 2
		if j >= 0 {
			// to reverse to the ascii code what you actually do is you are adding bit with 48
			res[j] = byte(bit + '0')
		} else {
			res = append([]byte{'1'}, res...)
			break
		}
		i--
		j--
	}
	return string(res)
}
