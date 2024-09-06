CREATE DATABASE IF NOT EXISTS voluntrack;

CREATE TABLE `volunteers` (
  `email` varchar(255) NOT NULL PRIMARY KEY,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `available` boolean NOT NULL DEFAULT true,
  `password` varchar(60) NOT NULL
);

CREATE TABLE `skills` (
  `skill_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(50) UNIQUE NOT NULL
);

CREATE TABLE `organizations` (
  `name` varchar(255) UNIQUE NOT NULL,
  `description` varchar(255) NOT NULL,
  `email` varchar(255)  PRIMARY KEY NOT NULL,
  `password` varchar(60) NOT NULL
);

CREATE TABLE `opportunities` (
  `opportunity_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `organization_email` varchar(255) NOT NULL,
  `status` integer NOT NULL
);

CREATE TABLE `opportunities_status` (
  `ost_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `status` integer
);

CREATE TABLE `opportunity_skills` (
  `osk_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `opportunity_id` integer NOT NULL,
  `skill_id` integer NOT NULL
);

CREATE TABLE `volunteer_skill` (
  `vsk_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `skill_id` integer NOT NULL
);

CREATE TABLE `volunteer_opportunity` (
  `vo_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `opportunity_id` integer NOT NULL
);

ALTER TABLE `volunteer_skill` ADD FOREIGN KEY (`email`) REFERENCES `volunteers` (`email`);

ALTER TABLE `volunteer_skill` ADD FOREIGN KEY (`skill_id`) REFERENCES `skills` (`skill_id`);

ALTER TABLE `volunteer_opportunity` ADD FOREIGN KEY (`email`) REFERENCES `volunteers` (`email`);

ALTER TABLE `opportunities` ADD FOREIGN KEY (`organization_email`) REFERENCES `organizations` (`email`);

ALTER TABLE `opportunities` ADD FOREIGN KEY (`status`) REFERENCES `opportunities_status` (`ost_id`);

ALTER TABLE `volunteer_opportunity` ADD FOREIGN KEY (`opportunity_id`) REFERENCES `opportunities` (`opportunity_id`);

ALTER TABLE `opportunity_skills` ADD FOREIGN KEY (`opportunity_id`) REFERENCES `opportunities` (`opportunity_id`);

ALTER TABLE `opportunity_skills` ADD FOREIGN KEY (`skill_id`) REFERENCES `skills` (`skill_id`);
