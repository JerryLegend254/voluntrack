CREATE TABLE `opportunities` (
  `opportunity_id` INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `description` varchar(255) NOT NULL,
  `organization_email` varchar(255) NOT NULL,
  `status` integer NOT NULL
);
