package database

// User table schema
var userTableSchema = `
create table if not exists users (
	id varchar(36) not null,
	email varchar(225) not null unique,
	username varchar(225),
	password varchar(225) not null,
	token varchar(15) not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	primary key (id)
);
`
