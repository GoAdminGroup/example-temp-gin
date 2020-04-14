-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: GoAdmin
-- ------------------------------------------------------
-- Server version	5.7.27-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goadmin_menu`
--

DROP TABLE IF EXISTS `goadmin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) unsigned NOT NULL DEFAULT '0',
  `type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `order` int(11) unsigned NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_menu`
--

LOCK TABLES `goadmin_menu` WRITE;
INSERT INTO `goadmin_menu` VALUES (1,0,1,3,'Admin','fa-tasks','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (2,1,1,6,'Users','fa-users','/info/manager',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (3,1,1,4,'Roles','fa-user','/info/roles',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (4,1,1,3,'Permission','fa-ban','/info/permission',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (5,1,1,5,'Menu','fa-bars','/menu',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (6,1,1,7,'Operation log','fa-history','/info/op',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_menu` VALUES (7,0,1,1,'Dashboard','fa-bar-chart','/','','2019-09-10 00:00:00','2020-03-26 16:15:07');
INSERT INTO `goadmin_menu` VALUES (8,0,0,2,'Demo','fa-align-left','/info/demo','Demo','2020-03-26 08:23:20','2020-03-26 08:23:20');
INSERT INTO `goadmin_menu` VALUES (9,8,0,2,'grade','fa-bars','/info/demo_grade','','2020-03-26 08:24:03','2020-03-26 16:25:56');
INSERT INTO `goadmin_menu` VALUES (10,8,0,2,'class','fa-bars','/info/demo_class','','2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_menu` VALUES (11,8,0,2,'student','fa-bars','/info/demo_student','','2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_menu` VALUES (12,8,0,2,'student-class','fa-bars','/info/demo_student_class','','2020-03-26 08:32:58','2020-03-26 16:33:44');
INSERT INTO `goadmin_menu` VALUES (13,8,0,2,'student-score','fa-bars','/info/demo_student_score','','2020-03-26 08:33:29','2020-03-26 08:33:29');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_permissions`
--

DROP TABLE IF EXISTS `goadmin_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_permissions`
--

LOCK TABLES `goadmin_permissions` WRITE;
INSERT INTO `goadmin_permissions` VALUES (1,'All permission','*','','*','2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_permissions` VALUES (2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_permissions` VALUES (3,'demo_grade-operator','demo_grade-operator','GET,PUT,POST,PATCH,OPTIONS','/info/demo_grade\r\n/info/demo_grade/detail\r\n/info/demo_grade/edit\r\n/edit/demo_grade\r\n/info/demo_grade/new\r\n/new/demo_grade\r\n/export/demo_grade','2020-03-26 07:41:29','2020-04-01 11:21:57');
INSERT INTO `goadmin_permissions` VALUES (4,'demo_class-operator','demo_class-operator','','/info/demo_class\r\n/info/demo_class/detail\r\n/info/demo_class/edit\r\n/edit/demo_class\r\n/info/demo_class/new\r\n/new/demo_class\r\n/export/demo_class','2020-03-26 08:27:59','2020-04-01 11:22:45');
INSERT INTO `goadmin_permissions` VALUES (5,'demo_student-operator','demo_student-operator','','/info/demo_student\r\n/info/demo_student/detail\r\n/info/demo_student/edit\r\n/edit/demo_student\r\n/info/demo_student/new\r\n/new/demo_student\r\n/export/demo_student','2020-03-26 08:29:33','2020-04-01 11:23:05');
INSERT INTO `goadmin_permissions` VALUES (6,'demo_student_class-operator','demo_student_class-operator','','/info/demo_student_class\r\n/info/demo_student_class/detail\r\n/info/demo_student_class/edit\r\n/edit/demo_student_class\r\n/info/demo_student_class/new\r\n/new/demo_student_class\r\n/export/demo_student_class','2020-03-26 08:31:48','2020-04-01 11:24:59');
INSERT INTO `goadmin_permissions` VALUES (7,'demo_student_score-operator','demo_student_score-operator','','/info/demo_student_score\r\n/info/demo_student_score/detail\r\n/info/demo_student_score/edit\r\n/edit/demo_student_score\r\n/info/demo_student_score/new\r\n/new/demo_student_score\r\n/export/demo_student_score','2020-03-26 08:32:08','2020-04-01 11:21:31');
INSERT INTO `goadmin_permissions` VALUES (8,'demo_grade-delete','demo_grade-delete','','/delete/demo_grade','2020-04-01 03:26:08','2020-04-01 11:26:08');
INSERT INTO `goadmin_permissions` VALUES (9,'demo_class-delete','demo_class-delete','','/delete/demo_class','2020-04-01 03:26:50','2020-04-01 11:26:50');
INSERT INTO `goadmin_permissions` VALUES (10,'demo_student-delete','demo_student-delete','','/delete/demo_student','2020-04-01 03:27:20','2020-04-01 11:27:20');
INSERT INTO `goadmin_permissions` VALUES (11,'demo_student_class-delete','demo_student_class-delete','','/delete/demo_student_class','2020-04-01 03:27:45','2020-04-01 11:27:45');
INSERT INTO `goadmin_permissions` VALUES (12,'demo_student_score-delete','demo_student_score-delete','','/delete/demo_student_score','2020-04-01 03:28:10','2020-04-01 11:28:10');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_menu`
--

DROP TABLE IF EXISTS `goadmin_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_role_menu` (
  `role_id` int(11) unsigned NOT NULL,
  `menu_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_menu`
--

LOCK TABLES `goadmin_role_menu` WRITE;
INSERT INTO `goadmin_role_menu` VALUES (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_role_menu` VALUES (1,8,'2019-09-11 10:20:55','2019-09-11 10:20:55');
INSERT INTO `goadmin_role_menu` VALUES (2,8,'2019-09-11 10:20:55','2019-09-11 10:20:55');
INSERT INTO `goadmin_role_menu` VALUES (1,7,'2020-03-26 08:15:07','2020-03-26 08:15:07');
INSERT INTO `goadmin_role_menu` VALUES (2,7,'2020-03-26 08:15:07','2020-03-26 08:15:07');
INSERT INTO `goadmin_role_menu` VALUES (3,7,'2020-03-26 08:15:07','2020-03-26 08:15:07');
INSERT INTO `goadmin_role_menu` VALUES (3,8,'2020-03-26 08:23:20','2020-03-26 08:23:20');
INSERT INTO `goadmin_role_menu` VALUES (1,9,'2020-03-26 08:25:56','2020-03-26 08:25:56');
INSERT INTO `goadmin_role_menu` VALUES (2,9,'2020-03-26 08:25:56','2020-03-26 08:25:56');
INSERT INTO `goadmin_role_menu` VALUES (3,9,'2020-03-26 08:25:56','2020-03-26 08:25:56');
INSERT INTO `goadmin_role_menu` VALUES (1,10,'2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_role_menu` VALUES (2,10,'2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_role_menu` VALUES (3,10,'2020-03-26 08:28:45','2020-03-26 08:28:45');
INSERT INTO `goadmin_role_menu` VALUES (1,11,'2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_role_menu` VALUES (2,11,'2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_role_menu` VALUES (3,11,'2020-03-26 08:31:03','2020-03-26 08:31:03');
INSERT INTO `goadmin_role_menu` VALUES (1,13,'2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_role_menu` VALUES (2,13,'2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_role_menu` VALUES (3,13,'2020-03-26 08:33:29','2020-03-26 08:33:29');
INSERT INTO `goadmin_role_menu` VALUES (1,12,'2020-03-26 08:33:44','2020-03-26 08:33:44');
INSERT INTO `goadmin_role_menu` VALUES (2,12,'2020-03-26 08:33:44','2020-03-26 08:33:44');
INSERT INTO `goadmin_role_menu` VALUES (3,12,'2020-03-26 08:33:44','2020-03-26 08:33:44');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_permissions`
--

DROP TABLE IF EXISTS `goadmin_role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_role_permissions` (
  `role_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_permissions`
--

LOCK TABLES `goadmin_role_permissions` WRITE;
INSERT INTO `goadmin_role_permissions` VALUES (1,1,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,2,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,3,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,4,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,5,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,6,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,7,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,8,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,9,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,10,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,11,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (1,12,'2020-04-01 03:28:22','2020-04-01 03:28:22');
INSERT INTO `goadmin_role_permissions` VALUES (2,2,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,3,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,4,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,5,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,6,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (2,7,'2020-03-26 08:35:05','2020-03-26 08:35:05');
INSERT INTO `goadmin_role_permissions` VALUES (3,2,'2020-04-01 03:28:30','2020-04-01 03:28:30');
INSERT INTO `goadmin_role_permissions` VALUES (3,3,'2020-04-01 03:28:30','2020-04-01 03:28:30');
INSERT INTO `goadmin_role_permissions` VALUES (3,4,'2020-04-01 03:28:30','2020-04-01 03:28:30');
INSERT INTO `goadmin_role_permissions` VALUES (3,5,'2020-04-01 03:28:30','2020-04-01 03:28:30');
INSERT INTO `goadmin_role_permissions` VALUES (3,6,'2020-04-01 03:28:30','2020-04-01 03:28:30');
INSERT INTO `goadmin_role_permissions` VALUES (3,7,'2020-04-01 03:28:30','2020-04-01 03:28:30');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_users`
--

DROP TABLE IF EXISTS `goadmin_role_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_role_users` (
  `role_id` int(11) unsigned NOT NULL,
  `user_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_users`
--

LOCK TABLES `goadmin_role_users` WRITE;
INSERT INTO `goadmin_role_users` VALUES (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_role_users` VALUES (2,2,'2020-03-26 08:14:36','2020-03-26 08:14:36');
INSERT INTO `goadmin_role_users` VALUES (3,3,'2020-03-26 07:44:12','2020-03-26 07:44:12');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_roles`
--

DROP TABLE IF EXISTS `goadmin_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_roles`
--

LOCK TABLES `goadmin_roles` WRITE;
INSERT INTO `goadmin_roles` VALUES (1,'Administrator','administrator','2019-09-10 00:00:00','2020-04-01 11:28:22');
INSERT INTO `goadmin_roles` VALUES (2,'Operator','operator','2019-09-10 00:00:00','2020-03-26 16:35:05');
INSERT INTO `goadmin_roles` VALUES (3,'Test','test','2020-03-26 07:33:50','2020-04-01 11:28:30');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_session`
--

DROP TABLE IF EXISTS `goadmin_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_session` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_session`
--

LOCK TABLES `goadmin_session` WRITE;
INSERT INTO `goadmin_session` VALUES (8,'fd6927af-6179-44c1-a157-bedeaac9907c','{\"user_id\":1}','2020-04-01 03:20:39','2020-04-01 03:20:39');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_user_permissions`
--

DROP TABLE IF EXISTS `goadmin_user_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_user_permissions` (
  `user_id` int(11) unsigned NOT NULL,
  `permission_id` int(11) unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_user_permissions`
--

LOCK TABLES `goadmin_user_permissions` WRITE;
INSERT INTO `goadmin_user_permissions` VALUES (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_user_permissions` VALUES (2,2,'2020-03-26 08:14:36','2020-03-26 08:14:36');
UNLOCK TABLES;

--
-- Table structure for table `goadmin_users`
--

DROP TABLE IF EXISTS `goadmin_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_users`
--

LOCK TABLES `goadmin_users` WRITE;
INSERT INTO `goadmin_users` VALUES (1,'admin','$2a$10$bUZFttkyk/rOp1tBriUPC.rnlSdLLloS6hHYZ6RWg1G6QT8X1KI32','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-10 00:00:00','2019-09-10 00:00:00');
INSERT INTO `goadmin_users` VALUES (2,'operator','$2a$10$Trvjo8139K7KDXXORCWAl.7lMl8Q9lZ4F20GvmTbmJD6yM1cz.Tti','Operator','',NULL,'2019-09-10 00:00:00','2020-03-26 16:14:36');
INSERT INTO `goadmin_users` VALUES (3,'test','$2a$10$NCI1bBSYqIHx9VZaAAajPOdjVdIFN6uIcFHVAuxQZiFsNizqwvCwm','test','',NULL,'2020-03-26 07:34:38','2020-03-26 15:44:12');
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-01  3:29:07
-- MySQL dump 10.13  Distrib 5.7.27, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: GoAdmin
-- ------------------------------------------------------
-- Server version	5.7.27-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goadmin_operation_log`
--

DROP TABLE IF EXISTS `goadmin_operation_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `goadmin_operation_log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=271 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-01  3:29:07

CREATE TABLE `goadmin_site` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` longtext COLLATE utf8mb4_unicode_ci,
  `description` varchar(3000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `state` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;