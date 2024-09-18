-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_detail`
(
    id              varchar(128)          not null primary key default(uuid()),
    user_id         varchar(128)          not null,
    name            varchar(255)    not null,
    phone           varchar(25)     null,
    address         varchar(255)    null default 1,
    image           varchar(100)    null default 0,
    finger_print    varchar(1000)   null,
    barcode         varchar(1000)   null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_detail`;
-- +goose StatementEnd
