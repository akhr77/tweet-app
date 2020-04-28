
-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    username VARCHAR(60) NOT NULL,
    email VARCHAR(100) NOT NULL,
    avater CHAR(100),
    user_profile CHAR(200),
    guest_flag CHAR(1),
    admin_flag CHAR(1),
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS relationships(
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    user_id INT(10) NOT NULL,
    folloer_id INT(10),
    followed_id INT(10),
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS posts(
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    user_id INT(10) NOT NULL,
    image VARCHAR(100) NOT NULL,
    comment VARCHAR(120),
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS likes(
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    user_id INT(10) NOT NULL,
    post_id INT(10) NOT NULL,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS relationships;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS likes;
