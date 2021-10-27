# How to run without Docker Compose
The config file assumes a nodejs backend listening at web-node:8080 and a go backend listening at web-go:8090

In order to run nginx outside Docker Compose, you can put ```conf.d/default.conf``` at ```/etc/nginx/sites-available``` and edit it according to your setup.

You also need to put ```html``` folder at ```/usr/share/nginx/```
