# Get User

## Valid

```
curl http://localhost:8080/users?id=123
```

## NotFound 404
```
curl http://localhost:8080/users?id=404
```

## Invalid 400
```
curl http://localhost:8080/users?id=abcd
```

# Create User
Note: please check your field names are matching with json in the request.

## Windows

### Valid

```
curl -X POST -H "Content-Type: application/json" -d "{\"firstName\": \"john\", \"lastName\": \"doe\", \"phone\": \"1234435\", \"age\":123, \"phoneVerified\":true}" http://localhost:8080/users
```

### Invalid Age

```
curl -X POST -H "Content-Type: application/json" -d "{\"firstName\": \"john\", \"lastName\": \"doe\", \"phone\": \"1234435\", \"age\":1235, \"phoneVerified\":true}" http://localhost:8080/users
```

## MacOS/Linux

### Valid

```
curl -X POST -H "Content-Type: application/json" -d '{"firstName": "john", "lastName": "doe", "phone": "1234435", "age":123, "phoneVerified":true}' http://localhost:8080/users
```

### Invalid Age

```
curl -X POST -H "Content-Type: application/json" -d '{"firstName": "john", "lastName": "doe", "phone": "1234435", "age":1235, "phoneVerified":true}' http://localhost:8080/users
```

# Delete User

## Valid

```
curl -X DELETE http://localhost:8080/users?id=123 -v
```

## NotFound 404
```
curl -X DELETE http://localhost:8080/users?id=404
```

## Invalid 400
```
curl -X DELETE http://localhost:8080/users?id=abcd
```
