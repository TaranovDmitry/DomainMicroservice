# Domain Service

### Prerequisites 
- docker
- migration tool

### How to start Domain Service with docker-compose
1. Start docker-compose: ``docker-compose up``
2. Execute DB Scheme Migration: ``migrate -path ./schema -database 'postgresql://postgres:12345@localhost:5432/postgres?sslmode=disable' up``

### API Request examples
1. Get all ports: <br></br>``curl -v --location --request GET 'http://localhost:8081/domain/v1/ports'``
2. Update ports: <br></br>``curl -v --location --request POST 'http://localhost:8081/domain/v1/ports' --header 'Content-Type: application/json' --data-raw '[{"id":0,"name":"consequat id officia","isActive":true,"company":"GEOFARM","email":"maymclean@geofarm.com","phone":"+1 (997) 434-3843","address":"907 National Drive, Foscoe, Oregon, 5061","about":"Id laborum \r\n","registered":"2021-02-14T01:46:57-02:00","latitude":70.822864,"longitude":156.088083}]'``