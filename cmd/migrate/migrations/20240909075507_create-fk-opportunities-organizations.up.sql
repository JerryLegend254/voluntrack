ALTER TABLE `opportunities` ADD FOREIGN KEY (`organization_email`) REFERENCES `organizations` (`email`);
