# SETUP PROJECT EXAMPLE

Go to `/authentication` folder and check the setup project example.

Implementation setps:

**1:**

- Create project structure folder according clean architechture design pattern;
  - **Domain** - Responsible for structure data models and dto's(request and response)
    - INTERACTS WITH REPOSITORY, USECASE and CONTROLLER
  - **Repository** - Responsible for doing operation in database
    - INTERACTS WITH USECASE
  - **UseCase** - Responsible for business logic
    - INTERACTS WITH CONTROLLER
  - **Controller** - Responsible to route all app and give implementation to handler methods
    - INTERACTS WITH USECASE
  - **EXTRA:**
    - **Initializers** - Responsible to set db, set http server, load env variables

**2:**

- Setup environment loader to load env variables - `/authentication/initializers/environment`
  - `github.com/joho/godotenv` package

**3:**

- Setup http server to handle with http requests and responses - `/authentication/initializers/http`
  - `echo` framework

**4:**

- Setup database connection - `/authentication/initializers/db`
  - `gorm` orm library (postgres)

**5:**

- Setup `init()` and `main()` functions to load environment variables, start database connection, start http server - `/authentication/main.go`
