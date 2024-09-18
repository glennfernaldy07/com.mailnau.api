-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `attendance`
(
    id              varchar(128)    not null primary key default(uuid()),
    user_id         varchar(128)    not null,
    status          varchar(20)     not null default 'OUT',
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',
    constraint unique_id_user_id UNIQUE (id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `attendance`;
-- +goose StatementEnd
