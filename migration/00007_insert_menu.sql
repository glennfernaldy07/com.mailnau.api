-- +goose Up
-- +goose StatementBegin
INSERT INTO menu (id, menu_name, description, created_at, created_by, updated_at, updated_by) VALUES
('mobile.*', '*', 'Acces All Menu', current_timestamp, 'SYSTEM', current_timestamp, 'SYSTEM'),
('dashboard.*', '*', 'Acces All Menu', current_timestamp, 'SYSTEM', current_timestamp, 'SYSTEM');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
