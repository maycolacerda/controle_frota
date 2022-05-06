-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `veiculos`;
CREATE TABLE `veiculos` (
  `Id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `identificador` varchar(5) NOT NULL,
  `tipo` varchar(50) NOT NULL,
  `marca` varchar(50) NOT NULL,
  `modelo` varchar(50) NOT NULL,
  `rastreador` tinyint(1) DEFAULT NULL,
  `placa` varchar(10) NOT NULL,
  `disponivel` tinyint(1) DEFAULT NULL,
  `km` tinyint(1) DEFAULT NULL,
  `tipo_carteira` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'N/A',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Id` (`Id`),
  UNIQUE KEY `identificador` (`identificador`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `veiculos` (`Id`, `identificador`, `tipo`, `marca`, `modelo`, `rastreador`, `placa`, `disponivel`, `km`, `tipo_carteira`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1,	'A01',	'Carro',	'Ford',	'Fusion',	1,	'ABC-1234',	1,	1,	'B',	'2022-05-04 06:41:47',	'2022-05-04 06:41:47',	NULL),
(2,	'C01',	'Trator',	'John Deere',	'5090',	1,	'N/A',	1,	1,	'N/A',	'2022-05-04 06:41:47',	'2022-05-04 06:41:47',	NULL),
(3,	'a03',	'carro',	'Fiat',	'Uno 2012',	0,	'abc-2345',	0,	0,	'B',	'2022-05-04 06:50:53',	'2022-05-04 06:50:53',	NULL),
(5,	'a04',	'carro',	'Fiat',	'Uno 2013',	0,	'abc-3456',	0,	0,	'B',	'2022-05-04 06:51:44',	'2022-05-04 06:51:44',	NULL);

-- 2022-05-04 06:53:53