# stations-service
![Build](https://github.com/rso-project-2021/station-service/actions/workflows/build.yml/badge.svg)
![Deploy](https://github.com/rso-project-2021/station-service/actions/workflows/deploy.yml/badge.svg)  
Microservice used for working with filling stations data.

## Environment file
In root of your local repository add `config.json` file.
```
{
    "db_driver" : "postgres",
    "db_source": "postgres://root:secret@localhost:5432/electric_station?sslmode=disable",
    "server_address": "0.0.0.0:8080",
    "gin_mode": "debug"
}
```

## Setup database
1. Run `docker pull postgres:alpine` to download [postgres image](https://hub.docker.com/_/postgres).
2. Run `make postgres` to run postgres image inside of container.
3. Run `make createdb` to create postgres database.
4. Run `make migrateup` to add "users" table.
5. Run `go mod tidy` to clean golang package dependecies.
6. Test project with command `make test`.
7. Run service with `go run .`.
8. Use [PostMan](https://www.postman.com/) to send query to `http://localhost:8080/v1/users/`.

## Seed database
Populate database with some stations. You can run this query in [TablePlus](https://tableplus.com/).
```sql
INSERT INTO stations("name", "lat", "lng", "provider")
VALUES 	('Zidani Most', 45.911460, 14.980200, 'OMV'),
	('Vrhnika', 45.966530, 14.298550, 'Petrol'),
    ('Ljubljanski grad', 46.0451374, 14.4917561, 'Petrol');
```