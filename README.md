# AspecT-Backend
<img src="./.github/images/Cover.jpg">
Aspect Backend

## About
 Motivation behind this project is, we care about indonesian for make it better next generation 


## Feature V1
-   Track all kid's bad words text
-   Add New Children to Monitorize
-   Report Daily, Weekly, Montly Progress
-   Achievement Good progress children (Soon)
-   Autotext Good Word (Soon)
## Documentation
Before start the project localy, make sure to all tools has been installed on your machine such:

- [X]    Golang

- [X]    Python

- [X]    Docker


## About The Project(Technical)
#### Apps (Service)
> You may ask why we use golang?

we use golang because it's FAST and the framework that we use is fiber, the fastest server, just take a look at this benchmark
<img src="./.github/images/benchmark-pipeline.png">
> benchmark from other framework

for the archicteture it self, we use DDD/Hexagonal in golang, MVC(Model View Controller) in python, and use Event Driven Architecture for build microservice between each service [See reference here](./#reference)

#### Reference
-   [DDD](https://engineering.grab.com/domain-driven-development-in-golang)
-   [Event Driven](https://medium.com/bliblidotcom-techblog/event-driven-architecture-ef3a312180ee)
-   [Microservice](https://www.nginx.com/blog/microservices-at-netflix-architectural-best-practices/)
-   [MVC](https://shravan-c.medium.com/mvc-for-flask-application-a636e6f58d72)

#### Additional Tool's and Service


#### HOW TO RUN
##### GOLANG

```bash
cd backend/apps
```
```bash
make run
```

##### PYTHON
to build image for docker
```bash
make build

```

to run the image as container

```bash
make run
```

stop the container
```bash
make stop
```

remove image 
```bash
make remove
```

## Resources

## Todo's Apps
- [ ] Account Service

- [ ] Auth Service

- [ ] Children Service

- [ ] Profanity Service

- [ ] Reporting Service

- [ ] AI Service (using AI Google Platform)




## Team

-   [@xxidbr9](https://github.com/xxidbr9) 
    -   fullname    : Barnando Akbarto Hidayatullah (Nando)
    -   [linkedin](https://linkedin.com/in/xxidbr9)
-   [@M4RIONETTE](https://github.com/M4RIONETTE)
    -   fullname : Fajar Nur Hidayatullah (Fajar)
    -   [linkedin](https://www.linkedin.com/in/marionette/)