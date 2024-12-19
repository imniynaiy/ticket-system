# ************************************************************
# Sequel Ace SQL dump
# Version 20077
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: localhost (MySQL 8.0.40)
# Database: some_db
# Generation Time: 2024-12-19 12:23:08 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `log`;

CREATE TABLE `log` (
  `log_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '操作ログ番号',
  `user_id` int unsigned NOT NULL COMMENT 'ユーザID',
  `function_id` int unsigned NOT NULL COMMENT '機能番号',
  `timestamp` datetime DEFAULT NULL COMMENT '操作日時',
  `result` int DEFAULT NULL COMMENT '操作結果　０＝注文中　１＝成功　２＝失敗',
  PRIMARY KEY (`log_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `log_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;

INSERT INTO `log` (`log_id`, `user_id`, `function_id`, `timestamp`, `result`)
VALUES
	(7,3,1,'2024-12-17 12:32:00',0),
	(8,3,1,'2024-12-17 12:34:33',0),
	(9,3,1,'2024-12-17 12:38:05',2),
	(10,3,1,'2024-12-17 12:45:06',2),
	(11,3,1,'2024-12-17 12:58:34',2),
	(12,3,1,'2024-12-17 12:59:27',2),
	(13,3,1,'2024-12-17 13:00:29',2),
	(14,3,1,'2024-12-17 13:00:43',2),
	(15,3,1,'2024-12-17 13:07:33',2),
	(16,3,1,'2024-12-17 13:10:10',0),
	(17,3,1,'2024-12-17 13:18:30',2),
	(18,3,1,'2024-12-17 13:20:05',2),
	(19,3,1,'2024-12-17 13:22:22',1),
	(20,3,1,'2024-12-17 13:24:42',1),
	(21,3,1,'2024-12-17 23:30:45',2),
	(22,5,1,'2024-12-17 23:46:34',1),
	(23,6,1,'2024-12-17 23:46:42',2),
	(24,10,1,'2024-12-17 23:46:42',2),
	(25,13,1,'2024-12-17 23:46:42',1),
	(26,7,1,'2024-12-17 23:46:42',2),
	(27,9,1,'2024-12-17 23:46:42',2),
	(28,12,1,'2024-12-17 23:46:42',2),
	(29,11,1,'2024-12-17 23:46:42',2),
	(30,8,1,'2024-12-17 23:46:42',2),
	(31,4,1,'2024-12-17 23:46:42',2),
	(32,8,1,'2024-12-18 13:22:19',1),
	(33,9,1,'2024-12-18 13:22:19',2),
	(34,6,1,'2024-12-18 13:22:19',1),
	(35,12,1,'2024-12-18 13:22:19',2),
	(36,11,1,'2024-12-18 13:22:19',2),
	(37,13,1,'2024-12-18 13:22:19',2),
	(38,10,1,'2024-12-18 13:22:19',2),
	(39,7,1,'2024-12-18 13:22:19',2),
	(40,5,1,'2024-12-18 13:22:19',2),
	(41,4,1,'2024-12-18 13:22:19',2),
	(42,8,1,'2024-12-18 13:22:51',1),
	(43,12,1,'2024-12-18 13:22:51',2),
	(44,13,1,'2024-12-18 13:22:51',2),
	(45,9,1,'2024-12-18 13:22:51',1),
	(46,10,1,'2024-12-18 13:22:51',2),
	(47,6,1,'2024-12-18 13:22:51',2),
	(48,5,1,'2024-12-18 13:22:51',2),
	(49,11,1,'2024-12-18 13:22:51',2),
	(50,4,1,'2024-12-18 13:22:51',2),
	(51,7,1,'2024-12-18 13:22:51',2),
	(52,4,1,'2024-12-18 13:25:23',2),
	(53,5,1,'2024-12-18 13:25:23',2),
	(54,10,1,'2024-12-18 13:25:23',2),
	(55,6,1,'2024-12-18 13:25:23',2),
	(56,9,1,'2024-12-18 13:25:23',2),
	(57,11,1,'2024-12-18 13:25:23',2),
	(58,13,1,'2024-12-18 13:25:23',2),
	(59,12,1,'2024-12-18 13:25:23',2),
	(60,7,1,'2024-12-18 13:25:23',2),
	(61,8,1,'2024-12-18 13:25:23',2),
	(62,10,1,'2024-12-18 13:25:34',1),
	(63,13,1,'2024-12-18 13:25:34',2),
	(64,7,1,'2024-12-18 13:25:34',2),
	(65,9,1,'2024-12-18 13:25:34',2),
	(66,8,1,'2024-12-18 13:25:34',2),
	(67,6,1,'2024-12-18 13:25:34',2),
	(68,11,1,'2024-12-18 13:25:35',2),
	(69,5,1,'2024-12-18 13:25:35',2),
	(70,12,1,'2024-12-18 13:25:35',2),
	(71,4,1,'2024-12-18 13:25:35',2),
	(72,9,1,'2024-12-18 13:35:09',1),
	(73,13,1,'2024-12-18 13:35:09',2),
	(74,8,1,'2024-12-18 13:35:09',2),
	(75,7,1,'2024-12-18 13:35:09',2),
	(76,6,1,'2024-12-18 13:35:09',2),
	(77,10,1,'2024-12-18 13:35:09',2),
	(78,4,1,'2024-12-18 13:35:09',2),
	(79,12,1,'2024-12-18 13:35:09',2),
	(80,11,1,'2024-12-18 13:35:09',2),
	(81,5,1,'2024-12-18 13:35:09',2);

/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table reservation
# ------------------------------------------------------------

DROP TABLE IF EXISTS `reservation`;

CREATE TABLE `reservation` (
  `reservation_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '予約番号',
  `user_id` int unsigned NOT NULL COMMENT 'ユーザー番号',
  `reservation_date` date NOT NULL COMMENT '予約年月日',
  PRIMARY KEY (`reservation_id`),
  KEY `customer_id` (`user_id`),
  CONSTRAINT `reservation_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `reservation` WRITE;
/*!40000 ALTER TABLE `reservation` DISABLE KEYS */;

INSERT INTO `reservation` (`reservation_id`, `user_id`, `reservation_date`)
VALUES
	(10,3,'2024-01-01'),
	(11,3,'2024-01-01'),
	(13,5,'2024-01-01'),
	(16,13,'2024-01-01'),
	(23,8,'2024-01-01'),
	(24,6,'2024-01-01'),
	(33,8,'2024-01-01'),
	(37,9,'2024-01-01'),
	(53,10,'2024-01-01'),
	(63,9,'2024-01-01');

/*!40000 ALTER TABLE `reservation` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table reservation_detail
# ------------------------------------------------------------

DROP TABLE IF EXISTS `reservation_detail`;

CREATE TABLE `reservation_detail` (
  `reservation_detail_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '予約明細番号',
  `route_id` int unsigned NOT NULL COMMENT '船便番号',
  `seat_id` int unsigned NOT NULL COMMENT '座席番号',
  `passenger_family_name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '乗船者姓',
  `passenger_first_name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '乗船者名',
  `reservation_id` int unsigned NOT NULL,
  PRIMARY KEY (`reservation_detail_id`),
  KEY `route_id` (`route_id`),
  KEY `seat_id` (`seat_id`),
  KEY `reservation_id` (`reservation_id`),
  CONSTRAINT `reservation_detail_ibfk_2` FOREIGN KEY (`route_id`) REFERENCES `route` (`route_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `reservation_detail_ibfk_3` FOREIGN KEY (`seat_id`) REFERENCES `seat` (`seat_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `reservation_detail_ibfk_4` FOREIGN KEY (`reservation_id`) REFERENCES `reservation` (`reservation_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `reservation_detail` WRITE;
/*!40000 ALTER TABLE `reservation_detail` DISABLE KEYS */;

INSERT INTO `reservation_detail` (`reservation_detail_id`, `route_id`, `seat_id`, `passenger_family_name`, `passenger_first_name`, `reservation_id`)
VALUES
	(1,2,19,'Smith','John',10),
	(2,2,20,'Smith','John',11),
	(3,2,15,'Test','User',13),
	(8,2,15,'Test','User',16),
	(13,2,22,'Test','User',23),
	(14,2,22,'Test','User',24),
	(16,2,23,'Test','User',33),
	(20,2,23,'Test','User',37),
	(21,2,24,'Test','User',53),
	(22,2,25,'Test','User',63);

/*!40000 ALTER TABLE `reservation_detail` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table route
# ------------------------------------------------------------

DROP TABLE IF EXISTS `route`;

CREATE TABLE `route` (
  `route_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '船便番号',
  `route_name` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '船便名',
  `departure_time` datetime NOT NULL COMMENT '出発日時',
  `arrival_time` datetime NOT NULL COMMENT '到着日時',
  `departure_from` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '出発地',
  `arrival_to` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '到着地',
  `distance` int NOT NULL COMMENT '航行距離',
  `basic_fee` int NOT NULL COMMENT '基本運賃',
  PRIMARY KEY (`route_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `route` WRITE;
/*!40000 ALTER TABLE `route` DISABLE KEYS */;

INSERT INTO `route` (`route_id`, `route_name`, `departure_time`, `arrival_time`, `departure_from`, `arrival_to`, `distance`, `basic_fee`)
VALUES
	(2,'Tokyo to Osaka','2024-03-20 09:00:00','2024-03-20 12:00:00','Tokyo','Osaka',515,14000),
	(3,'Tokyo to Osaka Express','2024-03-20 08:00:00','2024-03-20 11:00:00','Tokyo','Osaka',515,15000),
	(4,'Tokyo to Osaka','2024-03-20 09:00:00','2024-03-20 12:00:00','Tokyo','Osaka',515,14000),
	(5,'Tokyo to Osaka','2024-03-20 09:00:00','2024-03-20 12:00:00','Tokyo','Osaka',515,14000),
	(6,'Tokyo to Osaka','2024-03-20 09:00:00','2024-03-20 12:00:00','Tokyo','Osaka',515,14000);

/*!40000 ALTER TABLE `route` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table seat
# ------------------------------------------------------------

DROP TABLE IF EXISTS `seat`;

CREATE TABLE `seat` (
  `route_id` int unsigned NOT NULL COMMENT '船便番号',
  `seat_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '座席番号',
  `seatclass_id` int unsigned NOT NULL COMMENT '座席クラス番号',
  `status` int unsigned NOT NULL COMMENT '空席状況　１＝利用可能　０＝売られた',
  PRIMARY KEY (`seat_id`),
  KEY `route_id` (`route_id`),
  KEY `seatclass_id` (`seatclass_id`),
  CONSTRAINT `seat_ibfk_1` FOREIGN KEY (`route_id`) REFERENCES `route` (`route_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `seat_ibfk_2` FOREIGN KEY (`seatclass_id`) REFERENCES `seatclass` (`seatclass_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `seat` WRITE;
/*!40000 ALTER TABLE `seat` DISABLE KEYS */;

INSERT INTO `seat` (`route_id`, `seat_id`, `seatclass_id`, `status`)
VALUES
	(2,3,1,1),
	(2,4,1,1),
	(2,5,1,1),
	(2,6,1,1),
	(2,7,1,1),
	(2,8,1,1),
	(2,9,1,1),
	(2,10,1,1),
	(2,11,1,1),
	(2,12,1,1),
	(2,13,1,1),
	(2,14,1,1),
	(2,15,1,0),
	(2,16,1,1),
	(2,17,1,1),
	(2,18,1,1),
	(2,19,1,0),
	(2,20,1,0),
	(2,21,1,1),
	(2,22,1,0),
	(2,23,1,0),
	(2,24,1,0),
	(2,25,1,0),
	(2,26,1,1),
	(2,27,1,1),
	(2,28,1,1),
	(2,29,1,1),
	(2,30,1,1),
	(2,31,1,1),
	(2,32,1,1),
	(2,33,1,1),
	(2,34,1,1),
	(2,35,1,1),
	(2,36,1,1),
	(2,37,1,1),
	(2,38,1,1),
	(2,39,1,1),
	(2,40,1,1),
	(2,41,1,1),
	(2,42,1,1),
	(2,43,1,1),
	(2,44,1,1),
	(2,45,1,1),
	(2,46,1,1),
	(2,47,1,1),
	(2,48,1,1),
	(2,49,1,1),
	(2,50,1,1),
	(2,51,1,1),
	(2,52,1,1),
	(2,53,1,1),
	(2,54,1,1),
	(2,55,1,1),
	(2,56,1,1),
	(2,57,1,1),
	(2,58,1,1),
	(2,59,2,1),
	(2,60,2,1),
	(2,61,2,1),
	(2,62,2,1),
	(2,63,2,1),
	(2,64,2,1),
	(2,65,2,1),
	(2,66,2,1),
	(2,67,2,1),
	(2,68,2,1),
	(2,69,2,1),
	(2,70,2,1),
	(2,71,2,1),
	(2,72,2,1),
	(2,73,2,1),
	(2,74,2,1),
	(2,75,2,1),
	(2,76,2,1),
	(2,77,2,1),
	(2,78,2,1),
	(2,79,3,1),
	(2,80,3,1),
	(2,81,3,1),
	(2,82,3,1),
	(2,83,3,1),
	(2,84,3,1),
	(2,85,3,1),
	(2,86,3,1),
	(2,87,3,1),
	(2,88,3,1),
	(3,89,1,1),
	(3,90,1,1),
	(3,91,1,1),
	(3,92,1,1),
	(3,93,1,1),
	(3,94,1,1),
	(3,95,1,1),
	(3,96,1,1),
	(3,97,1,1),
	(3,98,1,1),
	(3,99,1,1),
	(3,100,1,1),
	(3,101,1,1),
	(3,102,1,1),
	(3,103,1,1),
	(3,104,1,1),
	(3,105,1,1),
	(3,106,1,1),
	(3,107,1,1),
	(3,108,1,1),
	(3,109,1,1),
	(3,110,1,1),
	(3,111,1,1),
	(3,112,1,1),
	(3,113,1,1),
	(3,114,1,1),
	(3,115,1,1),
	(3,116,1,1),
	(3,117,1,1),
	(3,118,1,1),
	(3,119,1,1),
	(3,120,1,1),
	(3,121,1,1),
	(3,122,1,1),
	(3,123,1,1),
	(3,124,1,1),
	(3,125,1,1),
	(3,126,1,1),
	(3,127,1,1),
	(3,128,1,1),
	(3,129,1,1),
	(3,130,1,1),
	(3,131,1,1),
	(3,132,1,1),
	(3,133,1,1),
	(3,134,1,1),
	(3,135,1,1),
	(3,136,1,1),
	(3,137,1,1),
	(3,138,1,1),
	(3,139,2,1),
	(3,140,2,1),
	(3,141,2,1),
	(3,142,2,1),
	(3,143,2,1),
	(3,144,2,1),
	(3,145,2,1),
	(3,146,2,1),
	(3,147,2,1),
	(3,148,2,1),
	(3,149,2,1),
	(3,150,2,1),
	(3,151,2,1),
	(3,152,2,1),
	(3,153,2,1),
	(3,154,2,1),
	(3,155,2,1),
	(3,156,2,1),
	(3,157,2,1),
	(3,158,2,1),
	(3,159,3,1),
	(3,160,3,1),
	(3,161,3,1),
	(3,162,3,1),
	(3,163,3,1),
	(3,164,3,1),
	(3,165,3,1),
	(3,166,3,1),
	(3,167,3,1),
	(3,168,3,1),
	(4,169,1,1),
	(4,170,1,1),
	(4,171,1,1),
	(4,172,1,1),
	(4,173,1,1),
	(4,174,1,1),
	(4,175,1,1),
	(4,176,1,1),
	(4,177,1,1),
	(4,178,1,1),
	(4,179,1,1),
	(4,180,1,1),
	(4,181,1,1),
	(4,182,1,1),
	(4,183,1,1),
	(4,184,1,1),
	(4,185,1,1),
	(4,186,1,1),
	(4,187,1,1),
	(4,188,1,1),
	(4,189,1,1),
	(4,190,1,1),
	(4,191,1,1),
	(4,192,1,1),
	(4,193,1,1),
	(4,194,1,1),
	(4,195,1,1),
	(4,196,1,1),
	(4,197,1,1),
	(4,198,1,1),
	(4,199,1,1),
	(4,200,1,1),
	(4,201,1,1),
	(4,202,1,1),
	(4,203,1,1),
	(4,204,1,1),
	(4,205,1,1),
	(4,206,1,1),
	(4,207,1,1),
	(4,208,1,1),
	(4,209,1,1),
	(4,210,1,1),
	(4,211,1,1),
	(4,212,1,1),
	(4,213,1,1),
	(4,214,1,1),
	(4,215,1,1),
	(4,216,1,1),
	(4,217,1,1),
	(4,218,1,1),
	(4,219,2,1),
	(4,220,2,1),
	(4,221,2,1),
	(4,222,2,1),
	(4,223,2,1),
	(4,224,2,1),
	(4,225,2,1),
	(4,226,2,1),
	(4,227,2,1),
	(4,228,2,1),
	(4,229,2,1),
	(4,230,2,1),
	(4,231,2,1),
	(4,232,2,1),
	(4,233,2,1),
	(4,234,2,1),
	(4,235,2,1),
	(4,236,2,1),
	(4,237,2,1),
	(4,238,2,1),
	(4,239,3,1),
	(4,240,3,1),
	(4,241,3,1),
	(4,242,3,1),
	(4,243,3,1),
	(4,244,3,1),
	(4,245,3,1),
	(4,246,3,1),
	(4,247,3,1),
	(4,248,3,1),
	(5,249,1,1),
	(5,250,1,1),
	(5,251,1,1),
	(5,252,1,1),
	(5,253,1,1),
	(5,254,1,1),
	(5,255,1,1),
	(5,256,1,1),
	(5,257,1,1),
	(5,258,1,1),
	(5,259,1,1),
	(5,260,1,1),
	(5,261,1,1),
	(5,262,1,1),
	(5,263,1,1),
	(5,264,1,1),
	(5,265,1,1),
	(5,266,1,1),
	(5,267,1,1),
	(5,268,1,1),
	(5,269,1,1),
	(5,270,1,1),
	(5,271,1,1),
	(5,272,1,1),
	(5,273,1,1),
	(5,274,1,1),
	(5,275,1,1),
	(5,276,1,1),
	(5,277,1,1),
	(5,278,1,1),
	(5,279,1,1),
	(5,280,1,1),
	(5,281,1,1),
	(5,282,1,1),
	(5,283,1,1),
	(5,284,1,1),
	(5,285,1,1),
	(5,286,1,1),
	(5,287,1,1),
	(5,288,1,1),
	(5,289,1,1),
	(5,290,1,1),
	(5,291,1,1),
	(5,292,1,1),
	(5,293,1,1),
	(5,294,1,1),
	(5,295,1,1),
	(5,296,1,1),
	(5,297,1,1),
	(5,298,1,1),
	(5,299,2,1),
	(5,300,2,1),
	(5,301,2,1),
	(5,302,2,1),
	(5,303,2,1),
	(5,304,2,1),
	(5,305,2,1),
	(5,306,2,1),
	(5,307,2,1),
	(5,308,2,1),
	(5,309,2,1),
	(5,310,2,1),
	(5,311,2,1),
	(5,312,2,1),
	(5,313,2,1),
	(5,314,2,1),
	(5,315,2,1),
	(5,316,2,1),
	(5,317,2,1),
	(5,318,2,1),
	(5,319,3,1),
	(5,320,3,1),
	(5,321,3,1),
	(5,322,3,1),
	(5,323,3,1),
	(5,324,3,1),
	(5,325,3,1),
	(5,326,3,1),
	(5,327,3,1),
	(5,328,3,1),
	(6,329,1,1),
	(6,330,1,1),
	(6,331,1,1),
	(6,332,1,1),
	(6,333,1,1),
	(6,334,1,1),
	(6,335,1,1),
	(6,336,1,1),
	(6,337,1,1),
	(6,338,1,1),
	(6,339,1,1),
	(6,340,1,1),
	(6,341,1,1),
	(6,342,1,1),
	(6,343,1,1),
	(6,344,1,1),
	(6,345,1,1),
	(6,346,1,1),
	(6,347,1,1),
	(6,348,1,1),
	(6,349,1,1),
	(6,350,1,1),
	(6,351,1,1),
	(6,352,1,1),
	(6,353,1,1),
	(6,354,1,1),
	(6,355,1,1),
	(6,356,1,1),
	(6,357,1,1),
	(6,358,1,1),
	(6,359,1,1),
	(6,360,1,1),
	(6,361,1,1),
	(6,362,1,1),
	(6,363,1,1),
	(6,364,1,1),
	(6,365,1,1),
	(6,366,1,1),
	(6,367,1,1),
	(6,368,1,1),
	(6,369,1,1),
	(6,370,1,1),
	(6,371,1,1),
	(6,372,1,1),
	(6,373,1,1),
	(6,374,1,1),
	(6,375,1,1),
	(6,376,1,1),
	(6,377,1,1),
	(6,378,1,1),
	(6,379,2,1),
	(6,380,2,1),
	(6,381,2,1),
	(6,382,2,1),
	(6,383,2,1),
	(6,384,2,1),
	(6,385,2,1),
	(6,386,2,1),
	(6,387,2,1),
	(6,388,2,1),
	(6,389,2,1),
	(6,390,2,1),
	(6,391,2,1),
	(6,392,2,1),
	(6,393,2,1),
	(6,394,2,1),
	(6,395,2,1),
	(6,396,2,1),
	(6,397,2,1),
	(6,398,2,1),
	(6,399,3,1),
	(6,400,3,1),
	(6,401,3,1),
	(6,402,3,1),
	(6,403,3,1),
	(6,404,3,1),
	(6,405,3,1),
	(6,406,3,1),
	(6,407,3,1),
	(6,408,3,1);

/*!40000 ALTER TABLE `seat` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table seatclass
# ------------------------------------------------------------

DROP TABLE IF EXISTS `seatclass`;

CREATE TABLE `seatclass` (
  `seatclass_id` int unsigned NOT NULL AUTO_INCREMENT,
  `seatclass_name` char(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `factor` decimal(10,2) NOT NULL,
  PRIMARY KEY (`seatclass_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `seatclass` WRITE;
/*!40000 ALTER TABLE `seatclass` DISABLE KEYS */;

INSERT INTO `seatclass` (`seatclass_id`, `seatclass_name`, `factor`)
VALUES
	(1,'Business Class',1.80),
	(2,'First Class',2.50),
	(3,'Normal Class',1.00);

/*!40000 ALTER TABLE `seatclass` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `user_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '顧客番号',
  `agent_flag` varchar(100) DEFAULT NULL COMMENT '代理店フラグ',
  `family_name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '姓',
  `first_name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '名',
  `sex` int DEFAULT NULL COMMENT '性別　０＝女性　１＝男性',
  `tel` int DEFAULT NULL COMMENT '電話番号',
  `address` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '住所',
  `email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'メールアドレス',
  `password_hash` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'パスワードハッシュ値',
  `is_admin` int NOT NULL DEFAULT '0' COMMENT '0＝普通ユーザ　１＝管理員',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`user_id`, `agent_flag`, `family_name`, `first_name`, `sex`, `tel`, `address`, `email`, `password_hash`, `is_admin`)
VALUES
	(1,'','','',0,0,'','admin@example.com','$2a$10$RW57bs3zsCOSuBlTfrvcV.MuRSUFjH0kz87iGMQ0TfxePIEmASTy.',1),
	(3,'','','',0,0,'','user01@example.com','$2a$10$l1kOYqRhwjPtbz1ZZxT.v.cEo91IeT8Z6plHsfQKC7gghRmodI/9q',0),
	(4,'','','',0,0,'','test_user_0@test.com','$2a$10$Fm/Fb6Diu.3uQOKmX22QQeruYr0/eMf3qNOXYf1ZfDcFkmuLgivki',0),
	(5,'','','',0,0,'','test_user_1@test.com','$2a$10$KZBOAsMuJneuCaTmt0NjIeIj9edNaiEzt7GmcCPiC5ST4OevDyiWy',0),
	(6,'','','',0,0,'','test_user_2@test.com','$2a$10$2IaT5mA8XGxbGHGyLaeoqOT/otzOWHFXsd22DtLkxP1hERplDhlTO',0),
	(7,'','','',0,0,'','test_user_3@test.com','$2a$10$ynEqG.9s2tEWdfO3kK85T.pQ6Y/FmgQeu4cMG3YvPy1Qi7Ljd7pYK',0),
	(8,'','','',0,0,'','test_user_4@test.com','$2a$10$hsPTdqqgjVTfbw3MAI9TZ.bwOkFeaLbwNESr0eGqbd/A/zjUFN6jm',0),
	(9,'','','',0,0,'','test_user_5@test.com','$2a$10$JpkixfUg1F5q0CEGg6nPZ.UI6bw5MoUl7//DgIStN22wtRrtzXRG2',0),
	(10,'','','',0,0,'','test_user_6@test.com','$2a$10$7U5LHVDXrCO.FQnHEup5qOGdA3UJG5RMHjUPkeYuz/RvnRSUtBuua',0),
	(11,'','','',0,0,'','test_user_7@test.com','$2a$10$68E2fcw.jRc1tetOfbHgMOz9Up7uuyWPkMUbeCy39idAycfLgQKAa',0),
	(12,'','','',0,0,'','test_user_8@test.com','$2a$10$aQCXOfp2C39cFTSjxUhjG.n/iTpD//lbj/PhiLZjZfwzbqBzQdNgK',0),
	(13,'','','',0,0,'','test_user_9@test.com','$2a$10$OL5XYrgoP9tTPBgO9EZxIO6B6KSORUHn5RVXi/33nir168A6QMdcK',0),
	(14,'','','',0,0,'','user01','$2a$10$OO5Ld2LyjQszm3INrZTPdO04SKbaJUWkhwU.FVPqTT5loQiWJrYnK',0),
	(15,'','','',0,0,'','user02@example.com','$2a$10$TZXDjEUl7Ku/hR5M45yFK.4MfJM7ACkGZQurmuajXIsGpvLtDSed6',0);

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
