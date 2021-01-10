CREATE TABLE IF NOT EXISTS "user"(
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    name varchar(255) not null,
    age integer not null,
    height integer not null,
    sex varchar(255) not null,
    activity_level integer not null,
    email varchar unique not null
);