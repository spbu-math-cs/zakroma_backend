ALTER SEQUENCE tags_for_meals_tag_id_seq RESTART WITH 1;

insert into tags_for_meals(tag_id, tag) values (0, 'завтрак');
insert into tags_for_meals(tag) values
    ('обед'), ('ужин'), ('перекус');