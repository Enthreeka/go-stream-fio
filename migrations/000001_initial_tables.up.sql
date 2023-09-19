CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE gen as enum ('male','female');

CREATE TABLE IF NOT EXISTS person(
    id uuid DEFAULT uuid_generate_v4(),
    firstname varchar(100),
    lastname varchar(100),
    age integer,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS gender(
    id int generated always as identity,
    person_id uuid,
    gender gen,
    probability float,
    primary key (id),
    foreign key (person_id)
        references person (id) on delete cascade
);


CREATE TABLE IF NOT EXISTS  address(
    id int generated always as identity,
    person_id uuid,
    country_code varchar(2),
    probability float,
    primary key (id),
    foreign key (person_id)
        references person (id) on delete cascade
);




SELECT (person.id,person.firstname,person.lastname,person.age,gender.gender,gender.probability,address.country_code,address.probability)
    FROM person
    JOIN  address on person.id = address.person_id
    JOIN  gender  on person.id = gender.person_id;

