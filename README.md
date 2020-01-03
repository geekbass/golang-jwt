# Golang JWT Course
Playing around with Golang, APIs, JWT with Postgres Backend. This is from the Udemy course "[Golang: Intro to JWT Auth with Golang](https://www.udemy.com/course/build-jwt-authenticated-restful-apis-with-golang/)" by Mike M. All credit goes out to Mike for this awesome course!

This course is really good for any beginner Golang developer looking to learn how to create APIs with a DB backend. Try it out for yourself!

# Setup
NOTE: The course uses [ElephantSQL](https://www.elephantsql.com/) but this uses Docker locally. 

Using Docker setup PGSQL:
```bash
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```

Create the Table and insert some stuff:
```bash
create table users ( id serial primary key, email text not null unique, password text not null);

insert into users (email, password) values ('test2@example.com', 'abcd');

insert into users (email, password) values ('test@example.com', '12345');

select * from users;
```

Setup the Golang packages:
```bash
go get -u github.com/gorilla/mux
go get github.com/dgrijalva/jwt-go
go get -h
go get -u github.com/lib/pq
go get -u github.com/davecgh/go-spew/spew
go get -u golang.org/x/crypto/bcrypt
```

Running the program:
```bash
go run main.go
```