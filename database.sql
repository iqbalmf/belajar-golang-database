create table customer
(
    id   varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) engine InnoDb;
select *
from customer;
desc customer;

delete
from customer;

alter table customer
    add column email      varchar(100),
    add column balance    int       default 0,
    add column rating     double    default 0.0,
    add column created_at timestamp default current_timestamp,
    add column birth_date date,
    add column married    boolean   default false;

insert into customer(id, name, email, balance, rating, birth_date, married)
values ('iqbal', 'Iqbal', 'test@test.com', 100000, 98.9, '2024-03-11', false);
insert into customer(id, name, email, balance, rating, birth_date, married)
values ('fauzan', 'fauzan', 'fauzan@test.com', 123321, 89.9, '2024-03-10', true);
select id, name, email, balance, rating, birth_date, married, created_at from customer;

insert into customer(id, name, email, balance, rating, birth_date, married)
values ('nulldata', 'null', null, null, null, null, null);

create table user
(
    username varchar(100) not null ,
    password varchar(100) not null ,
    primary key (username)
) engine InnoDb;
select * from user;
desc user;
insert into user(username, password) values ('admin', 'admin')

create table comments
(
    id int auto_increment not null ,
    email varchar(100) not null ,
    comment text,
    primary key (id)
) engine InnoDb;
desc comments;
select * from comments;