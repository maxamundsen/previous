CREATE TABLE `sessions` (
  `id` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `useragent` varchar(255) DEFAULT NULL,
  `ipaddr` varchar(45) DEFAULT NULL,
  `logintime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
