# MY GOLANG BACKEND CHALLENGE
## About
In this challenge I focused on learning about parsing CSV files into an API format. 

## Built using
* [github.com/gorilla/mux](github.com/gorilla/mux) for the routes
* [github.com/nu7hatch/gouuid](github.com/nu7hatch/gouuid) for generating the UUIDs

## Installation and Tests
In order to host the API:
1. Clone the repository
2. `go test -v` to run tests
3. `go build`
4. `./gochallenge`

Then visit: 
[http://localhost:8000/customers](http://localhost:8000/customers)

***

## API endpoints

### GET /
`returns some information about me`
### GET /customers
`returns all customers `
### POST /customers
format:
```json
{
    "name": "majid",
    "email": "email@email.com",
    "text": "creationtime and ID are handled through some imported packages"
}
```
_any additional fields will be rejected_
***

## Thoughts
Although this was not my first time dealing with APIs, it was my first proper project using Go. Working with CSV file was a challenge that I overcame with customer_loader and I  found writing tests to be much easier than in other languages. Importing go packages is also as easy as pointing to the git repository. Overall, it is very intuitive and quick to get used to, and I can't wait to learn more.

***

## Potential to-dos and learning points

* Set up authentication for get requests API
* Sort API response 
* Learn how to structure a go project with imports, etc.