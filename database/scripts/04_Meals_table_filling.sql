ALTER SEQUENCE meals_meal_id_seq RESTART WITH 1;

insert into meals
values (0, 'Завтрак качка', 'd154c7edfc68899a46e017394f5427f3d4a8a5c659570e1e47be9dbdf7e4d1c6', null);

insert into meals (meal_name, meal_hash, tag_id)
values ('Перекус качка', '22157cc64dafa7000bcfe91d280818e3821b6b376aa6a9c65e9de85eadb4c047', null),
       ('Обед качка', 'ce5c5d66381d5c636f1b6be0bc417ec4e5c2c997cba4cc60810ba36aa6ad9623', null),
       ('Ужин качка', '5427a2e2b9105a6f727857f7c7cb5838a18bfb887abda8ee4b09f8c06d2ca793', null),
       ('Завтрак перед пробежкой', '5a2dd78474ab98c9f122df158edd08fe23052aa658be654cb7ae715e24d1814d', null),
       ('Перекус перед качалкой', '6f950e170171c8ee587a4891f6ec8af8e1e1723365203e6f4d12b8735514d176', null),
       ('Обед для хорошего сна', 'd8f3ae36061b4dade40e75c94573b2ca4d1315e491b2ea50eac5ce0c342f0e98', null),
       ('Ужин перед взрывной тренировкой', '00431e39782111b01e6d5a4b7dd5a020c97de7e845eea760377613c1c27b89f8', null),
       ('Завтрак "Cheat meal"', '8e866eb4a7ec7da22076f859aa87c584bc9f0fdec932ea3dc90a522bcc6a420b', null),
       ('Перекус "Cheat meal"', 'e489c83030c43c2a954122a097a55f2fe200f4ec633bef0a1ae0ea72f95d04ac', null),
       ( 'Обед "Cheat meal"', 'a7174c9ded91c375742daf7eaadfff524883158662018571f36fb136ca1af2da', null),
       ( 'Ужин "Cheat meal"', '57408602f767b565e6f2b94e9610caeaa58fa15a53e74d62dc958f8e72ced6c6', null);
