-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `menu`
(
    id              varchar(50)     not null primary key,
    menu_name       varchar(255)    not null,
    description     varchar(255)    null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `menu`;
-- +goose StatementEnd
