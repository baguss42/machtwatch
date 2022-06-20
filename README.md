## Description
This repository is task assignment from jamtangan.com for Software Engineer - Back End Test.
The project is using repository pattern and use logging on middleware level to log each request

## Assumptions
* There is carts that will pass to make order
* We have price reduction in products to store discount
* Each transaction details save product histories
* Better run this application on linux environment

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

## Note
Thank you for reviewing my task, I realize there are still many things that need to be improved. Because I am only doing this task only 2 day (weekend), 

If you have feedback, please notice me by email me at bagussadewo42@gmail.com.

Cheers, Thank you