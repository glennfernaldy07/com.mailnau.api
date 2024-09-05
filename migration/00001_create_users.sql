-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users`
(
    id              bigint          not null primary key AUTO_INCREMENT,
    nik             varchar(50)     not null,
    email           varchar(255)   not null,
    password        varchar(255)    not null,
    status          varchar(20)         not null default 'ACTIVE',
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)        not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)        not null default 'SYSTEM',
    constraint unique_email UNIQUE (email),
    constraint unique_nik UNIQUE (nik)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `users`;
-- +goose StatementEnd
