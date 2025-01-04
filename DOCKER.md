## Docker imagem pronta no dockerHub

## Criando uma imagen e push para dockerhub a partir do projeto.

1. **Deletar imagens antigas**
    ```bash
        docker rmi frist-go:0.0.1
    ``` 
2. **Build do projeto**
    ```bash
        docker build -t frist-go:0.0.1 .
    ``` 
3. **Tag do build**
   ```bash
        docker tag 6ce2760d343b lucasgalo/frist-go:0.0.1
   ```
4. **Login dockerhub**
   ```bash
        docker login
   ```
5. **Push da tag**
    ```bash
        docker push lucasgalo/frist-go:0.0.1
    ```

## Container create example:

1. **Create work directories:**
    ```bash
    mkdir -p /home/frist-go/logs    
    ```
2. **Permissao de leitura e escrita para a pasta**
   ```bash
     sudo chmod 777 -R /home/frist-go/logs
   ```
3. **Create container:**
    ```bash
    docker run \
    --name frist-go \
    --hostname ambienteclean \
    --detach --restart unless-stopped \
    --publish 8080:8080 \
    --volume /home/frist-go/logs:/opt/ambienteclean/logs:rw \
    --volume /home/frist-go/application-prod.properties:/opt/ambienteclean/src/resources/application-prod.properties:ro \
    lucasgalo/frist-go:0.0.1
    ```   
4. **Create container, servidor externo**
    ```bash
     docker run --name frist-go --network="host" --detach --restart unless-stopped --publish 8080:8080 --volume /home/frist-go/logs:/opt/ambienteclean/logs:rw --volume /home/frist-go/application-prod.properties:/opt/ambienteclean/src/resources/application-prod.properties:ro lucasgalo/frist-go:0.0.1
    ```
5. **Other option create container local:**
    ```bash
    docker run --name fristgo --detach --restart unless-stopped --publish 8080:8080 fristgo:0.0.1
    ```
6. **Olhar os logs do container criado:**
    ```bash
    docker container logs 7b2   
    ```

