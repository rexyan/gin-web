create table if not exists community
(
    id int auto_increment primary key,
    community_id int not null,
    community_name varchar(128) not null,
    introduction varchar(256) not null,
    create_time timestamp default CURRENT_TIMESTAMP not null,
    update_time timestamp default CURRENT_TIMESTAMP null
);