-- +goose Up
-- +goose StatementBegin
INSERT INTO `role`( role_name, eselon, opd, created_at, created_by, updated_at, updated_by)
VALUES ( 'admin', '1', '1', current_timestamp, 'SYSTEM', current_timestamp, 'SYSTEM');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE from role where id = 1 and role_name = 'admin';
-- +goose StatementEnd
