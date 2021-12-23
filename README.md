#### Executar em modo interativo
* docker run -it \$\{container}

#### Buildar
* docker build -t \$\{nome container} ${path_dockerfile}

#### Mapear porta host:container
* docker -p 12345:4001

#### Iniciar o container
* docker start -i \$\{container}

#### Executar um comando dentro de um container em execução
* docker exec -it \$\{nome_container} \$\{comando, ex: bash}
