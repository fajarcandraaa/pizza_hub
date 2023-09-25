# pizza_hub
<!-- ABOUT THE PROJECT -->
## About The Project

This is a simple app to implement time.sleep and waitgroup to handle concurency in golang, with back story is Pizza Hub :smile:

### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Golang](https://golang.com)
* [PostgreSQL](https://www.postgresql.org/)

<!-- GETTING STARTED -->
## Getting Started
Before we get started, it's important to know that that this code use a custom command to execute it with makefile to make more simple command like :
1. make update
2. make tidy
3. make start

So, let start it.
1. After clone this repository, just run `make update`.
2. Setup your `.env` file such as database connection and redis connection based on default setting on you device.
3. To make sure that all dependency is run well, than run `make tidy`.
4. Finally, you can run your project with command: `make start`.
5. Go to postman and set url like `http://localhost:8080/`, for information that port to run this project depend on configuratin on `.env`

And for additional information, i'm alredy setup unit-testing, just run `make test-service`.

## Documentation
Here is more information on the process flow and API Contract documentation:

[POSTMAN API COLLECTION AND API CONTRACT](https://drive.google.com/file/d/1UeCsU7lZR-a8pOspmluKJMISgynhh2oQ/view?usp=sharing)

## Afterword
Hopefully, it can be easily understood and useful. Thank you~
