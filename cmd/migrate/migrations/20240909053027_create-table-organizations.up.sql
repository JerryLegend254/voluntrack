CREATE TABLE `organizations` (
  `name` varchar(255) UNIQUE NOT NULL,
  `description` varchar(255) NOT NULL,
  `email` varchar(255)  PRIMARY KEY NOT NULL,
  `password` varchar(60) NOT NULL
);
