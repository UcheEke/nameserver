#Nameserver

A simple demonstration of Http in Angular 2

###Overview

Illustration of the use of the Angular 2 Http service. 

A basic HTML table is populated by JSON datagenerated by a simple webserver written in golang. The server provides a JSON array of 10 randomly generated names, each with additional 'age' and 'id' tags, per GET request.
Within the main angular file 'main.ts', the service 'PeopleSvc' retrieves this data from the webserver using Observables (as opposed to Promises in Angular 1). This service is then subscribed to by the App, and the data interpolated by its template. 

###Requirements
- node
- go 1.5.3

Run 'npm install' to load the dependencies in the package.json file
