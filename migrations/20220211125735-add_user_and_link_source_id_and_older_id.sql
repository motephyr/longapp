
-- +migrate Up

Insert into users (id, username, password)
values (1, 'aaa', '$argon2id$v=19$m=65536,t=1,p=2$fy9iFofpebQXSyPTOSYBMQ$/aZGGXySjzOiNX6z0tmXv2jJ+MAmcYdM5ayVQtWCjN4');

INSERT INTO user_idstrings (user_id, idstring)
SELECT DISTINCT 1, idstring From sources;

INSERT INTO user_olders (user_id, older_id)
SELECT DISTINCT 1, id From olders;
-- +migrate Down

TRUNCATE TABLE user_olders;
TRUNCATE TABLE user_idstrings;
delete from  users where id = 1;

