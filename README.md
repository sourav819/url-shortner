# URL- Shortner
url shortner service

### Configuration Management
#### ENV Variables
required env variables have been listed in `.env.example`

#### Server Configuration
- Used [Gin](https://github.com/gin-gonic/gin) Web Framework

#### Database Configuration
- Using [GORM](https://github.com/go-gorm/gorm) as an ORM.
- Using postgresSQL for relational database

#### Local Setup Instruction
Follow these steps:
- Copy [.env.example](.env.example) as .env and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go`

### Boilerplate Structure

<pre>├── <font color="#3465A4"><b>controllers</b></font>
├─── <font color="#3465A4"><b>models</b></font>
├─── <font color="#3465A4"><b>objects</b></font> 
├─── <font color="#3465A4"><b>pkg</b></font> 
│   ├───config
│   ├───database
│   ├───helper
│   └───logger
├─── <font color="#3465A4"><b>routers</b></font> 
└─── <font color="#3465A4"><b>utils</b></font>
</pre>
# Future TODO list

- [] Build frontend with React.js
- [] Feature: Count of clicks on short url
- [] feature: Caching with Redis