1º atividade:

package main

import (
	"fmt"
	"strings"
)

//Escreva um programa que receba uma string e
//converta todas as letras minúsculas para maiúsculas.
//Informe ao usuário o resultado.

func main() {
	var x string // criação da string

	fmt.Print("Digite uma palavra com letras minusculas: ")
	fmt.Scan(&x) // Escaneamento da Palavra

	o := strings.ToUpper(x) // converção de strings minusculas para maiusculas

	fmt.Printf("A palavra que você digitou %s foi convertida para maiusculo, dando: %s", x, o)
}

2º Atividade:

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string
// e remova todos os espaços em branco.
// Informe ao usuário o resultado.

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma palavra: ")
	scanner.Scan()
	p := scanner.Text()

	o := strings.ReplaceAll(p, " ", "")

	fmt.Printf("A palavra que você digitou %s foi retirada os espaços, dando: %s", p, o)
}

3º // 9º Atividade:

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Escreva um programa que receba uma string e substitua todas as
//ocorrências desse caractere na string por outro caractere especificado
//pelo usuário. Informe ao usuário o resultado.

func main() {

	var i, j string // Criação das variaveis em string

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma palavra: ")
	scanner.Scan()
	p := scanner.Text()

	fmt.Print("Digite a letra que queira substituir: ")
	fmt.Scan(&i)

	fmt.Print("Digite a letra para substituir a letra que você queira substituir: ")
	fmt.Scan(&j)
	o := strings.ReplaceAll(p, i, j)

	fmt.Printf("A palavra que você digitou %s foi retirada os espaços, dando: %s", p, o)
}

4º Atividade:

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Solicite ao usuário duas strings
//e informe se elas são iguais ou diferentes.

func main() {

	scanner := bufio.NewScanner(os.Stdin) // Escaneamento da string

	x := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma palavra: ")
	scanner.Scan()
	p := scanner.Text()

	fmt.Print("Digite outra palavra: ")
	x.Scan()
	v := x.Text()

	if p == v {
		fmt.Printf("As palavras %s e %s são iguais", p, v)
	} else {
		fmt.Printf("As palavras %s e %s são diferentes", p, v)
	}
}

5º atividade: (não terminada)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Solicite ao usuário duas strings
//e informe se elas são iguais ou diferentes.

func main() {
	var x string
	scanner := bufio.NewScanner(os.Stdin) // Escaneamento da string

	fmt.Print("Digite uma número: ")
	fmt.Scan(&x)

	scanner.Scan()
	p := scanner.Text()

	z, err := strconv.ParseFloat(x, 64)

	if err != nil {
		fmt.Printf("O número %0.2f é um numero em float", z)
	} else {
		fmt.Printf("O numero %0.2f não é um número em float", z)
	}

}

