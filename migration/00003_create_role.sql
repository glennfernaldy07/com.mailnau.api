-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `role`
(
    id              int          not null primary key AUTO_INCREMENT,
    role_name       varchar(255)    not null,
    eselon          varchar(50)     not null,
    opd             varchar(50)     not null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',
    constraint unique_role_name_eselon_opd UNIQUE (role_name, eselon, opd)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `role`;
-- +goose StatementEnd
