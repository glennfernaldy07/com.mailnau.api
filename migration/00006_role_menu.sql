-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `role_menu`
(
    id              int             not null primary key AUTO_INCREMENT,
    role_id         int             not null,
    menu_id         varchar(50)      not null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    FOREIGN KEY (menu_id) REFERENCES menu(id) ON DELETE CASCADE,
    UNIQUE KEY role_menu_unique (role_id, menu_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `role_menu`;
-- +goose StatementEnd
