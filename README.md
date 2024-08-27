# ourchain-agent

aid server with ourchain

## How to use?

please refer to the following file:
[api.md](./doc/api.md)

## How to use docker-compose to deploy?

quick start:

1. run `docker volume create node_data` start disk volume
2. run `docker-compose up -d` in this project root directory
3. use `docker-compose down` in this project root directory to stop container

note:

- `docker volume create node_data` 建立 node_data volume 作為節點資料外部儲存區
- `docker volume rm node_data` 刪除 node_data volume
- `docker-compose up -d` 啟動容器群集
- `docker-compose down` 停止容器群集
- `docker-compose pull` 更新容器群集
- `docker run -it --rm -v node_data:/node busybox sh` 進入節點資料外部儲存區
- `sudo docker exec -it go-aid-chain-1 /bin/bash` 進入區塊鏈(可以用 `bitcoin-cli generate 1` 重置)
- `docker-compose logs -f` 查看容器群集日誌
- `docker-compose ps` 查看容器群集狀態

## How to use nginx to proxy aid services?

- `sudo apt install nginx` 安裝 nginx
- `sudo nginx -t` 檢查 nginx 設定檔語法
- `sudo systemctl restart nginx` 重啟 nginx
- `sudo systemctl status nginx` 查看 nginx 狀態
- `sudo systemctl stop nginx` 停止 nginx
- `sudo systemctl start nginx` 啟動 nginx
- `sudo systemctl enable nginx` 設定 nginx 開機自動啟動
- `sudo vim /etc/nginx/nginx.conf` 編輯 nginx 設定檔
