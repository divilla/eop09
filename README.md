# eop09
### 📖 Go, fully featured, gRPC client and server demo, with REST client built using echo framework and server backed by MongoDB.

## 🏃‍♀️ Run
```sh
# run everything with docker-compose
docker compose up

# run from client - check:
cd client
# list make options
make

# run from server - check:
cd server
# list make options
make
```

## 📚 What it does?
- best to try with **postman**
- **http://127.0.0.1:8080/import** imports client/data/ports.json via gRPC to MongoDB connected to server
- **http://127.0.0.1:8080/ports?page=1** will list pages of ports
- **http://127.0.0.1t:8080/ports** provides full REST interface: GET, GET/:key, POST, PATCH/:key, PUT/:key, DELETE/:key
- client is built with Go **echo** framework
- server is Go gRPC server with: KeepaliveEnforcementPolicy, StreamInterceptor, UnaryInterceptor, GracefulShutdown and HealthCheck demo implementation
