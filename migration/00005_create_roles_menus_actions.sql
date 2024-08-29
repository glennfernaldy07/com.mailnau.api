-- +goose Up
-- +goose StatementBegin
CREATE TABLE `roles`
(
    id              bigint                           not null primary key AUTO_INCREMENT,
    name            varchar(50)                      not null,
    echelon         enum('II', 'III', 'IV', 'V')     not null,
    created_at      datetime                         not null default CURRENT_TIMESTAMP,
    updated_at      datetime                         default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (name)
);

CREATE TABLE `menus`
(
    id              bigint          not null primary key AUTO_INCREMENT,
    name            varchar(50)     not null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (name)
);

CREATE TABLE `actions`
(
    id              bigint          not null primary key AUTO_INCREMENT,
    name            varchar(50)     not null,
    created_at      datetime        not null default CURRENT_TIMESTAMP,
    updated_at      datetime        default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (name)
);

CREATE TABLE `role_menu_action`
(
    id          bigint      not null primary key AUTO_INCREMENT,
    role_id     bigint      not null,
    menu_id     bigint      not null,
    action_id   bigint      not null,

    foreign key (role_id) references roles(id),
    foreign key (menu_id) references menus(id),
    foreign key (action_id) references actions(id)
);

-- populate menus
INSERT INTO menus (name) VALUES
('All'),
('Dashboard'),
('User'),
('Aktifitas'),
('Role dan Akses'),
('Master Data'),
('Absensi'),
('Tunjangan Kinerja'),
('Laporan');

-- populate actions
INSERT INTO actions (name) VALUES
('All'),
('Create'),
('Preview'),
('Edit'),
('Delete');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `role_menu_action`;
DROP TABLE IF EXISTS `actions`;
DROP TABLE IF EXISTS `menus`;
DROP TABLE IF EXISTS `roles`;
-- +goose StatementEnd
