# HowRU
This go program is the frontend server that connects to another small client called Buddy that sits on the server.

The Buddy Server can be found [Here](https://github.com/dminGod/Buddy)

### HowRU Components
3 Major components of this application right now :


Architecture Diagram:
![Alt text](/architecture.jpg?raw=true "HowRU Architecture")


- HTTP server to serve - Frontend Web Interface
- Web Sockets based server to service requests from the Frontend
- GRPC client to talk to Buddy clients that are deployed on remote servers


#### Frontend Web Interface & HTTP server
This is an angular based Frontend. Here is where you will see all the data related to your nodes.

- Metrics
- Logs
- Alarms
- Buttons to execute remote actions
- (...) Any other features we chose to add in the future

This is meant to be extensible which means if there is a module today to monitor a cassandra cluster later it can be
easily extended to monitor a PostgresXL cluster etc.

All the functionality is achieved using a Web Socket connection to the HowRU server

#### Web Sockets Server
This is also part of the HTTP server. The Web Sockets port is used to send and recieve messages from the HTTP interface to the HowRU server.

#### GRPC client to talk to Buddy Servers
This client has a list of servers that have the Buddy Servers installed on them. Based on the request from the frontend
a gRPC message request will be sent to the remote server, a response will be returned from the Buddy client over gRPC,
this will be processed and returned back to the frontend interface which will display this content to the end user.


## Buddy Clients
On each of the the remote servers a Buddy client is installed the Buddy Client consists of the following :
- GRPC Server
- Execution module

#### gRPC Server
This is a server that listens to requests sent by HowRU, processes it based on the execution engine and responds back
with a the reply.

#### Execution Module
This is part of the application that actually executes the requests after they have come, and returns back a response.



## Notes
- Frontend currently do not have authentication on it but will have users roles and rights as
  the application grows. JWT is tokens are planned for this so Authentication and Authorization both can be handled.
- There is a whitelist based + another factor authentication mechanism planned for communication between the HowRU Server and Buddy Server.


