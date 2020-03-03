# go-poc-archref
POC to define an architecture reference in Golang

# run
- Install docker and docker-compose
- From the root project folder run: docker-compose -f **docker/compose-app.yml** up --build
- Optionally, you can add **-d** to free your terminal: docker-compose -f **docker/compose-app.yml** up --build -d

# usage
- **HTTP/POST**: *http://localhost:3000/api/v1/customers*
    - body (json):  
{  
 "name": "Gerald",  
 "birthYear": 1768,  
 "birthDay": 6,  
 "birthMonth": 1,  
 "eMail": "gerald@witcher-mail.com"  
}
    - other customer's payloads on: https://github.com/vagner-nascimento/go-poc-archref/blob/master/tests/payloads.json
 
 - **HTTP/GET**: *http://localhost:3000/api/v1/customers/{id}*
 - **HTTP/PUT**: *http://localhost:3000/api/v1/customers/{id}*
    - body (json):  
      {  
       "name": "Gerald",  
       "birthYear": 1768,  
       "birthDay": 6,  
       "birthMonth": 1,  
       "eMail": "gerald@witcher-mail.com",  
"userId": "026616e2-063a-408b-aa45-99dc671081db"  
}
 
 - **AMQP/UPDATE**:
    - insert a customer through **HTTP/POST** method 
    - access *http://localhost:15672/#/queues/%2F/q-user*
    - on  **Publish message** menu, put a user that have the same e-mail address of inserted customer into **Payload** field
    - click on **Publish message**
    - to verify changes, call **HTTP/GET** with inserted customer's id
    
 - **AMQP/GET**: operations that change customer's data are published into customer's queue.
    - access **http://localhost:15672/#/queues/%2F/q-customer**
    - on **Get messages** menu increase **Messages** field to 50 (or more) and click on **Get Message(s)**

# next steps
- Build get list by query params and patch

# utils
- Docker installation: https://docs.docker.com/install/
- Docker-compose installation: https://docs.docker.com/compose/install/
- Golang: https://golang.org/
