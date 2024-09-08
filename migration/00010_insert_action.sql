-- +goose Up
-- +goose StatementBegin
INSERT INTO `action` (id, action_name) VALUES
('action.*', 'All'),
('action.create', 'Create'),
('action.preview', 'Preview'),
('action.edit', 'Edit'),
('action.delete', 'Delete');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `action`;
-- +goose StatementEnd
