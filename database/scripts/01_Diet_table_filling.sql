ALTER SEQUENCE diet_diet_id_seq RESTART WITH 1;

insert into diet
values (0, 'Набор массы 3000 кило', 'a337b13aaa3d71ffb24707d9f73d3f5ad6bcb7388da5a35618965aa0dbd18aab', false);

insert into diet (diet_name, diet_hash, diet_is_personal)
values ('Создана группой', 'f14a528413bc023996568dceaf09295b2b680937d89ab278eaab75551428be52', false);
