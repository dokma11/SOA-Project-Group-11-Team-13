<h1 align='center'>Service-oriented architectures</h1>

<div>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
  <img src="https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white" />
  <img src="https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white" />
  <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white" />
  <img src="https://img.shields.io/badge/Neo4j-018bff?style=for-the-badge&logo=neo4j&logoColor=white" />
  <img src="https://img.shields.io/badge/Fluentd-599CD0?style=for-the-badge&logo=fluentd&logoColor=white&labelColor=599CD0" />
  <img src="https://img.shields.io/badge/Prometheus-000000?style=for-the-badge&logo=prometheus&labelColor=000000" />
  <img src="https://img.shields.io/badge/Grafana-F2F4F9?style=for-the-badge&logo=grafana&logoColor=orange&labelColor=F2F4F9" />
  <img src="https://img.shields.io/badge/Elastic_Search-005571?style=for-the-badge&logo=elasticsearch&logoColor=white" />
  <img src="https://img.shields.io/badge/Kibana-005571?style=for-the-badge&logo=Kibana&logoColor=white" />  
</div>

# Overview:
This project is part of the course on Service-Oriented Architectures and involves refactoring a previously developed monolith project into a set of microservices. The aim is to decompose the monolith into manageable, loosely coupled services that communicate over a network, providing a more scalable, flexible, and maintainable architecture.

# Microservices:
- Tours: Responsible for managing tour-related functionalities such as tours, their key points, equipment, facilities and reviews.
- Blogs: Responsible for managing blog-related functionalities such as blogs, their comments, votes and recommendations.
- Followers: Responsible for managing users followings and followers.
- JWT: Responsible for managing authentication using JSON Web Tokens.
- TourSearch: Implements Elasticsearch for efficient search and filtering of tours.

Other relevant parts from the previously made monolith project:

[API Gateway](https://github.com/dokma11/soa-group-11-team-13-back-end)

[Front-end](https://github.com/dokma11/soa-group-11-team-13-front-end)

# Technologies:
- Golang: Programming language for the microservices.
- PostgreSQL: Relational database used in tours microservice.
- MongoDB: NoSQL document database used in blogs microservice.
- Neo4j: Graph database used in followers microservice.
- Docker: Containerization tool to ensure consistent environments.
- Prometheus: Monitoring and alerting toolkit.
- Jaeger: Distributed tracing system.
- Loki: Log aggregation system.
- Fluentd: Data collection and log forwarding.
- Grafana: Data visualization & monitoring.
- Elasticsearch: Distributed, search and analytics engine designed for horizontal scalability and real-time search capabilities.
- Kibana: Data visualization and exploration tool used for visualizing Elasticsearch data and navigating the Elastic Stack.
- Remote Procedure Call (RPC): Messaging protocol.

# Getting started
<h3>Prerequisites</h3>

- Git: Ensure you have Git installed. You can download it from [here](https://git-scm.com/downloads).
- Docker: Ensure you have Docker installed. You can download it from [here](https://docs.docker.com/desktop/install/windows-install/).

<h3>Steps</h3>

1. Clone the repository<br>
    git clone https://github.com/dokma11/SOA-Project-Group-11-Team-13.git
2. Navigate to monitoring directory<br>
    cd ./src/tours/monitoring
3. Run docker compose for monitoring stack<br>
    docker compose up
4. Navigate to root directory<br>
    cd ../..
5. Run docker compose for the whole project<br>
    docker compose up

# Authors:
- [Veljko Nikolić RA 121/2020](https://github.com/Veljko121)
- [Spasoje Brborić RA 107/2020](https://github.com/spasoje2001)
- [Nina Kuzminac RA 119/2020](https://github.com/kuzminacc)
- [Vukašin Dokmanović RA 89/2020](https://github.com/dokma11)
