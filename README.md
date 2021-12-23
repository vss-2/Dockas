#### Executar em modo interativo
* docker run -p 4001:4001 --net=bridge -it \$\{container}

#### Buildar
* docker build -t \$\{nome container} ${path_dockerfile}

#### Executar container em moto detach (no background)
* docker run -d \$\{nome_container}

#### Mapear porta host:container
* docker -p 12345:4001

#### Iniciar o container
* docker start -i \$\{container}

#### Executar um comando dentro de um container em execução
* docker exec -it \$\{nome_container} \$\{comando, ex: bash}
* Obs: -i (interative) -t (tty) e -d (detach) podem ser usados juntos em -dit

#### Criando uma rede interna para containers
* docker network create \$\{nome_container_net} --driver bridge >> container_net_id.txt
* Obs: adicionar --network \$\{nome_container_net} para runs ou execs

#### Ver quais containers estão conectados na bridge
* docker inspect \$\{nome_container_net}
