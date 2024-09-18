-- +goose Up
-- +goose StatementBegin
INSERT INTO users ( nik, email, password, status, created_at, created_by, updated_at, updated_by, role_id)
VALUES ( '1', 'admin@admin.com', '$2a$14$AzpCbGw9YjeQwdaQSxJd8OQZW42pT9yTcu2XlQ36MG0oyvJYOsCqS', '', '2024-09-05 11:52:02', '', '2024-09-05 11:52:02', '', 1);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users where id = 1 and nik = 1;
-- +goose StatementEnd
