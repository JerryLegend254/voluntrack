ALTER TABLE `opportunities` ADD FOREIGN KEY (`status`) REFERENCES `opportunities_status` (`ost_id`);
