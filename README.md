# Stress Tester

## Pré-requisitos

- Docker
- Go

---

## **Descrição Geral**

O **Stress Tester** é uma ferramenta escrita em **Go** para simular cargas de trabalho em servidores web, permitindo verificar desempenho, estabilidade e comportamento sob alta concorrência. Configurável via linha de comando, o Stress Tester oferece relatórios detalhados sobre as respostas HTTP recebidas.

---

## **Como Funciona**

### Simulação de Carga

- O Stress Tester utiliza **goroutines** para criar múltiplas requisições simultâneas.
- Suporta configuração de **número total de requisições** e **nível de concorrência**.

### Coleta de Métricas

- **Respostas HTTP** são categorizadas por código de status.
- Exibe contadores de **sucessos** e **erros**, além de contabilizar o tempo total da execução.

### Relatório Final

- Gera um relatório formatado com métricas detalhadas:
- Requisições bem-sucedidas (`200`)
- Outras respostas HTTP (`4xx`, `5xx`)
- Total de requisições realizadas
- Duração total da execução
- Exemplo:

```
      METRIC              |  VALUE
--------------------------|-----------
Successful Requests (200) |        5
Other HTTP Responses      | 429: 995
Total Requests            |     1000
Total Duration (s)        |     0.21
```

---

### Testes unitários

```bash
go test ./... -v
```

### Execução com Docker

- É necessário rodar o seguinte comando para gerar o compilado

```bash
docker build -t stress-tester .
```

- Para executar o stress tester, depois de gerado o compilado, basta executar:

```bash
docker run <sua imagem docker> —url=<http://google.com> —requests=1000 —concurrency=10
```
