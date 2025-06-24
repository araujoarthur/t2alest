# T2Alest

Esta aplicação foi implementada na forma de REPL (Read Eval Print Loop) para simular uma árvore de diretórios.

## Partes

### main.go

O ponto de entrada do programa contém somente o inicializador do loop do REPL, que é gerenciado pelo próprio package `repl`

### Package `repl`

A package `repl` contem o loop, estrutura e registro dos comandos utilizados dentro da aplicação. Os comandos por sua vez preparam e padronizam p input para invocar os métodos da árvore

### Package `tree`

Implementa a arvore em dua partes. No arquivo `nodes.go` esta presente toda a logica referente aos nodos da arvore e seus metodos. No arquivo `tree.go` encontram-se os métodos de manejo do ADT.

