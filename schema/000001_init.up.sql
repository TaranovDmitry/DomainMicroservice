CREATE TABLE IF NOT EXISTS ports
(
    id integer not null unique,
    name varchar not null,
    isActive boolean null,
    company varchar not null,
    email varchar null,
    phone varchar null,
    address varchar null,
    about  varchar null,
    registered date null,
    latitude float null,
    longitude float null
);
