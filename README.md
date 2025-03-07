<div align="center">
<h1>Technical Assessment</h1>
<h2>SOA Microservices</h2>
</div>

# Intructions

## Folders structure
1. **apis**: define paths of apis
2. **repositories**: define **mongodb** connection
3. **services**: define logics
4. **common**: define supported packages
5. **models**: define structs
6. **servers**: define server to run api
7. **tmp**: include file log of this project
8. **build**: bash to build source code

## Detailed steps to config and run api
1. You copy .env_example to .env and write approriate configuration data to .env
```
# PostgreSQL
PG_HOST=
PG_DATABASE=
PG_USERNAME=
PG_PASSWORD=
PG_PORT=
```

2. Use 2 below curls to test apis, run on your terminal or postman:


Notice: you should replace <i><b>http://localhost:8000</b></i> with your domain api when you deployed on your server. If you run 2 curls on your local terminal, please skip this notice.


3. Run source code

You can run source code on your computer or your server:
```
go run main.go
```
Or use file .exe:
```
./build/build.sh
```

this command will exported app.exe file, helping you to deploy.

I also have already created a **Dockerfile**, you can run to build an image and deploy to the server.

Build Docker file:

```
docker build --pull --rm -f "Dockerfile" -t soa:latest .
```


<div align="center">
Copyright © 2025 AnhLe. All rights reserved.
</div>