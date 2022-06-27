package main

import "fmt"

type nomeAluno struct {
	aluno string
}

type notas struct {
	nota1 float64
	nota2 float64
}

func (nome nomeAluno) perimeter() string {
	return nome.aluno
}

func (media notas) perimeter() float64 {
	return (media.nota1 + media.nota2) / 2
}

func main() {
	var (
		notaUm       float64
		notaDois     float64
		nomeQualquer string
		qAlunos      int
		i            int
	)
	print("quantos alunos irá cadastrar? ")
	fmt.Scan(&qAlunos)
	for i = 0; i < qAlunos; i++ {
		print("digite o nome do aluno: ")
		fmt.Scan(&nomeQualquer)
		nome := nomeAluno{nomeQualquer}
		println("\ndigite as notas do aluno: ")
		fmt.Scan(&notaUm)
		fmt.Scan(&notaDois)
		media := notas{notaUm, notaDois}

		println("\nnome do aluno: ", nome.perimeter())
		fmt.Println("media do aluno: ", media.perimeter())
	}

}