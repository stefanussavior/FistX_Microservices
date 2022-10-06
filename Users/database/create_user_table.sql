-- +migrate Up
CREATE TABLE users (
   id int(10) unsigned NOT NULL AUTO_INCREMENT,
   id_role int(10) NOT NULL,
   id_province int(10) NOT NULL,
   id_city int(10) NOT NULL,
   username varchar(191) NOT NULL,
   password varchar(191) NOT NULL,
   email varchar(191) UNIQUE NOT NULL,
   created_at timestamp NULL DEFAULT NULL,
   updated_at timestamp NULL DEFAULT NULL,
   deleted_at timestamp NULL DEFAULT NULL,
   PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down
DROP TABLE users;

