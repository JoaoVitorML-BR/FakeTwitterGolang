package main

import (
	"fmt"
)

func main(){
	openMainMenu()
}

func openMainMenu(){
	for {
		fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
		fmt.Println("Bem Vindo(a)!")
		fmt.Println()

		fmt.Println("Digite a opção que deseja fazer: ")

		fmt.Printf("Fazer cadastro - 1 \n")
		fmt.Printf("Fazer Login - 2 \n")
		fmt.Printf("Sair - 0 \n")
		fmt.Println("Digite a opção desejada: ")

		var op int
		_,_ = fmt.Scan(&op)

		if(op == 1){
			RegisterLister()
		}

	}
}

