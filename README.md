# go-poc-archref
POC to define an architecture reference in Golang

# run
- Install docker and docker-compose
- From the root project folder run: docker-compose -f **docker/compose-app.yml** up --build
- Optionally, you can add **-d** to free your terminal: docker-compose -f **docker/compose-app.yml** up --build -d

# usage
- POST -> *http://localhost:3000/api/v1/customers*
  body: take one item of customers array found into tests folder
  
# next steps
- Build others http resources

# utils
- Docker installation: https://docs.docker.com/install/
- Docker-compose installation: https://docs.docker.com/compose/install/
- Golang: https://golang.org/
