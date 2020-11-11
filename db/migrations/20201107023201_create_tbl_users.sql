
-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `login_id` VARCHAR(255) NOT NULL COMMENT '',
  `password` VARCHAR(60) NOT NULL COMMENT '',
  `remember_token`  VARCHAR(100),
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`),
  UNIQUE(`login_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'ユーザー管理';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd