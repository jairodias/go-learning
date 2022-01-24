package main

import "fmt"

func main() {
	var variavel1 string = "Variável explicita"
	variavel2 := "Variável implícita - inferência de tipo"

	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "var"
		variavel4 string = "var"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var", "var"
	fmt.Println(variavel5, variavel6)

	const const1 string = "const"
	fmt.Println(const1)

}
