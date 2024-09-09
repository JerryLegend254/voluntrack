CREATE TABLE `volunteers` (
  `email` varchar(255) NOT NULL PRIMARY KEY,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `available` boolean NOT NULL DEFAULT true,
  `password` varchar(60) NOT NULL
);
