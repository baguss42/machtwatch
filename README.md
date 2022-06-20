## Prerequisites

* Golang 1.16 or higher
* Mysql

## Installation
If you are using linux and docker installed, just simply run this command:
```
make deploy
```

If you don't have docker installed:
* Configure your own mysql database
* Execute schema in `database/schema/schema.sql`
```
cp env .env
make run
```
Then application will run on localhost:8080

If you want to run testing run this command:
```
make test
```