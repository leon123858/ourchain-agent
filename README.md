# go-aid

aid server with ourchain

## How to use?

read [chainAPI](./chain.http)

## How to run deploy by docker-compose?

1. install docker and docker-compose
2. build container image follow [ourChain](https://github.com/leon123858/OurChain/blob/main/doc/dev-docker.md#%E7%99%BC%E5%B8%83-container)
3. run `docker build -t go-aid .` in this project root directory
4. run `docker-compose up -d` in this project root directory

note: use `docker-compose down` in this project root directory to stop container