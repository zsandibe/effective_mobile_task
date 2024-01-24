# Effective Mobile Test Task


## Clone the project

```
$ git clone https://github.com/zsandibe/effective_mobile_task
$ cd effective_mobile_task
```

## Launch a project

```
$ make run
```

## Execute migrations

```
$ make migrate-up
$ make migrate-down
```


## API server provides the following endpoints:
* `GET /person` - returns a list of persons (parameters can be used: gender, nationality, page, size)
* `GET /person/:id` - returns person by id
* `POST /person/add` - adds a person

```
{
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich" <-optionally
}
```

* `PUT /person/update/:id` - updates person by id

```
{
    "id": 1,
    "name": "Genadi",
    "surname": "Shapkin",
    "patronymic": "Petrovich", <-optionally
    "age": 21,
    "gender":"male",
    "nationality":"RU"
}
```

* `DELETE /person/delete/:id` - deletes person by id

# .env file
## Enrichment configuration

```
AGE_API="https://api.agify.io"
GENDER_API="https://api.genderize.io"
NATIONALITY_API="https://api.nationalize.io"
```

## Server configuration

```
SERVER_HOST=localhost
SERVER_PORT=7777
```

## Postgres configuration

```
DRIVER="postgres"
DB_USER="postgres"
DB_PASSWORD="test"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="effective_mobile"
```

