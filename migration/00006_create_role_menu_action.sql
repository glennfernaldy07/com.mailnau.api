-- +goose Up
-- +goose StatementBegin
CREATE TABLE `role_menu_action`
(
    id              int             not null primary key AUTO_INCREMENT,
    role_id         int             not null,
    menu_id         varchar(50)     not null,
    action_id       varchar(50)             not null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    created_by      varchar(255)    not null default 'SYSTEM',
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_by      varchar(255)    not null default 'SYSTEM',

    foreign key (role_id) references `role`(id),
    foreign key (menu_id) references menu(id),
    foreign key (action_id) references `action`(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `role_menu_action`;
-- +goose StatementEnd
