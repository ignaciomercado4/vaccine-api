## Vaccine API backend
REST API made following a take-home coding challenge. Task available [here](https://github.com/ignaciomercado4/vaccine-api/blob/master/TASK.md). Made using Go, Gin, Gorm and PostgreSQL.

### Request examples
User creation:
```
curl -X POST http://localhost:8080/api/signup \
-H "Content-Type: application/json" \
-d '{
  "name": "test",
  "email": "test@email.com",
  "password": "password"
}'
```

User login:
```
curl -X POST http://localhost:8080/api/login \
-H "Content-Type: application/json" \
-d '{
  "name": "test",
  "email": "test@email.com",
  "password": "password"
}'

```

Drug creation:
```
curl -X POST http://localhost:8080/api/drug \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>" \
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
curl -X GET http://localhost:8080/api/drugs \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>"
```

Delete a drug:
```
curl -X DELETE http://localhost:8080/api/drugs/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>"
```

Modify an existing drug:
```
curl -X PUT http://localhost:8080/api/drugs/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>" \
-d '{
  "name": "drug1modified",
  "approved": false,
  "minDose": 2,
  "maxDose": 2,
  "availableAt": "2025-12-22T00:00:00Z"
}'
```

Vaccination creation:
```
curl -X POST http://localhost:8080/api/vaccination \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>" \
-d '{
  "name": "vaccination1",
  "drugId": 1,
  "dose": 2,
  "date": "2025-12-23T00:00:00Z"
}'
```

Get all vaccinations:
```
curl -X GET http://localhost:8080/api/vaccination \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>"
```

Delete a vaccination:
```
curl -X DELETE http://localhost:8080/api/vaccination/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <your_jwt_token>"
```

### Frontend
Not done yet :P.
