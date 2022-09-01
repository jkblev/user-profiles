# How to build and start server
1. `cd` into `web-service-gin`
2. Run `go build` to compile the code into an executable.
3. On Linux or OS X, run `./web-service-gin`. On Windows, run `web-service-gin.exe`
4. You should see the following output: "Listening and serving HTTP on localhost:8080"

# How to interact with endpoints
## POST /users
1. Follow instructions for how to start server above.
2. Open terminal
3. Enter the following:
```
curl http://localhost:8080/users \
    --header "Content-Type: application/json" \
    --request "POST" \
  --data '[{"user_id": "1", "name": "Joe Smith", "date_of_birth":"1983-05-12", "created_on" : 1642612034},{"user_id": "2", "name": "Jane Doe", "date_of_birth":"1989-04-29", "created_on": 0}]'
```
4. You should receive JSON containing user data.


## GET /users
1. Follow instructions for how to start server above.
2. Open terminal
3. Enter the following:
```
curl http://localhost:8080/users \
--header "Content-Type: application/json" \
--request "GET"
```
4. You should receive JSON containing user data. May be empty if no users have been submitted yet.

## POST /image
1. Follow instructions for how to start server above.
2. Open terminal
3. Enter the following. Be sure to replace /Replace/With/FilePath.jpeg with your own jpeg file location:
```
curl -X POST http://localhost:8080/image \
-F "file=@/Replace/With/FilePath.jpeg" \
-H "Content-Type: multipart/form-data" > result.png
```

For example:
```
curl -X POST http://localhost:8080/image \
-F "file=@/Users/jblevins/Desktop/testing_image.jpeg" \
-H "Content-Type: multipart/form-data" > result.png
```

4. You should have received a resized PNG and piped it to result.png!

# Development Notes
## Running SQL
1. From commandline, run `mysql -u root -p` to start mysql
2. Create a database by running `create database userprofiles`
3. Change to the database you just created so we can run the sql files: `use userprofiles;`
4. Run the file /user-profiles/Project3/create-tables.sql via `source` with the fully qualifed path name, e.g: `source /Users/jblevins/Development/user-profiles/Project3/create-tables.sql`

## How to run tests
1. `cd` into each directory
2. Run `go test -cover -v`
