# CRUD Go Application

Esta é uma aplicação CRUD simples escrita em Go, utilizando o PostgreSQL como banco de dados. A aplicação permite a criação, leitura e manipulação de produtos, com base na arquitetura de repositórios e casos de uso.

## Tecnologias Utilizadas

- **Go** (1.24.2)
- **PostgreSQL**
- **Gin** (framework web)
- **Docker** (para containerização)

## Funcionalidades

- **Criar produto**: Adiciona um novo produto ao banco de dados.
- **Obter todos os produtos**: Retorna uma lista de todos os produtos cadastrados.
- **Obter produto por ID**: Retorna um produto específico com base no ID fornecido.

## Estrutura do Projeto

- **repository**: Camada responsável pela interação com o banco de dados.
- **use_case**: Camada de lógica de negócios.
- **model**: Contém os modelos das entidades (como o `Product`).
- **cmd**: Contém o ponto de entrada da aplicação.

## Como Rodar a Aplicação

### Pré-requisitos

- Go (1.24.2)
- Docker e Docker Compose

### Passos para Rodar Localmente

1. Clone o repositório:

    ```bash
    git clone https://github.com/seu-usuario/seu-repositorio.git
    cd seu-repositorio
    ```

2. Construa e rode os containers Docker:

    ```bash
    docker-compose up --build
    ```

3. A aplicação estará disponível em `http://localhost:8080`.

### Endpoints

- **POST** `/products` - Cria um novo produto. Exemplo de body:

    ```json
    {
        "name": "Produto A",
        "price": 10.5
    }
    ```

- **GET** `/products` - Retorna todos os produtos.

- **GET** `/products/{id}` - Retorna um produto específico pelo ID.

## Como Contribuir

1. Fork o repositório
2. Crie uma nova branch (`git checkout -b feature/nova-funcionalidade`)
3. Faça commit das suas alterações (`git commit -am 'Adiciona nova funcionalidade'`)
4. Envie para o repositório remoto (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
