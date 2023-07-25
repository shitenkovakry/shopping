create table buyers(
    id serial primary key,
	name varchar(255),
	email varchar(255),
	balance float
);

-- нельзя запускать в sql
insert into buyers(name, email, balance)
values ('Ondrys', 'rianby64@mail.ru', 100.0),
        ('Kry', 'kry@gmail.com', 500.0),
        ('Jenya', 'yurkevich2@mail.ru', 150.5),
        ('Marina','yukmari@mail.ru', 300.5);


create table items(
    id serial primary key,
	name varchar(255),
	price float
);

-- нельзя запускать в sql
insert into items(name, price)
values ('ramen', 50),
       ('cucumber', 25),
       ('cheese', 10),
       ('kimchi', 26);


create table purchases (
    id_purchase serial primary key,
	buyer_id int,
	item_id int,
	foreign key (buyer_id) references buyers(id),
	foreign key (item_id) references items(id)
);

-- нельзя запускать в sql
insert into purchases(buyer_id, item_id)
values (1, 3),
       (2, 4),
       (4, 2),
       (4, 1),
       (3, 3);

alter table purchases add column balance float;

alter table items add column status varchar(255);

alter table items
add constraint chk_status check (status in ('published', 'unpablished'));
