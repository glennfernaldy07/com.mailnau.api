-- +goose Up
-- +goose StatementBegin
INSERT INTO role_menu_action (role_id, menu_id, action_id) VALUES
( 1, 'mobile.*', 'action.*'),
( 1, 'dashboard.*', 'action.*');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from role_menu_action;
-- +goose StatementEnd
