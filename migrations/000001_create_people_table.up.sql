create table people(
    id serial primary key,
    name varchar,
    surname varchar,
    patronymic varchar,
    age integer not null,
    gender varchar,
    nationality varchar
);