# About this project
Simple Golang service for handling Account objects, built for a backend with a microservices architecture.
## Built with
- Golang
- gRPC
- Microservice architecture
## Why this stack?
I designed this microservice while working on a larger Multiplayer Game Development project. In online gaming, speed and latency are major factors that need to be accounted for. This stack meets this crateria for the following reasons :
- Golang : while not the absolute fastest language available, golang strikes a good balance between performance and developer friendliness.
- gRPC : to minimize latency, JSON is generally avoided as a communication protocol in game development. Protobuf is a more efficient protocol, and gRPC uses protobufs while also making it easy to work with them with zero performance tradeoff.
