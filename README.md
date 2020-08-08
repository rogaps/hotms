# hotms (Hotel Management System)

## Entity Relationship Diagram

![Hotel Management System ERD](https://github.com/rogaps/hotms/blob/master/images/hotms-erd.png?raw=true)


## Features

hotms has following features:
- managing multiple hotels
- searching available rooms for date range
- dynamic room price
- reservation for multiple rooms and multiple guests

## Usage

To run hotms you need to run this command:
```
go build ./cmd/hotms
./hotms
 ```
hotms will serve on `localhost:8080`. Postman collection can be found in `postman` folder. hotms' API documentations can be found in the `docs` folder.
