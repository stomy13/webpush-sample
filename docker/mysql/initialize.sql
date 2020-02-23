CREATE DATABASE webpush default character set 'utf8';
create user 'webpush' identified by 'webpush';
grant all on webpush.* to 'webpush';
USE webpush;

CREATE TABLE IF NOT EXISTS `endpoints` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `endpoint` VARCHAR(2048) NOT NULL,
  `key` VARCHAR(255) NOT NULL,
  `token` VARCHAR(255) NULL,
  PRIMARY KEY (`id`));

COMMIT;