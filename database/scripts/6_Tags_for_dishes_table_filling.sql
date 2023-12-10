ALTER SEQUENCE tags_for_dishes_tag_id_seq RESTART WITH 1;

insert into tags_for_dishes
values (0, 'breakfast');

insert into tags_for_dishes(tag)
values ('breakfast'),
       ('snack'),
       ('lunch'),
       ('dinner');