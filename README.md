# go-poc-archref
POC to define an architecture reference in Golang

# run
- Install docker and docker-compose
- From the root project folder run: docker-compose -f **docker/compose-app.yml** up --build
- Optionally, you can add **-d** to free your terminal: docker-compose -f **docker/compose-app.yml** up --build -d

# usage
- POST: *http://localhost:3000/api/v1/customers*
    - body (json):  
{  
 "name": "Gerald",  
 "birthYear": 1768,  
 "birthDay": 6,  
 "birthMonth": 1,  
 "eMail": "gerald@witcher-mail.com"  
}
    - other payloads on: https://github.com/vagner-nascimento/go-poc-archref/blob/master/tests/payloads.json
 
 - GET: *http://localhost:3000/api/v1/customers/{id}*

# next steps
- Build others http resources

# utils
- Docker installation: https://docs.docker.com/install/
- Docker-compose installation: https://docs.docker.com/compose/install/
- Golang: https://golang.org/
