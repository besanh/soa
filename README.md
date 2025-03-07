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
8. **pkgs**: include all third party libraries

## Detailed steps to config and run api
You copy .env_example to .env and write approriate configuration data to .env
To save time, I offer use secret key to secure apis.
So I provided secret key, so you only need to pass below field to .env file
```
SECRET=8HsicGYxLnf9xNmFjF5WuRgu1VHwcktDLQIR6EPMs8kTwJlBgT
```

```
# PostgreSQL
PG_HOST=
PG_DATABASE=
PG_USERNAME=
PG_PASSWORD=
PG_PORT=
```

Notice: you should replace <i><b>http://localhost:8000</b></i> with your domain api when you deployed on your server. If you run 2 curls on your local terminal, please skip this notice.


Run source code

You can run source code on your computer or your server:
```
go run main.go
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