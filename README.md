# SuperTests

## :ballot_box_with_check: Repositório do projeto supertests

Este projeto é um portal de testes interativo com integração com o Facebook

<p align="center">
  <img src="upload/o-que-o-coelhinho-da-pascoa-vai-trazer-para-voce/coelinho230317.jpg" />
</p>

### Dependências
Para instalar este projeto, é necessário ter o docker-compose e o docker instalado.

### Instalação e execução
Para executar, você pode rodar o seguinte comando abaixo
```
make run
```

Para interromper os containers gerados e refazer o banco de dados, execute o comando abaixo:
```
make down
```

### Testes
Para executar os testes, rode o seguinte comando abaixo
```
make test
```