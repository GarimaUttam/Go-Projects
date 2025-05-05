1. to create a go.mod file

```
go mod init ./
```

2. to run the project 

``
go run main.go
``

3. to run the Apis or routers run the below cmd on terminal

``
curl http/localhost:8080/books
``

4. to run the create cmd from the termincal using an json file data in the same folder

``
curl -Uri http://localhost:8080/books `
     -Method Post `
     -Headers @{ "Content-Type" = "application/json" } `
     -Body (Get-Content -Raw -Path .\body.json)
``

5. to find the book by the help of its Id 

``
curl.exe http://localhost:8080/books/8
``


6. to change the quantity of book - checkout

``
curl.exe http://localhost:8080/checkout?id=3 --request "PATCH"
``