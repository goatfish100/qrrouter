curl -H "Contect-Type: application/json" -X POST -d {"id":3,"description":"yahoo","protected":false,"action":"forward","address":"http://www.yahoo.com"}` http://localhost:8001/createorg/




curl -H "Contect-Type: application/json" -X POST -d {"username":"jose","email":"test@yahoo.com","name":"jose cuevo","password":"password"}` http://localhost:8001/createuser/