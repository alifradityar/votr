CREATE SCHEMA `votr` DEFAULT CHARACTER SET utf8 ;
CREATE TABLE topic (
    `id` CHAR(36) CHARACTER SET utf8 NOT NULL,
    `title` VARCHAR(100) NOT NULL,
    `upvote` INT(10) NOT NULL,
    `downvote` INT(10) NOT NULL,
    `score` INT(10) NOT NULL,
    `created` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id_unique` (`id`),
    INDEX `score_index` (`score`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8;