CREATE TABLE IF NOT EXISTS ports
(
    id serial not null unique,
    name varchar(255) not null,
    isActive boolean null,
    company varchar(255) not null,
    email varchar(255) null,
    phone varchar(255) null,
    address varchar(255) null,
    about varchar(255) null,
    registered date null,
    latitude float null,
    longitude float null
);
