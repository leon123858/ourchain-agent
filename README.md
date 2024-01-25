# go-aid

aid server with ourchain

## How to use?

read functional test in [test](./test) directory

## How to run deploy by docker-compose?

1. run `docker-compose up -d` in this project root directory

note: use `docker-compose down` in this project root directory to stop container

## How to use nginx to proxy aid services?

- `sudo apt install nginx` 安裝 nginx
- `sudo nginx -t` 檢查 nginx 設定檔語法
- `sudo systemctl restart nginx` 重啟 nginx
- `sudo systemctl status nginx` 查看 nginx 狀態
- `sudo systemctl stop nginx` 停止 nginx
- `sudo systemctl start nginx` 啟動 nginx
- `sudo systemctl enable nginx` 設定 nginx 開機自動啟動
- `sudo vim /etc/nginx/nginx.conf` 編輯 nginx 設定檔