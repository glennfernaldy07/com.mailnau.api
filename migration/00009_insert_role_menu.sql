-- +goose Up
-- +goose StatementBegin
INSERT INTO role_menu ( role_id, menu_id, created_at, created_by, updated_at, updated_by) VALUES
( 1, 'mobile.*', current_timestamp, 'SYSTEM', current_timestamp, 'SYSTEM'),
( 1, 'dashboard.*', current_timestamp, 'SYSTEM', current_timestamp, 'SYSTEM');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from role_menu where role_id = 1;
-- +goose StatementEnd
