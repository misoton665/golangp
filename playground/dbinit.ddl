CREATE TABLE users (id SERIAL, user_id varchar(32) NOT NULL UNIQUE, name varchar(64) NOT NULL, email varchar(64) NOT NULL, create_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE messages (id serial PRIMARY KEY, user_id varchar(32) NOT NULL REFERENCES users(user_id), body varchar(140) NOT NULL, create_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP);
