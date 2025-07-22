# MConf Book Search

Sistema de busca de livros com API Go e cliente Python.

## Como executar

### API
```bash
cd apps/api
docker build -t mconf/api:candidato-1 .
docker run -ti --rm -p 3000:3000 mconf/api:candidato-1
```

### Runner
```bash
cd apps/runner
docker build -t mconf/runner:candidato-1 .
docker run -ti --rm -e API_PORT=3000 mconf/runner:candidato-1 "Lord of the Rings"
```

## Comandos disponíveis no Makefile

| Comando | Descrição |
|---------|-----------|
| `make build` | Constrói as imagens Docker da API e do runner |
| `make api` | Executa a API na porta 3000 |
| `make runner` | Executa o runner com query padrão "Lord of the Rings" |
| `make down` | Para o container da API |
| `make clean` | Para containers e remove imagens |


## Endpoints

- `GET /health` - Health check
- `GET /api/v1/search?q=<query>` - API versionada


