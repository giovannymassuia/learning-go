create table category (
    id string primary key,
    name string,
    description string
);

create table course (
    id string primary key,
    name string,
    description string,
    category_id string references category(id)
);