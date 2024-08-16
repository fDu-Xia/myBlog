create table posts
(
    id         bigint unsigned auto_increment
        primary key,
    username   varchar(255) not null,
    post_id    varchar(256) not null,
    content    longtext     not null,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    constraint post_id_key
        unique (post_id)
)
    collate = utf8mb4_bin;

create index post_username_index
    on posts (username);

create table users
(
    id         bigint unsigned auto_increment
        primary key,
    username   varchar(255) not null,
    password   varchar(255) not null,
    nickname   varchar(30)  not null,
    email      varchar(256) not null,
    created_at datetime     null,
    updated_at datetime     null,
    deleted_at datetime     null,
    constraint users_pk_2
        unique (username)
)
    collate = utf8mb4_bin;


