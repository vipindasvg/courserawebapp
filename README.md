# courserawebapp

I have used postgresql as database

Database Configuration:

Create database user:
create user cloudstick with password 'coursera';

Create database:
create database cloudstick with owner coursera;

Run the application:

cd courserawebapp
go mod tidy
go build
./courserawebapp