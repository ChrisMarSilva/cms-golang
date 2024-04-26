
CREATE DATABASE IF NOT EXISTS qasys DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

create table qasys.users
(
    id         bigint unsigned auto_increment primary key,
    username   varchar(255) not null,
    password   varchar(255) not null,
    email      varchar(255) not null,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    constraint email unique (email),
    constraint username unique (username)
);



create table qasys.questions
(
    id         bigint unsigned auto_increment primary key,
    user_id    int          not null,
    title      varchar(255) not null,
    content    text         not null,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null
);

create index questions_user_id_index on qasys.questions (user_id);


create table qasys.answers
(
    id          bigint unsigned auto_increment primary key,
    question_id int      not null,
    user_id     int      not null,
    content     text     not null,
    created_at  datetime null,
    updated_at  datetime null,
    deleted_at  datetime null
);

create index answers_question_id_index on qasys.answers (question_id);
create index answers_user_id_index on qasys.answers (user_id);