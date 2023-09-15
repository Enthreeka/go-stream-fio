CREATE TYPE gen as enum ('male','female');

CREATE TABLE IF NOT EXISTS person(
    id int generated always as identity,
    firstname varchar(100),
    lastname varchar(100),
    age integer,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS gender(
    id int generated always as identity,
    person_id int,
    gender gen,
    probability float,
    primary key (id),
    foreign key (person_id)
        references person (id) on delete cascade
);


CREATE TABLE IF NOT EXISTS  address(
    id int generated always as identity,
    person_id int,
    country_code varchar(2),
    probability float,
    primary key (id),
    foreign key (person_id)
        references person (id) on delete cascade
);