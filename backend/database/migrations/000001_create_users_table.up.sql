create table if not exists users(
    id ulid NOT NULL DEFAULT gen_ulid() PRIMARY KEY,
    username text not null,
    password text not null
);