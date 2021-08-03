create table if not exists post
(
    id bigint auto_increment primary key,
    post_id bigint not null,
    title varchar(128) not null,
    content varchar(10000) not null,
    user_id bigint not null,
    community_id bigint not null,
    status tinyint default 1 not null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null
);

