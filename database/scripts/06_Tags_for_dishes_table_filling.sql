ALTER SEQUENCE tags_for_dishes_tag_id_seq RESTART WITH 1;

insert into tags_for_dishes
values (0, 'Завтрак');

insert into tags_for_dishes(tag)
values ('Перекус'),
       ('Обед'),
       ('Ужин'),
       ('Мясное')
       ('Веганское')
       ('Вегетерианское')
       ('Лёгкое');