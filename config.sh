# remove old docker
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
# install docker and docker-compose
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
# install nginx
sudo apt-get install nginx
# cp config file
sudo cp ./nginx.conf /etc/nginx/nginx.conf
# check ssl.crt and ssl.key if exist, if not, throw error
if [ ! -f "./ssl.crt" ]; then
  echo "ssl.crt not found, please put ssl.crt and ssl.key in the same directory as this script."
  exit 1
fi
if [ ! -f "./ssl.key" ]; then
  echo "ssl.crt not found, please put ssl.crt and ssl.key in the same directory as this script."
  exit 1
fi
# cp ssl.crt and ssl.key
sudo cp ./ssl.crt /etc/nginx/ssl.crt
sudo cp ./ssl.key /etc/nginx/ssl.key
sudo nginx -t
sudo systemctl restart nginx
