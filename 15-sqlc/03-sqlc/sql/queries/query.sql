-- name: ListCategories :many
select * from categories;

-- name: GetCategory :one
select * from categories where id = ?;

-- name: CreateCategory :exec
insert into categories (id, name, description) values (?, ?, ?);

-- name: UpdateCategory :exec
update categories set name = ?, description = ? where id = ?;

-- name: DeleteCategory :exec
delete from categories where id = ?;
