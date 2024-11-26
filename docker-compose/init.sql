create table users
(
    id       uuid not null
        primary key,
    login    text
        unique,
    password text
);

alter table users
    owner to postgres;

create table img
(
    id   uuid not null
        primary key,
    data bytea
);

alter table img
    owner to postgres;

create table projects
(
    id       uuid not null
        primary key,
    owner_id uuid
        references users
            on update cascade on delete cascade,
    title    text
);

alter table projects
    owner to postgres;

create table pages
(
    id         uuid not null
        primary key,
    owner_id   uuid
        references users
            on update cascade on delete cascade,
    title      text,
    data       jsonb,
    project_id uuid
        references projects
            on update cascade on delete cascade
);

alter table pages
    owner to postgres;

create table participants
(
    user_id uuid
        references users,
    page_id uuid
        references pages
);

alter table participants
    owner to postgres;



