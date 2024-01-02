# SENSORS
Sensors and human Apps using multiple microservices

### Tech Stack
- Golang
- gRPC
- Echo Framework
- MySQL

### How to run
- Complete `config.toml` file as your configuration
- To run REST API :
  `go run cmd/apiService/main.go`
- To run gRPC :
  `go run cmd/grpcService/main.go`

### Example gRPC payload
- Insert Sensor :
```
{
    "id_1": "A",
    "id_2": 1,
    "timestamp": "2023-01-01",
    "type": "sensor",
    "value": 55.1
}
```
- Get Sensors 
- Get Sensor by id :
```
{
    "id_1": "A",
    "id_2": 1
}
```
- Edit Sensor :
```
{
    "filter_id1": "A",
    "filter_id2": 1,
    "id_1": "B",
    "id_2": 2,
    "timestamp": "2023-02-02",
    "type": "Sensor",
    "value": -98.81
}
```
- Delete Sensor :
```
{
    "id_1": "B",
    "id_2": 2
}
```
