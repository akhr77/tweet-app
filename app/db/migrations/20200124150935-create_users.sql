
-- +migrate Up
CREATE TABLE IF NOT EXISTS tweets(
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    tweet CHAR(30),
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE IF EXISTS tweets;
