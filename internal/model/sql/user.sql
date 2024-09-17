CREATE TABLE IF NOT EXISTS users (
    id varchar(255) NOT NULL ,
    name varchar(255) NULL COMMENT 'The username',
    password varchar(255) NOT NULL DEFAULT '' COMMENT 'The password',
    mobile varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
    gender char(10) NOT NULL DEFAULT 'male' COMMENT 'gender,male|female|unknown',
    nickname varchar(255) NOT NULL DEFAULT '' COMMENT 'The nickname',
    role_id varchar(255) NULL COMMENT 'The user role id',
    avatar varchar(255) NULL COMMENT 'The user avatar',
    status tinyint(1) NOT NULL DEFAULT 1 COMMENT 'The user status,0:disable,1:enable',
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "创建时间", 
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT "更新时间",
    deleted_at datetime NULL COMMENT "删除时间" ,
    UNIQUE KEY `uk_mobile` (`mobile`),
    UNIQUE KEY `uk_name` (`name`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB COLLATE utf8mb4_general_ci COMMENT='The user table';