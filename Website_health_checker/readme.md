1. To init project locally 

```
go mod init website_health_checker
```

2. To install packages used in the project

```
go mod tidy
```

3. To intall the cli package

```
go get github.com/urfave/cli/v2
```

4. To run and test for a particular domain 

```
go run . --domain google.com
```