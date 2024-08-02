-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_detail`
(
    id              bigint          not null primary key AUTO_INCREMENT,
    task_id         bigint          not null,
    task_name       varchar(255)    not null,
    description     varchar(1000)   not null,
    priority        tinyint         not null default 1,
    is_done         tinyint         not null default 0,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    constraint task_unique_title UNIQUE (task_id, task_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `user_detail`;
-- +goose StatementEnd
