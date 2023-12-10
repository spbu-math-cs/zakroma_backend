ALTER SEQUENCE diet_day_diet_day_id_seq RESTART WITH 1;

insert into diet_day
values (0, 'Жесточайший набор массы 1');

insert into diet_day (diet_day_name)
values ('Жесточайший набор массы 1'),
       ('Жесточайший набор массы 2'),
       ('Жесточайший набор массы 3'),
       ('Жесточайший набор массы 4'),
       ('Жесточайший набор массы 5'),
       ('Жесточайший набор массы 6'),
       ('Жесточайший набор массы 7');
