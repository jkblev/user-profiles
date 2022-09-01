DROP TABLE IF EXISTS userpassword;

DROP TABLE IF EXISTS user;
CREATE TABLE user (
                       id         INT AUTO_INCREMENT NOT NULL,
                       firstname      VARCHAR(128) NOT NULL,
                       lastname     VARCHAR(128) NOT NULL,
                       zipcode      INT(5) NOT NULL,
                       PRIMARY KEY (`id`)
);

CREATE TABLE userpassword (
    id INT AUTO_INCREMENT NOT NULL,
    hashedpassword VARCHAR(255) NOT NULL,
    lastchanged DATETIME DEFAULT NULL,
    active BOOLEAN NOT NULL,
    userid INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (userid)
                          REFERENCES user(id)
                          ON DELETE CASCADE
);