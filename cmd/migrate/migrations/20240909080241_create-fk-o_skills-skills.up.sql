ALTER TABLE `opportunity_skills` ADD FOREIGN KEY (`opportunity_id`) REFERENCES `opportunities` (`opportunity_id`);
