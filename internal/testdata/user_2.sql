CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `email` varchar(100) NOT NULL DEFAULT '',
    `register_time` timestamp NOT NULL,
    `password` varchar(255) NOT NULL DEFAULT '',
    `status` int(10) unsigned NOT NULL DEFAULT 1,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3