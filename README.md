# Guide

## Run in local

```
go run main.go
```
### Run in docker
### Build image docker 
```
docker build -t ahfrd/esb-assesment:v1 .
```

### Run esb-assesment image on docker
```
docker run -d -p 9018:9018 -v config:/app/config --name esb-assesment-v1 ahfrd/esb-assesment:v1
```


#### Import Database
```
Import database from folder db esb in zip and file name esb.sql to your dbms example (dbeaver,datagrip dan lain lain) 
```

##### Import postman
```
Import collection postman from folder collection postman esb in zip
```