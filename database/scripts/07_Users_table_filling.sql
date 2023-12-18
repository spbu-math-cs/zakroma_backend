ALTER SEQUENCE users_user_id_seq RESTART WITH 1;

insert into users values
        (0, 'admin', 'adminov', 'YAVSEHNENAVIZHY@yandex.com', '$2a$10$YDFE2R4tVHf1gG7qSF29kuYLxZ1sk4OeWZQb7XFa0zWEDVab868QK', '1999-12-31', '55f0d301f35c0157e25c665d9169abf99186d84756061d5a87a4e8bdcd0e7138');

insert into users (user_name, user_surname, user_email, password_hash, birth_date, user_hash)
values
       ('Gigachad', 'AAAAAAA', 'eeeeeeeeeee@gg.gg', 'ef92b7784aee0b8a8a7e4c144be76b7c72f9e436a6c5761a844e7d5e454d29c3', '2007-02-01', '43cb59219f2c150a5c9cf84eafadf1fc2262cbcd1ec13ddf43ce653da9339784');