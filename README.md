# shortcut
Shortcut is a web service that written in golang for people to create short URLs that can be easily shared.

Online: [online_shortcut_services](https://share.mdzz.dev/)

## How to use it

### Prepare

1. go environment
2. mysql
3. redis
4. google reCAPTCHA v2

After you have prepared, run `go get github.com/grt1st/shortcut`

### Set config file

The config file is at `./conf/default.yaml`, and you have to edit it. Here is an example:

```
protocol: http  # http or https
host: ip:port   # listen on where, such as 127.0.0.1:8001
domain: domain  # your domain, such as github.com. you can also set it the same as host
mysql: mysql//user:password@(host:port)/database?charset=utf8&parseTime=True&loc=Local # mysql client
redis: 127.0.0.1:6379//password  # redis client
gkey: your_site_client_key       # you key of google recaptcha
gsecret: your_site_server_key    # you secret of google recaptcha
```

and my config file is:

```
protocol: https
host: 127.0.0.1:8001
domain: mdzz.dev
mysql: mysql//root:xxxxxx@(127.0.0.1:3306)/ddd?charset=utf8&parseTime=True&loc=Local
redis: 127.0.0.1:6379//
gkey: 6LfRopgUAAAxxxxxxxxxxx
gsecret: 6LfRopgUAAxxxxxxxx
```

### Start to run

Into the directory first(`$GOPATH/github.com/grt1st/shortcut`), then you can simple run it by `go run main.go` or build it use `go build -o short && ./short`.

### Nginx config

```
upstream share_pool{
    server 127.0.0.1:8001;
}


server {
	listen 80;
	listen [::]:80;

	server_name share.mdzz.dev;

	access_log /var/log/nginx/share.log;
    error_log /var/log/nginx/share.error;

    location /static/~ {
        root   /var/www/static/; # static file dir, make sure to have permission
    }

	location / {
        proxy_set_header Host $host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://share_pool;
	}
	client_max_body_size 20m;
}
```

## todo

1. beautify ui
2. add no mysql, no redis, no captcha support
3. add google reCAPTCHA v3 support