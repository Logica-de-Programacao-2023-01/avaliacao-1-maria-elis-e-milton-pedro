package q1

import (
	"fmt"
	"testing"
)

func TestDivideWatermelon(t *testing.T) {
	tests := []struct {
		weight      int
		expected    bool
		expectedErr bool
	}{
		{weight: -1, expected: false, expectedErr: true},
		{weight: 0, expected: false, expectedErr: true},
		{weight: 1, expected: false, expectedErr: false},
		{weight: 2, expected: false, expectedErr: false},
		{weight: 3, expected: false, expectedErr: false},
		{weight: 4, expected: true, expectedErr: false},
		{weight: 5, expected: false, expectedErr: false},
		{weight: 6, expected: true, expectedErr: false},
		{weight: 7, expected: false, expectedErr: false},
		{weight: 8, expected: true, expectedErr: false},
		{weight: 9, expected: false, expectedErr: false},
		{weight: 10, expected: true, expectedErr: false},
		{weight: 11, expected: false, expectedErr: false},
		{weight: 12, expected: true, expectedErr: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("weight=%d", test.weight), func(t *testing.T) {
			result, err := DivideWatermelon(test.weight)
			if test.expectedErr && err == nil {
				t.Errorf("Erro esperado, mas nenhum erro foi retornado")
			}

			if result != test.expected {
				t.Errorf("Resultado esperado: %t, Resultado obtido: %t", test.expected, result)
			}
		})
	}
}
 18 changes: 18 additions & 0 deletions18  
q2/q2.go
@@ -0,0 +1,18 @@
package q2

//Um dia, três melhores amigos - Pedro, Vanessa e Tônia - decidiram formar uma equipe e participar de concursos de
//programação. Os participantes geralmente recebem vários problemas durante esses concursos. Muito antes do início, os
//amigos decidiram que implementariam um problema somente se pelo menos dois deles tivessem certeza da solução. Caso
//contrário, os amigos não escreveriam a solução do problema.
//
//Este concurso oferece n problemas aos participantes. Para cada problema, sabemos qual amigo tem certeza da solução. Você
//receberá uma matriz booleana de n linhas e 3 colunas, em que a i-ésima linha representa as opiniões de Pedro, Vanessa e
//Tônia, respectivamente, sobre o i-ésimo problema. O valor "true" indica que o amigo tem certeza da solução, e "false"
//indica o contrário.
//
//Ajude os amigos a encontrar o número de problemas para os quais eles escreverão uma solução.

func ProblemsSolved(answers [][3]bool) int {
	// Seu código aqui
	return 0
}
 116 changes: 116 additions & 0 deletions116  
q2/q2_test.go
@@ -0,0 +1,116 @@
package q2

import (
	"fmt"
	"testing"
)

func TestProblemsSolved(t *testing.T) {
	tests := []struct {
		answers [][3]bool
		want    int
	}{
		{
			answers: [][3]bool{
				{true, true, true},
				{true, true, false},
				{true, false, false},
			},
			want: 2,
		},
		{
			answers: [][3]bool{
				{false, true, false},
				{false, true, true},
				{false, false, false},
			},
			want: 1,
		},
		{
			answers: [][3]bool{
				{false, true, false},
				{false, true, true},
				{false, false, false},
				{true, true, true},
				{true, true, false},
			},
			want: 3,
		},
		{
			answers: [][3]bool{
				{false, true, false},
			},
			want: 0,
		},
		{
			answers: [][3]bool{
				{true, false, false},
				{true, true, true},
			},
			want: 1,
		},
		{
			answers: [][3]bool{
				{false, false, false},
				{false, true, true},
				{true, true, true},
				{false, true, false},
				{true, false, true},
				{true, true, true},
				{false, false, true},
				{true, false, false},
				{true, true, false},
				{true, false, true},
				{false, true, false},
				{false, false, true},
				{true, true, false},
				{false, true, false},
				{true, true, false},
				{false, false, false},
				{true, true, true},
				{true, false, true},
				{false, false, true},
				{true, true, false},
				{true, true, true},
				{false, true, true},
				{true, true, false},
				{false, false, false},
				{false, false, false},
				{true, true, true},
				{false, false, false},
				{true, true, true},
				{false, true, true},
				{false, false, true},
				{false, false, false},
				{false, false, false},
				{true, true, false},
				{true, true, false},
				{true, false, true},
				{true, false, false},
				{true, false, true},
				{true, false, true},
				{false, true, true},
				{true, true, false},
				{true, true, false},
				{false, true, false},
				{true, false, true},
				{false, false, false},
				{false, false, false},
				{false, false, false},
				{false, false, true},
				{true, true, true},
				{false, true, true},
				{true, false, true},
			},
			want: 29,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test %d", i+1), func(t *testing.T) {
			if got := ProblemsSolved(test.answers); got != test.want {
				t.Errorf("ProblemsSolved() = %v, want %v", got, test.want)
			}
		})
	}
}
 18 changes: 18 additions & 0 deletions18  
q3/q3.go
@@ -0,0 +1,18 @@
package q3

//Você recebe um tabuleiro retangular de M x N quadrados. Além disso, você tem um número ilimitado de peças de dominó
//padrão de 2 x 1 quadrados. Você pode girar as peças. Você deve colocar o máximo de peças de dominó possível no
//tabuleiro, seguindo as seguintes condições:
//
//1. Cada peça de dominó cobre completamente dois quadrados.
//2. Nenhuma duas peças de dominó se sobrepõem.
//3. Cada peça de dominó está completamente dentro do tabuleiro. É permitido que a peça toque as bordas do tabuleiro.
//
//Encontre o número máximo de peças de dominó que podem ser colocadas sob essas restrições.
//
//Se M ou N forem iguais ou menores que 0, a função deve retornar um erro.

func DominoPieces(m, n int) (int, error) {
	// Seu código aqui
	return 0, nil
}
 43 changes: 43 additions & 0 deletions43  
q3/q3_test.go
@@ -0,0 +1,43 @@
package q3

import (
	"fmt"
	"testing"
)

func TestDominoPieces(t *testing.T) {
	tests := []struct {
		m       int
		n       int
		want    int
		wantErr bool
	}{
		{m: 0, n: 0, want: 0, wantErr: true},
		{m: -1, n: 1, want: 0, wantErr: true},
		{m: 1, n: -1, want: 0, wantErr: true},
		{m: 2, n: 4, want: 4, wantErr: false},
		{m: 3, n: 3, want: 4, wantErr: false},
		{m: 1, n: 5, want: 2, wantErr: false},
		{m: 1, n: 6, want: 3, wantErr: false},
		{m: 1, n: 15, want: 7, wantErr: false},
		{m: 1, n: 16, want: 8, wantErr: false},
		{m: 2, n: 5, want: 5, wantErr: false},
		{m: 15, n: 15, want: 112, wantErr: false},
		{m: 14, n: 16, want: 112, wantErr: false},
		{m: 11, n: 13, want: 71, wantErr: false},
		{m: 1, n: 1, want: 0, wantErr: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("m=%d,n=%d", test.m, test.n), func(t *testing.T) {
			got, err := DominoPieces(test.m, test.n)
			if (err != nil) != test.wantErr {
				t.Errorf("DominoPieces() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("DominoPieces() got = %v, want %v", got, test.want)
			}
		})
	}
}
 15 changes: 15 additions & 0 deletions15  
q4/q4.go
@@ -0,0 +1,15 @@
package q4

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	// Seu código aqui
	return 0, nil
}
 35 changes: 35 additions & 0 deletions35  
q4/q4_test.go
@@ -0,0 +1,35 @@
package q4

import (
	"fmt"
	"testing"
)

func TestClassifyPrices(t *testing.T) {
	tests := []struct {
		prices  []int
		want    int
		wantErr bool
	}{
		{prices: []int{1, 2, 3, 4, 5}, want: 1, wantErr: false},
		{prices: []int{5, 4, 3, 2, 1}, want: 2, wantErr: false},
		{prices: []int{1, 2, 3, 4, 3}, want: 3, wantErr: false},
		{prices: []int{}, want: 0, wantErr: true},
		{prices: []int{1}, want: 3, wantErr: false},
		{prices: []int{1, 20, 60, 90}, want: 1, wantErr: false},
		{prices: []int{10, 5, 2}, want: 2, wantErr: false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("prices=%v", test.prices), func(t *testing.T) {
			got, err := ClassifyPrices(test.prices)
			if (err != nil) != test.wantErr {
				t.Errorf("ClassifyPrices() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("ClassifyPrices() got = %v, want %v", got, test.want)
			}
		})
	}
}
 20 changes: 20 additions & 0 deletions20  
q5/q5.go
@@ -0,0 +1,20 @@
package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	// Seu código aqui
	return ""
}
 36 changes: 36 additions & 0 deletions36  
q5/q5_test.go
@@ -0,0 +1,36 @@
package q5

import (
	"testing"
)

func TestProcessString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "CEUB", want: ".c.b"},
		{input: "ceub", want: ".c.b"},
		{input: "aBAcAba", want: ".b.c.b"},
		{input: "Codeforces", want: ".c.d.f.r.c.s"},
		{input: "codeforces", want: ".c.d.f.r.c.s"},
		{input: "a", want: ""},
		{input: "A", want: ""},
		{input: "tour", want: ".t.r"},
		{input: "TOUR", want: ".t.r"},
		{input: "tourist", want: ".t.r.s.t"},
		{input: "TOURIST", want: ".t.r.s.t"},
		{input: "tourists", want: ".t.r.s.t.s"},
		{input: "AISDJasikdjnasodnhAIOOS", want: ".s.d.j.s.k.d.j.n.s.d.n.h.s"},
		{input: "AJISASJIjasiasjAIsjASJIAIAJSAJISASNMDKLEQWIOfNAKSJLDnmQOWELASKDMNQOIEWFNLAAnajksnfoaiwsLANK", want: ".j.s.s.j.j.s.s.j.s.j.s.j.j.s.j.s.s.n.m.d.k.l.q.w.f.n.k.s.j.l.d.n.m.q.w.l.s.k.d.m.n.q.w.f.n.l.n.j.k.s.n.f.w.s.l.n.k"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := ProcessString(test.input)
			if got != test.want {
				t.Errorf("ProcessString(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}
 17 changes: 17 additions & 0 deletions17  
utils/testing.go
@@ -0,0 +1,17 @@
package utils

import "math"

func AssertFloatWithPrecision(a, b, e float64) bool {
	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(b)) < e
}
