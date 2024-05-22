# Hotel reservation backend

# Project environment variables
```
HTTP_LISTEN_ADDRESS=:4000
JWT_SECRET=somethingsupersecretthatNOBODYKNOWS
MONGO_DB_NAME=hotel-reservation
MONGO_DB_URL=mongodb://localhost:27017
MONGO_DB_URL_TEST=mongodb://localhost:27017
```

## Project outline
- users -> book room from a hotel
- admins -> going to check reservation/bookings
- Authentication and authorization -> JWT tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database management -> seeding, migration

## Resources
### Mongodb driver
Documentation
```
https://mongodb.com/docs/drivers/go/current/quick-start
```

Installing mongodb client
```
go get go.mongodb.org/mongo-driver/mongo
```

### go fiber
Documentation
```
https://gofiber.io
```

Installing go fiber
```
go get github.com/gofiber/fiber/v2
```

## Docker
### Installing mongodb as a Docker container
```
docker run -d -p 27017:27017 --name test-mongo mongo:latest
