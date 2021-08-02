create table user
(
    id          bigint auto_increment               primary key,
    user_id     bigint                              not null,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         null,
    gender      tinyint   default 0                 not null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null,
    constraint user_user_id_uindex
        unique (user_id),
    constraint user_username_uindex
        unique (username)
);


