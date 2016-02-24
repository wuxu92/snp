## Shared Navigation Page Project (SNP) ##
personal project, using go, beego and mongo to implement a new navigation page


### How to Run
you should install begoo at first, Run snp as a begoo project.

snp only provide apis at 8080 port(or others you pointed), it renders no pages/views, so it need a http server before it.

make a vhost rooted at `paht/to/snp/html`, that is the `html` folder under project. proxy all request to snp server except the index.html request, for:

all requests are pointed to index.html under `html`. using try_files to redirect api requests to snp, so you should config you vhost as below(using nginx):

```
upstream go_snp {
    server 127.0.0.1:8080;
    keepalive 300;
}

location / {
    try_files $uri /index.html;
    root   E:\go\src\snp\html;
    index  index.html index.htm;
}

location /api {
    proxy_pass http://go_snp;
    proxy_http_version 1.1;
    proxy_set_header Connection "";
}
```
this config is a little tedious... though it is a really bad design, i like it. because i build it.

### setup MongoDB
what you should do before run snp in your server is config the MongoDB. config it at `/path/to/project/conf/app.conf`
replace the mgo host, port, username and password with your own.

### TODO
this project is partly  RESTful implemented, that is, for some apis they meet RESTful requires, but others designed at first is not.

so one important thing is to change all apis to meet RESTful