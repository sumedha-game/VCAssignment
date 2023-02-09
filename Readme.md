# API

## SignUp
```
GET /signUP HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 56

{
    "username":"user1",
    "password":"New@123"
}

A User will be added in the DB.
```

## SignIn
```
GET /login HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Content-Length: 56

{
    "username":"user1",
    "password":"New@123"
}

Use the token generated in the response in the Bearer Token.
```

## CreateEmployee
``` 
POST /createNew HTTP/1.1
Host: localhost:8080
Authorization: Bearer <--token-->
Content-Type: application/json
Content-Length: 111

{
    "name":"Tony",
    "department":"mgr",
    "city":"LA",
    "salary":52025.54
}

```

## GetEmployee
``` 
POST /createNew HTTP/1.1
Host: localhost:8080
Authorization: Bearer <--token-->
Content-Type: application/json
Content-Length: 111

{
    "name":"Tony",
    "department":"mgr",
    "city":"LA",
    "salary":52025.54
}
```

## UpdateEmployee
``` 
PUT /update?id=2 HTTP/1.1
Host: localhost:8080
Authorization: Bearer <--token-->
Content-Type: application/json
Content-Length: 111

{
    "name":"Tony",
    "department":"mgr",
    "city":"LA",
    "salary":25000.54
}
```

## DeleteEmployee
``` 
DELETE /delete?id=1 HTTP/1.1
Host: localhost:8080
Authorization: Bearer <--token-->
```