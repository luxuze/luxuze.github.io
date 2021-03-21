# Kong

1. kong database

   ```bash
    docker run -d --name kong-database \
                    -p 5432:5432 \
                    -e "POSTGRES_USER=kong" \
                    -e "POSTGRES_DB=kong" \
                    postgres:9.6
   ```

2. kong network

   ```bash
    docker network create kong-net
   ```

3. kong volume

   ```bash
    docker volume create kong-vol
    docker volume inspect kong-vol

        [
            {
                "CreatedAt": "2019-05-28T12:40:09Z",
                "Driver": "local",
                "Labels": {},
                "Mountpoint": "/var/lib/docker/volumes/kong-vol/_data",
                "Name": "kong-vol",
                "Options": {},
                "Scope": "local"
            }
        ]
   ```

4. Config File
   * /var/lib/docker/volumes/kong-vol/\_data/kong.yml
   * run docker image

     ```bash
     docker run -d --name kong \
        --network=kong-net \
        -v "kong-vol:/usr/local/kong/declarative" \
        -e "KONG_DATABASE=off" \
        -e "KONG_DECLARATIVE_CONFIG=/usr/local/kong/declarative/kong.yml" \
        -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
        -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
        -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
        -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
        -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
        -p 8000:8000 \
        -p 8443:8443 \
        -p 8001:8001 \
        -p 8444:8444 \
        kong:latest
     ```

## Useful cmds

* Reload Kong in a running container

```bash
docker exec -it kong kong reload
```

