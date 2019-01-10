-- ==================================
-- Crean up old tables
-- ==================================
DROP TABLE IF EXISTS tx_payments;
DROP TABLE IF EXISTS tx_deposits;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS users;

-- ==================================
-- CREATE TABLES:
--   users
--   rooms
--   questions
--   tx_deposits
--   tx_payments
-- ==================================
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `nickname` VARCHAR(255) DEFAULT NULL,
  `mail` VARCHAR(255) DEFAULT NULL,
  `ethaddress` VARCHAR(255) NOT NULL,
  `encrypted_password` VARCHAR(255) DEFAULT NULL,
  `salt` VARCHAR(255) DEFAULT NULL,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `mail_UNIQUE` (`mail` ASC),
  UNIQUE INDEX `ethaddress_UNIQUE` (`ethaddress` ASC)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
AUTO_INCREMENT=1
COMMENT = 'master table of host users';

CREATE TABLE IF NOT EXISTS `rooms` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `owner_id` BIGINT UNSIGNED NOT NULL,
  `owner_address` VARCHAR(255) DEFAULT NULL,
  `title` VARCHAR(255) DEFAULT NULL,
  `description` TEXT DEFAULT NULL,
  `event_code` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) DEFAULT NULL,
  `create_tx_hash` VARCHAR(255) DEFAULT NULL,
  `is_private` TINYINT(1) DEFAULT 0,
  `wei_balance` BIGINT UNSIGNED DEFAULT 0,
  `wei_prize` BIGINT UNSIGNED DEFAULT 0,
  `start_time` DATETIME DEFAULT NULL,
  `end_time` DATETIME DEFAULT NULL,
  `active` TINYINT(1) DEFAULT 1,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `event_code_UNIQUE` (`event_code` ASC),
  CONSTRAINT `fk_rooms_users`
    FOREIGN KEY (`owner_id`)
    REFERENCES `users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COMMENT='master table of rooms';

CREATE TABLE IF NOT EXISTS `questions` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `room_id` BIGINT UNSIGNED NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `owner_id` BIGINT UNSIGNED DEFAULT NULL,
  `body` TEXT DEFAULT NULL,
  `adopted` TINYINT(1) DEFAULT 0,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_questions_rooms`
    FOREIGN KEY (`room_id`)
    REFERENCES `rooms` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COMMENT='master table of questions';

CREATE TABLE IF NOT EXISTS `tx_deposits` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `room_id` BIGINT UNSIGNED NOT NULL,
  `tx_hash` VARCHAR(255) DEFAULT NULL,
  `confirmed` TINYINT(1) DEFAULT 0,
  `success` TINYINT(1) DEFAULT 0,
  `sender` VARCHAR(255) NOT NULL DEFAULT '0x0',
  `receiver` VARCHAR(255) NOT NULL DEFAULT '0x0',
  `wei_amount` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX (`room_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COMMENT='transaction table of deposits';

CREATE TABLE IF NOT EXISTS `tx_payments` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `question_id` BIGINT UNSIGNED NOT NULL,
  `tx_hash` VARCHAR(255) DEFAULT NULL,
  `confirmed` TINYINT(1) DEFAULT 0,
  `success` TINYINT(1) DEFAULT 0,
  `sender` VARCHAR(255) NOT NULL DEFAULT '0x0',
  `receiver` VARCHAR(255) NOT NULL DEFAULT '0x0',
  `wei_amount` BIGINT UNSIGNED NOT NULL DEFAULT 0,
  `updated_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX (`question_id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8
COMMENT='transaction table of payments';

