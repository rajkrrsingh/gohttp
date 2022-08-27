A simple http server to listen on the configured port to simulate apps running behind a reverse proxy.
#### Build Docker image
```
docker build -t gohttp .  
```

#### Run Docker image
```

docker run -e "PORT=1111" -p 1111:1111 -h app1 -t  gohttp 
docker run -e "PORT=2222" -p 2222:2222 -h app2 -t  gohttp 
docker run -e "PORT=3333" -p 3333:3333 -h app3 -t  gohttp 
```