CREATE DATABASE IF NOT EXISTS `printdb`;
grant all privileges on *.* to 'user'@'%';

USE `printdb`;

CREATE TABLE `materials` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `type` varchar(255) NOT NULL,
  `price` float NOT NULL,
  `density` float NOT NULL,
  `color` varchar(255) NOT NULL
);

CREATE TABLE `profiles` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `options` text,
  `material_type` varchar(255) NOT NULL,
  `level_of_detail` varchar(255) NOT NULL
);

CREATE TABLE `printers` (
  `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `power_supply` int NOT NULL,
  `profile_id` int NOT NULL,
  `width` int NOT NULL,
  `length` int NOT NULL,
  `height` int NOT NULL
);

-- fill tables with some data

INSERT INTO `materials` (`type`, `price`, `density`, `color`) VALUES
  ('PLA', 20.0, 1.25, 'red'),
  ('ABS', 25.0, 1.05, 'blue'),
  ('PETG', 30.0, 1.15, 'green');

INSERT INTO `profiles` (`name`, `options`, `material_type`, `level_of_detail`) VALUES
  ('Fast', '', 'PLA', 'low'),
  ('Normal', '', 'PLA', 'medium'),
  ('High', '', 'PLA', 'high');

INSERT INTO `printers` (`name`, `power_supply`, `profile_id`, `width`, `length`, `height`) VALUES
  ('Prusa i3', 220, 1, 200, 200, 200),
  ('Anet A8', 220, 2, 220, 220, 240),
  ('Creality CR-10', 220, 3, 300, 300, 400);
