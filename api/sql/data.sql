insert into users (name, nick, email, password)
values
("User 1", "User_1", "usuario1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario1
("User 2", "User_2", "usuario2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"), -- usuario2
("User 3", "User_3", "usuario3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"); -- usuario3

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into publications(title, contenet, author_id)
values
("Publication from User 1", "This is the publication from user 1! Yey!", 1),
("Publication from User 2", "This is the publication from user 2! Yey!", 2),
("Publication from User 3", "This is the publication from user 3! Yey!", 3);