-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `attendance_logs`
(
    id              varchar(128)    not null primary key default(uuid()),
    attendance_id   varchar(128)    not null,
    action          varchar(10)     not null default 'OUT',
    location        varchar(1080)   null,
    image           varchar(1080)   null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',
    constraint unique_id_attendance_id UNIQUE (id, attendance_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `attendance_logs`;
-- +goose StatementEnd
