CREATE TABLE users
(
    id              serial not null unique,
    username        varchar(255) not null unique, 
    password_hash   varchar(255) not null,
    role            varchar(255) not null
);

CREATE TABLE actors
(
    id          serial not null unique,
    name        varchar(128) not null unique, 
    gender      varchar(32) not null,
    birth_date  date not null
);

CREATE TABLE films
(
    id              serial not null unique,
    name            varchar(150) not null unique, 
    description     varchar(1000) not null,
    release_date    date not null,
    rating          smallint not null
);

CREATE TABLE actors_films
(
    id          serial not null unique,
    actor_id    int references actors(id) on delete cascade not null, 
    film_id     int references films(id) on delete cascade not null
);