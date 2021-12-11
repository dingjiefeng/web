CREATE TABLE `user` (
    `id` INT(11) AUTO_INCREMENT,
    `name` VARCHAR(100) not null,
    `email` VARCHAR(100) not null,
    `created` DATETIME not null,
    `updated` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    primary key (`id`)
) ENGINE=InnoDb DEFAULT CHARSET=utf8;