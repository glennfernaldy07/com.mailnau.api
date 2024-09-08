-- +goose Up
-- +goose StatementBegin
ALTER TABLE `users`
ADD COLUMN role_id int,
ADD CONSTRAINT fk_users_role_id FOREIGN KEY (role_id)
REFERENCES `role`(id)
ON DELETE SET NULL
ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP FOREIGN KEY fk_users_role_id;
-- +goose StatementEnd
