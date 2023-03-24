# ZeroGO-API

A backend service for my internal project project with Gin, MongoDB, JWT Authentication Middleware, Test, and Docker inspired by https://github.com/amitshekhariitbhu/go-backend-clean-architecture

I have improved some feature such as : 
- Updated using Go 1.20 and newer module dependencies.
- Add live reloading while on development, no neeed to rebuild the image container.
- Add default JSON structure for http response.
  ```json
  {
    "code" : 200, // or 400,500
    "status" : "OK", // or BAD REQUEST, INTERNAL SERVER ERROR
    "data" : {...}   // Object for single or array for multiple
  }    
  ```
- Integrated with OpenAI API.

## Architecture Layers of the project

- Router
- Controller
- Usecase
- Repository
- Domain



## Major Packages used in this project

- **gin**: Gin is an HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need a smashing performance, get yourself some Gin.
- **mongo go driver**: The Official Golang driver for MongoDB.
- **jwt**: JSON Web Tokens are an open, industry-standard RFC 7519 method for representing claims securely between two parties. Used for Access Token and Refresh Token.
- **viper**: For loading configuration from the `.env` file. Go configuration with fangs. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.
- **mockery**: A mock code autogenerator for Golang used in testing.
- Check more packages in `go.mod`.


#### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.
- Access API using `http://localhost:8080`


## CI/CD
- using jenkins poll SCM.
