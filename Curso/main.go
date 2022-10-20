package main

import (
	"fmt"
	"os"

	"github.com/ichiban/prolog"
)

func main() {
	proyectos()
}

func proyectos() {
	p := prolog.New(os.Stdin, os.Stdout)
	if err := p.Exec(`
	%base de Conocimientos
    %hechos
    nivel(ana,tecnico).
    nivel(juan,pasante).
    nivel(luis,tuitulado).
    nivel(rosa,maestria).
    certificado_en(ana,c).
    certificado_en(rosa,php).
    certificado_en(ana,js).
    certificado_en(juan,c).
    %reglas
    certificado(X):-certificado_en(X,Y).
    programadorjr(X):-nivel(X,tecnico);nivel(X,pasante).
    programadorsr(X):-nivel(X,titulado);nivel(X,pasante),certificado(X).
    lider(X):-nivel(X,maestria).
    puede_proyecto1:-programadorjr(X),programadorsr(Y),X\=Y.
    puede_proyecto2:-programadorsr(X),programadorsr(Y),X\=Y.
    puede_proyecto3:-lider(X),programadorsr(Y),X\=Y.
    `); err != nil {
		panic(err)
	}
	sols, err := p.Query(`certificado(?).`, "rosa")
	if err != nil {
		panic(err)
	}
	defer sols.Close()

	// Iteramos sobre las soluciones a la consulta.
	for sols.Next() {
		fmt.Printf("Si.\n") //.
	}

	// Checar si hubo algun error durante la consulta.
	if err := sols.Err(); err != nil {
		panic(err)
	}
}
