package main

import (
	"errors"
	"fmt"
)

func numbers()  {
	// Tipos para numeros inteiros baseado no tamanho do número
	var integer8 int8 = 1
	var integer16 int16 = 1
	var integer32 int32 = 1
	var integer64 int64 = 1

	// O tipo int utiliza um tamanho compativel com a arquitetura da maquina
	var integer int = -1

	// Semelhante ao int inclusive em suas variações (int32, int64, etc) 
	// porem não suporta valores negativos
	var unsygnedInteger uint = 1

	// Tipos para numeros quebrados baseado no tamanho do número
	var floatNum32 float32 = 1.1
	var floatNum64 float64 = 1.1
	floatNum := 1.1 


	fmt.Println(
		integer, integer8, integer16, integer32, integer64, 
		unsygnedInteger, 
		floatNum, floatNum32, floatNum64)
}

func strings()  {
	// Texto comum referenciado por ""
	var text string = "A"

	// GO não possui um tipo especifico para chars 
	// contudo se voce referenciar um texto utilizando ''
	//  ele sera convertido para o inteiro cerrespondente na tabela ASC
	char := 'A'

	fmt.Println(text, char)
}

func boolean()  {
	// Tipo para valores booleanos (verdadeiro e falso)
	trueVar := true
	falseVar := false

	fmt.Println(trueVar, falseVar)
}

func err()  {
	// GO possui um tipo especifico para erros
	// para declarar um é nessesario o uso da lib interna "errors"
	existentError := errors.New("Deu ruim")

	// O valor zero pera este tipo é nil
	var nonExistentError error

	fmt.Println(existentError, nonExistentError)
}

func main()  {
	err()
}