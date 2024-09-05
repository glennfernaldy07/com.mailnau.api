-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_role`
(
    id              int             not null primary key AUTO_INCREMENT,
    user_id         bigint             not null,
    role_id         int             not null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_role`;
-- +goose StatementEnd
