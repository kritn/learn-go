-- -----------------
-- CREATE DATABASE
-- ------------------
-- CREATE DATABASE IF NOT EXISTS dpost CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE clean_arc;

-- ------------
-- CreateTABLE
-- -------------

CREATE TABLE users(
	id int(6) PRIMARY KEY  NOT NULL AUTO_INCREMENT,
	username varchar(50),	
	password varchar(100),
	createdatetime datetime DEFAULT current_timestamp()	
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE items(
	id int(6) PRIMARY KEY  NOT NULL AUTO_INCREMENT,
	item varchar(50),	
	detail varchar(255),
	createdatetime datetime DEFAULT current_timestamp()	
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
