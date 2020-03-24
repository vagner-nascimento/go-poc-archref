# go-poc-archref
POC to define an architecture reference in Golang

# run
- Install docker and docker-compose
- From the root project folder run: docker-compose -f **docker/compose-app.yml** up --build
- Optionally, you can add **-d** to free your terminal: docker-compose -f **docker/compose-app.yml** up --build -d

# usage
- **HTTP/POST**: http://localhost:3000/api/v1/customers
    - body (json):  
{  
 "name": "Gerald",  
 "birthYear": 1768,  
 "birthDay": 6,  
 "birthMonth": 1,  
 "eMail": "gerald@witcher-mail.com"  
}
    - other customer's payloads on: https://github.com/vagner-nascimento/go-poc-archref/blob/master/tests/payloads.json
 - **HTTP/PUT**: http://localhost:3000/api/v1/customers/{id}
     - body (json):  
       {  
        "name": "Gerald",  
        "birthYear": 1768,  
        "birthDay": 6,  
        "birthMonth": 1,  
        "eMail": "gerald@witcher-mail.com",  
 "userId": "026616e2-063a-408b-aa45-99dc671081db"  
 }  
 - **HTTP/GET** (by id): http://localhost:3000/api/v1/customers/{id}
 - **HTTP/GET** (by params): http://localhost:3000/api/v1/customers?p=val&p1=val1&valArr=["it1", "it2"]
    - params can be any Customer attributes. If you want to search in a range of some value, send as array, like above
    - don't send multiple params with same name, like ?name=Jhon&name=Mary, it will consider only the first one   
  
 - **AMQP/UPDATE**:
    - insert a customer through **HTTP/POST** method 
    - access http://localhost:15672/#/queues/%2F/q-user
    - on  **Publish message** menu, put a user that have the same **e-mail** of inserted customer into **Payload** field
    - click on **Publish message**
    - to verify changes, call **HTTP/GET** with inserted customer's id
    
 - **AMQP/GET**: operations that change customer's data are published into customer's queue.
    - access **http://localhost:15672/#/queues/%2F/q-customer**
    - on **Get messages** menu increase **Messages** field to 50 (or more) and click on **Get Message(s)**

# next steps
- Develop delete and patch http endpoint
- Layout restructure
- Http errors handler
- App errors handler

# utils
- Golang: https://golang.org/
- Docker installation: https://docs.docker.com/install/
- Docker-compose installation: https://docs.docker.com/compose/install/
