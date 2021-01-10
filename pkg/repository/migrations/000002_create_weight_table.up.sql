CREATE TABLE IF NOT EXISTS weight(
    id serial PRIMARY KEY,
    created_at      timestamp with time zone default now() not null,
    updated_at      timestamp with time zone,
    weight integer not null,
    bmr integer not null,
    daily_caloric_intake integer,
    user_id integer not null,
    FOREIGN KEY (user_id) REFERENCES "user" (id)
);