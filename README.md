# Backend Developer Assessment

## Requirement
To run this project, you need to install:

- [go](https://golang.org/dl/)
- [mysql](https://www.mysql.com/downloads/)

## Clone the project

```
$ git clone https://github.com/ItonLazaro/backend-assessment.git
$ cd backend-assessment
```

## Create a database

```
Open a NEW and SEPARATE terminal
$ mysql -u <username> -p 
$ <Enter password>
$ create database storia
```
## 


## Run the Program

```
Run MySQL in the background locally using Laragon/XAMPP (or any preferred local development tool)
Go back to /backend-assessment directory
$ go run main.go
```


## List of Endpoints
**For faster POSTMAN endpoint setup, you can import Storia.postman_collection.json file**
```
    [POST] localhost:8080/api/register
    [POST] localhost:8080/api/login


    Endpoints below needs a Bearer Token that is generated and a response from Login endpoint
    [GET] localhost:8080/api/task
    [POST] localhost:8080/api/task
    [PUT] localhost:8080/api/task/{task_id}
    [DELETE] localhost:8080/api/task/{task_id}
```
