# Example Architecture using Golang

In order to show an example for good practices in Golang, I created a simple architecture for a Web API.

## Folder Structure

```
|
pkg/  -> Exported models, it can use as a library/dependency
|
├── models/
|    ├── requests/   -> Exported Request models
|    └── responses/  -> Exported Response models
|
internal/  -----------> Application code / it cannot use as a library/dependency
|
├── app/     -> Application folder
|    ├── data/          -> Repositories / Data Access
|    ├── controllers/   -> HTTP Handlers
|    ├── helpers/       -> functions helpers for application
|    |    ├── constants/constants.go --> constants of application
|    |    ├── ioc/dependencies.go ----> configuration all dependencies for the application
|    |    ├── cache.go  --------------> configuration for cache
|    |    ├── errors.go --------------> all errors of application
|    |    └── log.go  ----------------> configuration for log object
|    ├── middlewares/   -> middlewares
|    ├── models/        -> internal models/entities
|    ├── routes/        -> routes of app
|    ├── services/      -> bussiness logic layer / services of application
|    └── startup.go     -> Where all function for setup are called
|
├── core/    -> Frameworks & Drivers
|    ├── ioc/          -> Inversion of control: Dependency Injection provider
|    ├── server/       -> HTTP Providers (Gin, Echo, etc)
|    ├── database/     -> Databases Providers (Gorm, etc)
|    ├── utils/        -> shared functions
|
...
```

understand the folder structure:

- everything in app folder can see core folder
- core folder cannot see app folder
- pkg contains all shared models

## Data Flow

```
|
| REQUEST
├──> Echo/Gin/WebHandler
|    └──> Middlewares
|        └──> Http Handler
|    	    └──> Service
|    	    	└──> Repository
|    	    	└──> Clients
```

## How to run

```bash
go run main.go
```
