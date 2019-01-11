grep -i docker /etc/group

sudo gpasswd -a $(whoami) docker

grep -i docker /etc/group
