-- Initial SQL script to initialize the database for Shape API project.

-- For PRODUCTION: need to have database named `shape` CREATED
--   to be able to run this script.

-- For DEVELOPMENT: this script will be executed when database 
--   docker-compose container run for the first time

USE shape;

CREATE TABLE users (
  username VARCHAR(255) NOT NULL PRIMARY KEY,
  password VARCHAR(255) NOT NULL,
  password_salt VARCHAR(255) NOT NULL
)
