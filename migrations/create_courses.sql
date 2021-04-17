-- +migrate Up

create table if not exists courses
(
    id              serial primary key,
    cid             text, 
    course_type     text,
    slug            text,
    name            text
);

-- +migrate Down
DROP TABLE IF EXISTS courses;