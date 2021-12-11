CREATE TABLE `user` (
    `id` INT(11) AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(100) NOT NULL,
    `password` TEXT NOT NULL,
    `created` DATETIME NOT NULL ,
    `updated` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    primary key (`id`)
) ENGINE=InnoDb DEFAULT CHARSET=utf8;