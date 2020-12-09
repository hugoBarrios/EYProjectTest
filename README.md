# serve localhost:3300
go run main.go

# build exe
go build main.go

## API routes
#Get all students
method "GET" route "/api/v1/student"

#Get a students
method "GET" route "/api/v1/student/{id}"

#Create a new students
method "POST" route "/api/v1/student"

#Delete a students
method "DELETE" route "/api/v1/student/{id}"

