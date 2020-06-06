---
title: 
author: Luis Rodriguez
type: post
date: 2020-01-24T00:57:15+00:00
url: /post/2020/01/24/personal-cloud-setup
draft: true
categories:
  - Projects
tags:
  - nas
  - docker
  - all-in-one
  - iot
  - home-server
  - server
  - cloud

---


apt-get install software-properties-common
apt-key adv --recv-keys --keyserver hkp://keyserver.ubuntu.com:80 0xF1656F24C74CD1D8
add-apt-repository "deb [arch=amd64,arm64,ppc64el] http://mariadb.mirror.liquidtelecom.com/repo/10.4/ubuntu $(lsb_release -cs) main"
apt update && apt -y install mariadb-server mariadb-client


If you were not prompted to set root password, run:
mysql_secure_installation



nano /etc/mysql/my.cnf

To optimize MySQL/Mariadb which is installed on 1gb VPS you need to add this configs in my.cnf under [mysqld] :

[mysqld]

#change these
max_allowed_packet = 4M
key_buffer_size = 32K
table_open_cache = 8
sort_buffer_size = 128K
read_buffer_size = 512K
read_rnd_buffer_size = 512K
innodb_file_per_table = 0

#add these
symbolic-links=0
skip-external-locking
net_buffer_length = 4K
thread_stack = 480K

systemctl restart mysqld



https://www.digitalocean.com/community/tutorials/how-to-optimize-nginx-configuration

apt-get install nginx -y

nano /etc/nginx/nginx.conf



https://linuxize.com/post/how-to-install-php-on-ubuntu-18-04/
https://guides.wp-bullet.com/adjusting-php-fpm-performance-low-memory/

apt install php-fpm php-mysql -y


nano /etc/nginx/sites-available/default

add under server{}

       location ~ \.php$ {
                include snippets/fastcgi-php.conf;
                fastcgi_pass unix:/var/run/php/php7.2-fpm.sock;
        }

systemctl restart nginx



configure your sites

add-apt-repository ppa:certbot/certbot
apt install cerbot python-certbot-nginx -y

certbot --nginx -d example.com -d www.example.com

add cron for certbot

certbot renew --dry-run

0 0 * * 0 certbot renew