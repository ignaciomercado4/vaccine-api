## Vaccine API backend
REST API made following a take-home coding challenge. Task available [here](https://github.com/ignaciomercado4/vaccine-api/blob/master/TASK.md). Made using Go, Gin, Gorm and PostgreSQL.

### Request examples
User creation:
```
curl -X POST http://localhost:8080/createUser \
-H "Content-Type: application/json" \
-d '{
  "name": "test",
  "email": "test@email.com",
  "password": "password"
}'
```

User login:
```
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
  "name": "test",
  "email": "test@email.com",
  "password": "password"
}'

```

Drug creation:
```
curl -X POST http://localhost:8080/drug \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your_jwt_token" \
-d '{
  "name": "drug1",
  "approved": true,
  "minDose": 1,
  "maxDose": 1,
  "availableAt": "2024-12-22T00:00:00Z"
}'
```

Get all drugs:
```
curl -X GET http://localhost:8080/drugs \
-H "Content-Type: application/json" \
-H "Authorization: Bearer your_jwt_token"
```

### Frontend
Not done yet :P.
