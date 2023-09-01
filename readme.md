<h1 align="center">Self payroll System</h1>
The main feature of this web service is employees can do salary withdrawals independently every month.

## Getting Started

### Prerequisites
- Go 1.18+
- PostgreSQL 14+

### Installation or Configure
```bash
# copy the env variable template and adjust your env configuration
$ cp .env.example .env

# install the dependencies
$ go mod tidy && go mod vendor

# run server
$ go run main.go
```

The list of endpoints is available in the [documenter](https://documenter.getpostman.com/view/4080490/2s83Ychhk4).