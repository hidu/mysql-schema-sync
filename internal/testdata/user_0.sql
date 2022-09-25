CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `email` varchar(1000) NOT NULL DEFAULT '',
    `register_time` timestamp NOT NULL,
    `password` varchar(1000) NOT NULL DEFAULT '',
    `status` tinyint unsigned NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3