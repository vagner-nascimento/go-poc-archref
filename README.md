# go-poc-archref
POC to define an architecture reference in Golang

#run
- Install docker and copose
- From the root project's folder run: docker-compose -f *docker/compose-app.yml* up --build

#usage
- POST -> *http://localhost:3000/api/v1/customers*
body
  {
   "name": "Gerald",
    "birthYear": 1768,
    "birthDay": 6,
    "birthMonth": 1,
    "eMail": "gerald@witcher-mail.com"
  }
