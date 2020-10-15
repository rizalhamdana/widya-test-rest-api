# widya-test-rest-api
This is a Golang Code project that is a test for joining widya.ai company


### How to Run


#### Prerequisite
1. Go (Programming Language) version >= 1.14.4 [Install Here](https://golang.org/doc/install#install)
2. PostgreSQL (Database) version >= 11 [Install Here](https://www.postgresql.org/download/)
3. Postman (For testing API URL) [Install Here](https://www.postman.com/downloads/)


#### Step-by-Step

1. Change Following Variables in the `.env` file with your own database credentials.
```
READ_DB_USER_POSTGRES=postgres
READ_DB_PASSWORD_POSTGRES=
READ_DB_HOST_POSTGRES=localhost
READ_DB_PORT_POSTGRES=5432
READ_DB_NAME_POSTGRES=postgres
```
2. Import the `widya-test.sql` file into the public schemas in your database.
3. Run `go run .` inside this project root address with your terminal/commad prompt.
4. Import the `WIDYA-TEST.postman_collection` file into the Postman and send all the requests that I have provided.
5. If there is any issue, Just write down on the Issues section of this repository
