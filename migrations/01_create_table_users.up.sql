

create table if not exists users(
    id uuid primary key,
    first_name varchar,
    last_name varchar,
    email varchar,
    age integer,
    created_at timestamp default current_timestamp
);