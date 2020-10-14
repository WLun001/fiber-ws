Example of API server send message to websocket server

```bash
API Server -> WS server -> Browser (Not showing here)
```

### Run
Open terminal, run ws server
```bash
go run ws.go
```

Open another terminal tab, run api server
```bash
go run api.go
```

Open another terminal tab, run the following command to call the api server
```bash
curl http://127.0.0.1:8081
```
A message will appear on ws server terminal tab
