# Aprenda Go com testes

Repositório de estudos focado em **Go (Golang)** e **TDD (Test-Driven Development)**, acompanhando o
guia [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/) (tradução em português do _Learn Go with
Tests_).

## Estrutura do projeto

Cada diretório na raiz corresponde a um capítulo do guia (ex.: `hello/`, `integers/`, `array/`), contendo o código de
implementação e seus respectivos arquivos de teste (`*_test.go`).

## Como rodar os testes

Rodar todos os testes do projeto:

```bash
go test ./...
```

Rodar os testes de um pacote específico:

```bash
go test ./interaction/
```

Rodar benchmarks:

```bash
go test ./interaction/ -bench=.
```

Verificar a cobertura de testes:

```bash
go test -cover ./...
```
