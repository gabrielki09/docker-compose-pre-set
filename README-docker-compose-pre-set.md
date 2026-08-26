# Docker Compose Pre Set

CLI simples em Go para gerar um arquivo `compose.yaml` com um serviço PostgreSQL pré-configurado.

A ideia é acelerar a criação de um `docker compose` básico para projetos locais, evitando escrever manualmente a mesma estrutura de PostgreSQL toda vez.

---

## Índice

- [Visão geral](#visão-geral)
- [Estrutura do projeto](#estrutura-do-projeto)
- [Pré-requisitos](#pré-requisitos)
- [Instalação das dependências](#instalação-das-dependências)
- [Build local](#build-local)
- [Instalar como comando global](#instalar-como-comando-global)
- [Uso](#uso)
- [Flags disponíveis](#flags-disponíveis)
- [Exemplo de uso](#exemplo-de-uso)
- [Exemplo de compose gerado](#exemplo-de-compose-gerado)
- [Validação do arquivo gerado](#validação-do-arquivo-gerado)
- [Comportamento atual](#comportamento-atual)
- [Problemas conhecidos](#problemas-conhecidos)
- [Próximos ajustes sugeridos](#próximos-ajustes-sugeridos)

---

## Visão geral

O comando gera um arquivo:

```text
compose.yaml
```

no diretório em que ele for executado.

O compose gerado contém:

```text
services
volumes
PostgreSQL
container_name
environment
ports
volume persistente
```

Serviço utilizado:

```yaml
image: postgres:18-alpine
```

---

## Estrutura do projeto

```text
docker-compose-pre-set/
├── cmd/
│   └── cli/
│       └── main.go
├── pkg/
│   ├── config.go
│   ├── constants.go
│   ├── content.go
│   ├── file.go
│   ├── runner.go
│   └── str_helper.go
├── go.mod
└── go.sum
```

### `cmd/cli/main.go`

Arquivo de entrada da CLI.

Responsável por:

```text
configurar o comando Cobra
declarar as flags
chamar o Runner
ativar logs de debug
```

### `pkg/runner.go`

Orquestra o fluxo principal:

```text
monta a configuração do compose
cria o arquivo compose.yaml
retorna erro quando algo falha
```

### `pkg/config.go`

Contém:

```text
struct DockerFile
validação dos dados
montagem dos nomes finais
porta do banco
volume
container
database
```

### `pkg/content.go`

Gera o conteúdo do arquivo `compose.yaml` usando `text/template`.

### `pkg/file.go`

Cria fisicamente o arquivo `compose.yaml` no diretório atual.

### `pkg/constants.go`

Centraliza:

```text
nome do arquivo
valores padrão
limites de tamanho
mensagens de erro
```

### `pkg/str_helper.go`

Contém helpers de string usados na normalização dos valores.

---

## Pré-requisitos

Necessário ter instalado:

```text
Go
Docker
Docker Compose
```

Para validar:

```bash
go version
docker --version
docker compose version
```

---

## Instalação das dependências

Na raiz do projeto:

```bash
go mod tidy
```

---

## Build local

Como o `main.go` fica em `cmd/cli`, a build deve apontar para esse pacote.

### Linux, WSL ou macOS

```bash
go build -o docker-compose-pre-set ./cmd/cli
```

### Windows PowerShell

```powershell
go build -o docker-compose-pre-set.exe ./cmd/cli
```

---

## Instalar como comando global

A melhor forma é colocar o binário em uma pasta que esteja no `PATH`.

Não é necessário colocar o executável na raiz do sistema.

### Linux / WSL / macOS

Criar pasta local para binários:

```bash
mkdir -p ~/.local/bin
```

Mover o binário:

```bash
mv docker-compose-pre-set ~/.local/bin/
chmod +x ~/.local/bin/docker-compose-pre-set
```

Adicionar ao `PATH`, caso ainda não esteja.

Para Bash:

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc
```

Para Zsh:

```bash
echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

Validar:

```bash
docker-compose-pre-set --help
```

### Windows PowerShell

Criar uma pasta para binários do usuário:

```powershell
mkdir "$env:USERPROFILE\bin"
```

Mover o executável:

```powershell
Move-Item .\docker-compose-pre-set.exe "$env:USERPROFILE\bin\docker-compose-pre-set.exe"
```

Adicionar ao `PATH` do usuário:

```powershell
setx PATH "$env:PATH;$env:USERPROFILE\bin"
```

Fechar e abrir o terminal novamente.

Validar:

```powershell
docker-compose-pre-set --help
```

---

## Uso

Na versão atual, o comando principal usa o subcomando:

```text
docker
```

Portanto, o uso fica assim:

```bash
docker-compose-pre-set docker [flags]
```

Exemplo:

```bash
docker-compose-pre-set docker \
  --name postgres \
  --container minha_api \
  --database minha_api \
  --user minha_api \
  --password secret \
  --port 5432 \
  --volume minha_api
```

---

## Flags disponíveis

| Flag | Atalho | Descrição |
|---|---:|---|
| `--name` | `-n` | Nome do serviço dentro do `compose.yaml`. |
| `--container` | `-c` | Nome base do container. |
| `--database` | `-b` | Nome base do banco de dados. |
| `--user` | `-u` | Usuário do banco de dados. |
| `--password` | `-w` | Senha do banco de dados. |
| `--port` | `-p` | Porta externa do PostgreSQL. |
| `--volume` | `-m` | Nome base do volume. |
| `--debug` | `-d` | Ativa logs em modo debug. |

---

## Exemplo de uso

Comando:

```bash
docker-compose-pre-set docker \
  --name postgres \
  --container minha_api \
  --database minha_api \
  --user minha_api \
  --password secret \
  --port 5432 \
  --volume minha_api
```

Arquivo gerado:

```text
compose.yaml
```

---

## Exemplo de compose gerado

```yaml
services:
    postgres:
        image: postgres:18-alpine
        container_name: minha_api_postgres
        restart: always
        environment:
            POSTGRES_DB: minha_api_db
            POSTGRES_USER: minha_api
            POSTGRES_PASSWORD: secret
        ports:
            - "5432:5432"
        volumes:
            - minha_api_postgres_data:/var/lib/postgresql/data

volumes:
    minha_api_postgres_data:
```

---

## Validação do arquivo gerado

Depois de gerar o `compose.yaml`, validar com:

```bash
docker compose config
```

Subir o banco:

```bash
docker compose up -d
```

Ver containers:

```bash
docker ps
```

Parar:

```bash
docker compose down
```

Parar e remover volume:

```bash
docker compose down -v
```

---

## Comportamento atual

### Arquivo existente

O app usa `os.O_EXCL`, então não sobrescreve `compose.yaml` existente.

Se o arquivo já existir, retorna erro.

```text
O arquivo compose.yaml já existe.
```

### Porta

A porta informada é usada no formato:

```text
porta_externa:5432
```

Exemplo:

```text
5433:5432
```

### Nomes derivados

Alguns campos recebem sufixos automaticamente.

Exemplo:

```text
container informado: minha_api
container final:     minha_api_postgres

database informado: minha_api
database final:     minha_api_db

volume informado: minha_api
volume final:     minha_api_postgres_data
```

---

## ## Build recomendada para uso pessoal

Fluxo simples:

```bash
go mod tidy
go fmt ./...
go build -o docker-compose-pre-set ./cmd/cli
```

Instalar localmente:

```bash
mkdir -p ~/.local/bin
mv docker-compose-pre-set ~/.local/bin/
chmod +x ~/.local/bin/docker-compose-pre-set
```

Testar:

```bash
docker-compose-pre-set --help
```

Gerar compose:

```bash
docker-compose-pre-set docker \
  --name postgres \
  --container minha_api \
  --database minha_api \
  --user minha_api \
  --password secret \
  --port 5432 \
  --volume minha_api
```

Validar:

```bash
docker compose config
```

Subir:

```bash
docker compose up -d
```

---

## Observação

Essa CLI é útil para uso pessoal e para acelerar setup local de projetos.

