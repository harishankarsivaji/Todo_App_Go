# Todo_App_Go

![Travis (.com)](https://img.shields.io/travis/com/harishankarsivaji/Todo_App_Go)

Full stack CRUD application using Go and React

## Backend API architecture

- Using Golang to create REST API's `GET, PUT, POST, DELETE, OPTION`
- Routing through Gin-Gonic framework because of its lightweight and faster performance 
- MongoDB 
- Creating various middleware function which performs the CRUD operations in mongoDB. 
- A custom middleware to handle preflight `CORS` checks on the server side.

## Front end 

A simple React application which consumes the REST API's to create, update and delete tasks. Using the semantic UI to create styled components.

## Workflow and Future Works

- [x] Backend API
- [x] React front end 
- [ ] Unit Testing 
- [x] Dockerize the API 
- [x] Continous Integration like CircleCI / Travis CI
- [ ] Mongodb Atlas cluster connection string
- [ ] Deploy in AWS (EBS, ECS)
- [ ] JWT - User Login

