-- MySQL dump 10.13  Distrib 5.7.32, for osx10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: game25_test_user0
-- ------------------------------------------------------
-- Server version	5.7.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `game25_test_user0`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `game25_test_user0` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;

USE `game25_test_user0`;

--
-- Table structure for table `log_billing_client_operation`
--

DROP TABLE IF EXISTS `log_billing_client_operation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_billing_client_operation` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `product_type` tinyint(3) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `platform` tinyint(3) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL,
  `message` text NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`),
  KEY `ix_user_id_platform_product_type_product_id_created_at` (`user_id`,`platform`,`product_type`,`product_id`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_billing_client_operation`
--

LOCK TABLES `log_billing_client_operation` WRITE;
/*!40000 ALTER TABLE `log_billing_client_operation` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_billing_client_operation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_user_caution`
--

DROP TABLE IF EXISTS `log_user_caution`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_user_caution` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `caution_id` bigint(20) unsigned NOT NULL,
  `caution_scene` tinyint(3) unsigned NOT NULL,
  `caution_type` tinyint(3) unsigned NOT NULL,
  `caution_title` text,
  `caution_message` text,
  `start_at` int(10) unsigned NOT NULL,
  `end_at` int(10) unsigned DEFAULT NULL,
  `looked_at` int(10) unsigned DEFAULT NULL,
  `reason` tinyint(3) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_user_caution`
--

LOCK TABLES `log_user_caution` WRITE;
/*!40000 ALTER TABLE `log_user_caution` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_user_caution` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_user_model`
--

DROP TABLE IF EXISTS `log_user_model`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_user_model` (
  `user_id` int(10) unsigned NOT NULL,
  `client_timestamp` bigint(20) unsigned NOT NULL,
  `server_timestamp` int(10) unsigned NOT NULL,
  `idx` int(10) unsigned NOT NULL,
  `category` varchar(191) NOT NULL,
  `val` longblob NOT NULL,
  PRIMARY KEY (`user_id`,`client_timestamp`,`idx`),
  KEY `server_timestamp_index` (`server_timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPRESSED
/*!50100 PARTITION BY RANGE ( client_timestamp)
(PARTITION p201909 VALUES LESS THAN (1569855600000) COMMENT = '2019-09 JST' ENGINE = InnoDB,
 PARTITION p201910 VALUES LESS THAN (1572534000000) COMMENT = '2019-10 JST' ENGINE = InnoDB,
 PARTITION p201911 VALUES LESS THAN (1575126000000) COMMENT = '2019-11 JST' ENGINE = InnoDB,
 PARTITION p201912 VALUES LESS THAN (1577804400000) COMMENT = '2019-12 JST' ENGINE = InnoDB,
 PARTITION p202001 VALUES LESS THAN (1580482800000) COMMENT = '2020-01 JST' ENGINE = InnoDB,
 PARTITION p202002 VALUES LESS THAN (1582988400000) COMMENT = '2020-02 JST' ENGINE = InnoDB,
 PARTITION p202003 VALUES LESS THAN (1585666800000) COMMENT = '2020-03 JST' ENGINE = InnoDB,
 PARTITION p202004 VALUES LESS THAN (1588258800000) COMMENT = '2020-04 JST' ENGINE = InnoDB,
 PARTITION p202005 VALUES LESS THAN (1590937200000) COMMENT = '2020-05 JST' ENGINE = InnoDB,
 PARTITION p202006 VALUES LESS THAN (1593529200000) COMMENT = '2020-06 JST' ENGINE = InnoDB,
 PARTITION p202007 VALUES LESS THAN (1596207600000) COMMENT = '2020-07 JST' ENGINE = InnoDB,
 PARTITION p202008 VALUES LESS THAN (1598886000000) COMMENT = '2020-08 JST' ENGINE = InnoDB,
 PARTITION p202009 VALUES LESS THAN (1601478000000) COMMENT = '2020-09 JST' ENGINE = InnoDB,
 PARTITION p202010 VALUES LESS THAN (1604156400000) COMMENT = '2020-10 JST' ENGINE = InnoDB,
 PARTITION p202011 VALUES LESS THAN (1606748400000) COMMENT = '2020-11 JST' ENGINE = InnoDB,
 PARTITION p202012 VALUES LESS THAN (1609426800000) COMMENT = '2020-12 JST' ENGINE = InnoDB,
 PARTITION p202101 VALUES LESS THAN (1612105200000) COMMENT = '2021-01 JST' ENGINE = InnoDB,
 PARTITION p202102 VALUES LESS THAN (1614524400000) COMMENT = '2021-02 JST' ENGINE = InnoDB,
 PARTITION p202103 VALUES LESS THAN (1617202800000) COMMENT = '2021-03 JST' ENGINE = InnoDB,
 PARTITION p202104 VALUES LESS THAN (1619794800000) COMMENT = '2021-04 JST' ENGINE = InnoDB,
 PARTITION p202105 VALUES LESS THAN (1622473200000) COMMENT = '2021-05 JST' ENGINE = InnoDB,
 PARTITION p202106 VALUES LESS THAN (1625065200000) COMMENT = '2021-06 JST' ENGINE = InnoDB,
 PARTITION p202107 VALUES LESS THAN (1627743600000) COMMENT = '2021-07 JST' ENGINE = InnoDB,
 PARTITION p202108 VALUES LESS THAN (1630422000000) COMMENT = '2021-08 JST' ENGINE = InnoDB,
 PARTITION p202109 VALUES LESS THAN (1633014000000) COMMENT = '2021-09 JST' ENGINE = InnoDB,
 PARTITION p202110 VALUES LESS THAN (1635692400000) COMMENT = '2021-10 JST' ENGINE = InnoDB,
 PARTITION p202111 VALUES LESS THAN (1638284400000) COMMENT = '2021-11 JST' ENGINE = InnoDB,
 PARTITION p202112 VALUES LESS THAN MAXVALUE COMMENT = '2021-12 JST' ENGINE = InnoDB) */;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_user_model`
--

LOCK TABLES `log_user_model` WRITE;
/*!40000 ALTER TABLE `log_user_model` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_user_model` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_user_set_birth_month_for_billing`
--

DROP TABLE IF EXISTS `log_user_set_birth_month_for_billing`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_user_set_birth_month_for_billing` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `birth_month_for_billing` int(10) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_user_set_birth_month_for_billing`
--

LOCK TABLES `log_user_set_birth_month_for_billing` WRITE;
/*!40000 ALTER TABLE `log_user_set_birth_month_for_billing` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_user_set_birth_month_for_billing` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_user_set_birthday`
--

DROP TABLE IF EXISTS `log_user_set_birthday`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_user_set_birthday` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `birth_month` int(10) unsigned DEFAULT NULL,
  `birth_day` int(10) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_user_set_birthday`
--

LOCK TABLES `log_user_set_birthday` WRITE;
/*!40000 ALTER TABLE `log_user_set_birthday` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_user_set_birthday` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_virtual_money_consume_history_detail`
--

DROP TABLE IF EXISTS `log_virtual_money_consume_history_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_virtual_money_consume_history_detail` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `history_nsec` bigint(20) unsigned NOT NULL,
  `deposit_history_detail_nsec` bigint(20) unsigned NOT NULL,
  `platform` tinyint(3) unsigned NOT NULL,
  `virtual_money_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `free_balance` int(10) unsigned NOT NULL,
  `paid_balance` int(10) unsigned NOT NULL,
  `is_exclude_accounting` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`),
  KEY `ix_user_id` (`user_id`),
  KEY `ix_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_virtual_money_consume_history_detail`
--

LOCK TABLES `log_virtual_money_consume_history_detail` WRITE;
/*!40000 ALTER TABLE `log_virtual_money_consume_history_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_virtual_money_consume_history_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log_virtual_money_deposit_history_detail`
--

DROP TABLE IF EXISTS `log_virtual_money_deposit_history_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log_virtual_money_deposit_history_detail` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `history_nsec` bigint(20) unsigned NOT NULL,
  `balance_nsec` bigint(20) unsigned NOT NULL,
  `platform` tinyint(3) unsigned NOT NULL,
  `virtual_money_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `free_balance` int(10) unsigned NOT NULL,
  `paid_balance` int(10) unsigned NOT NULL,
  `product_content_master_id` int(10) unsigned DEFAULT NULL,
  `is_exclude_accounting` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`),
  KEY `ix_user_id` (`user_id`),
  KEY `ix_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log_virtual_money_deposit_history_detail`
--

LOCK TABLES `log_virtual_money_deposit_history_detail` WRITE;
/*!40000 ALTER TABLE `log_virtual_money_deposit_history_detail` DISABLE KEYS */;
/*!40000 ALTER TABLE `log_virtual_money_deposit_history_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_accessory`
--

DROP TABLE IF EXISTS `user_accessory`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_accessory` (
  `user_id` int(10) unsigned NOT NULL,
  `user_accessory_id` bigint(20) unsigned NOT NULL,
  `accessory_master_id` int(10) unsigned NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `grade` int(10) unsigned NOT NULL,
  `attribute` tinyint(3) unsigned NOT NULL,
  `passive_skill_1_id` int(10) unsigned DEFAULT NULL,
  `passive_skill_1_level` int(10) unsigned DEFAULT NULL,
  `passive_skill_2_id` int(10) unsigned DEFAULT NULL,
  `passive_skill_2_level` int(10) unsigned DEFAULT NULL,
  `is_lock` tinyint(1) NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `acquired_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`user_accessory_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_accessory`
--

LOCK TABLES `user_accessory` WRITE;
/*!40000 ALTER TABLE `user_accessory` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_accessory` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_accessory_level_up_item`
--

DROP TABLE IF EXISTS `user_accessory_level_up_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_accessory_level_up_item` (
  `user_id` int(10) unsigned NOT NULL,
  `accessory_level_up_item_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`accessory_level_up_item_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_accessory_level_up_item`
--

LOCK TABLES `user_accessory_level_up_item` WRITE;
/*!40000 ALTER TABLE `user_accessory_level_up_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_accessory_level_up_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_accessory_rarity_up_item`
--

DROP TABLE IF EXISTS `user_accessory_rarity_up_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_accessory_rarity_up_item` (
  `user_id` int(10) unsigned NOT NULL,
  `accessory_rarity_up_item_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`accessory_rarity_up_item_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_accessory_rarity_up_item`
--

LOCK TABLES `user_accessory_rarity_up_item` WRITE;
/*!40000 ALTER TABLE `user_accessory_rarity_up_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_accessory_rarity_up_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_accessory_unique`
--

DROP TABLE IF EXISTS `user_accessory_unique`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_accessory_unique` (
  `user_id` int(10) unsigned NOT NULL,
  `accessory_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`accessory_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_accessory_unique`
--

LOCK TABLES `user_accessory_unique` WRITE;
/*!40000 ALTER TABLE `user_accessory_unique` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_accessory_unique` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_billing_destruction`
--

DROP TABLE IF EXISTS `user_billing_destruction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_billing_destruction` (
  `user_id` int(10) unsigned NOT NULL,
  `error_type` tinyint(3) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_billing_destruction`
--

LOCK TABLES `user_billing_destruction` WRITE;
/*!40000 ALTER TABLE `user_billing_destruction` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_billing_destruction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_billing_monthly_payment`
--

DROP TABLE IF EXISTS `user_billing_monthly_payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_billing_monthly_payment` (
  `user_id` int(10) unsigned NOT NULL,
  `year` int(10) unsigned NOT NULL,
  `month` int(10) unsigned NOT NULL,
  `price` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`year`,`month`),
  KEY `ix_year_month` (`year`,`month`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_billing_monthly_payment`
--

LOCK TABLES `user_billing_monthly_payment` WRITE;
/*!40000 ALTER TABLE `user_billing_monthly_payment` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_billing_monthly_payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_billing_purchase_counter`
--

DROP TABLE IF EXISTS `user_billing_purchase_counter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_billing_purchase_counter` (
  `user_id` int(10) unsigned NOT NULL,
  `product_type` tinyint(3) unsigned NOT NULL,
  `product_id` int(10) unsigned NOT NULL,
  `year` int(10) unsigned NOT NULL,
  `month` int(10) unsigned NOT NULL,
  `day` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `last_purchased_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`product_type`,`product_id`,`year`,`month`,`day`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_billing_purchase_counter`
--

LOCK TABLES `user_billing_purchase_counter` WRITE;
/*!40000 ALTER TABLE `user_billing_purchase_counter` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_billing_purchase_counter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_card`
--

DROP TABLE IF EXISTS `user_card`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_card` (
  `user_id` int(10) unsigned NOT NULL,
  `card_master_id` int(10) unsigned NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `love_point` int(10) unsigned NOT NULL,
  `is_favorite` tinyint(1) NOT NULL,
  `is_awakening` tinyint(1) NOT NULL,
  `is_awakening_image` tinyint(1) NOT NULL,
  `is_all_training_activated` tinyint(1) NOT NULL,
  `max_free_passive_skill` int(10) unsigned NOT NULL,
  `grade` int(10) unsigned NOT NULL,
  `training_life` int(10) unsigned NOT NULL,
  `training_attack` int(10) unsigned NOT NULL,
  `training_dexterity` int(10) unsigned NOT NULL,
  `active_skill_level` int(10) unsigned NOT NULL,
  `passive_skill_a_level` int(10) unsigned NOT NULL,
  `passive_skill_b_level` int(10) unsigned NOT NULL,
  `passive_skill_c_level` int(10) unsigned NOT NULL,
  `additional_passive_skill_1_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_2_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_3_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_4_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_5_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_6_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_7_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_8_id` int(10) unsigned NOT NULL,
  `additional_passive_skill_9_id` int(10) unsigned NOT NULL,
  `acquired_at` int(10) unsigned NOT NULL,
  `live_join_count` int(10) unsigned NOT NULL,
  `active_skill_play_count` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`card_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_card`
--

LOCK TABLES `user_card` WRITE;
/*!40000 ALTER TABLE `user_card` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_card` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_card_training_tree_cell`
--

DROP TABLE IF EXISTS `user_card_training_tree_cell`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_card_training_tree_cell` (
  `user_id` int(10) unsigned NOT NULL,
  `card_m_id` int(10) unsigned NOT NULL,
  `cell_id` int(10) unsigned NOT NULL,
  `activated_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`card_m_id`,`cell_id`),
  KEY `ix_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_card_training_tree_cell`
--

LOCK TABLES `user_card_training_tree_cell` WRITE;
/*!40000 ALTER TABLE `user_card_training_tree_cell` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_card_training_tree_cell` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_current_live`
--

DROP TABLE IF EXISTS `user_current_live`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_current_live` (
  `user_id` int(10) unsigned NOT NULL,
  `live_id` bigint(20) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  `live_difficulty_master_id` int(10) unsigned NOT NULL,
  `started_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_current_live`
--

LOCK TABLES `user_current_live` WRITE;
/*!40000 ALTER TABLE `user_current_live` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_current_live` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_custom_background`
--

DROP TABLE IF EXISTS `user_custom_background`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_custom_background` (
  `user_id` int(10) unsigned NOT NULL,
  `custom_background_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`custom_background_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_custom_background`
--

LOCK TABLES `user_custom_background` WRITE;
/*!40000 ALTER TABLE `user_custom_background` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_custom_background` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_daily_mission`
--

DROP TABLE IF EXISTS `user_daily_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_daily_mission` (
  `user_id` int(10) unsigned NOT NULL,
  `mission_m_id` int(10) unsigned NOT NULL,
  `mission_start_count` int(10) unsigned NOT NULL,
  `mission_count` int(10) unsigned NOT NULL,
  `is_cleared` tinyint(1) NOT NULL,
  `is_received_reward` tinyint(1) NOT NULL,
  `cleared_expired_at` int(10) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`mission_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_daily_mission`
--

LOCK TABLES `user_daily_mission` WRITE;
/*!40000 ALTER TABLE `user_daily_mission` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_daily_mission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_emblem`
--

DROP TABLE IF EXISTS `user_emblem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_emblem` (
  `user_id` int(10) unsigned NOT NULL,
  `emblem_m_id` int(10) unsigned NOT NULL,
  `emblem_param` varchar(191) DEFAULT NULL,
  `acquired_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`emblem_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_emblem`
--

LOCK TABLES `user_emblem` WRITE;
/*!40000 ALTER TABLE `user_emblem` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_emblem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_common_show_result`
--

DROP TABLE IF EXISTS `user_event_common_show_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_common_show_result` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `event_id` int(10) unsigned NOT NULL,
  `result_at` int(10) unsigned NOT NULL,
  `end_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_common_show_result`
--

LOCK TABLES `user_event_common_show_result` WRITE;
/*!40000 ALTER TABLE `user_event_common_show_result` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_common_show_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_marathon`
--

DROP TABLE IF EXISTS `user_event_marathon`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_marathon` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `event_point` int(10) unsigned NOT NULL,
  `opened_story_number` int(10) unsigned NOT NULL,
  `read_story_number` int(10) unsigned NOT NULL,
  `is_board_update` tinyint(1) NOT NULL,
  `is_use_booster` tinyint(1) NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_marathon`
--

LOCK TABLES `user_event_marathon` WRITE;
/*!40000 ALTER TABLE `user_event_marathon` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_marathon` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_marathon_booster`
--

DROP TABLE IF EXISTS `user_event_marathon_booster`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_marathon_booster` (
  `user_id` int(10) unsigned NOT NULL,
  `event_item_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_marathon_booster`
--

LOCK TABLES `user_event_marathon_booster` WRITE;
/*!40000 ALTER TABLE `user_event_marathon_booster` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_marathon_booster` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_marathon_result`
--

DROP TABLE IF EXISTS `user_event_marathon_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_marathon_result` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `event_marathon_id` int(10) unsigned NOT NULL,
  `result_at` int(10) unsigned NOT NULL,
  `end_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_marathon_result`
--

LOCK TABLES `user_event_marathon_result` WRITE;
/*!40000 ALTER TABLE `user_event_marathon_result` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_marathon_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_mining`
--

DROP TABLE IF EXISTS `user_event_mining`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_mining` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `event_point` int(10) unsigned NOT NULL,
  `event_voltage_point` int(10) unsigned NOT NULL,
  `opened_story_number` int(10) unsigned NOT NULL,
  `read_story_number` int(10) unsigned NOT NULL,
  `is_add_new_panel` tinyint(1) NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_mining`
--

LOCK TABLES `user_event_mining` WRITE;
/*!40000 ALTER TABLE `user_event_mining` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_mining` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_mining_nice`
--

DROP TABLE IF EXISTS `user_event_mining_nice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_mining_nice` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `thumbnail_cell_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`,`thumbnail_cell_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_mining_nice`
--

LOCK TABLES `user_event_mining_nice` WRITE;
/*!40000 ALTER TABLE `user_event_mining_nice` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_mining_nice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_event_mining_voltage`
--

DROP TABLE IF EXISTS `user_event_mining_voltage`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_event_mining_voltage` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `live_id` int(10) unsigned NOT NULL,
  `frame_no` int(10) unsigned NOT NULL,
  `live_difficulty_id` int(10) unsigned NOT NULL,
  `voltage` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`,`live_id`,`frame_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_event_mining_voltage`
--

LOCK TABLES `user_event_mining_voltage` WRITE;
/*!40000 ALTER TABLE `user_event_mining_voltage` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_event_mining_voltage` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_exchange_event_mining_content`
--

DROP TABLE IF EXISTS `user_exchange_event_mining_content`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_exchange_event_mining_content` (
  `user_id` int(10) unsigned NOT NULL,
  `event_item_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_exchange_event_mining_content`
--

LOCK TABLES `user_exchange_event_mining_content` WRITE;
/*!40000 ALTER TABLE `user_exchange_event_mining_content` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_exchange_event_mining_content` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_exchange_event_point`
--

DROP TABLE IF EXISTS `user_exchange_event_point`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_exchange_event_point` (
  `user_id` int(10) unsigned NOT NULL,
  `exchange_event_point_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`exchange_event_point_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_exchange_event_point`
--

LOCK TABLES `user_exchange_event_point` WRITE;
/*!40000 ALTER TABLE `user_exchange_event_point` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_exchange_event_point` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gacha_daily`
--

DROP TABLE IF EXISTS `user_gacha_daily`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gacha_daily` (
  `user_id` int(10) unsigned NOT NULL,
  `gacha_draw_m_id` int(10) unsigned NOT NULL,
  `day_id` int(10) unsigned NOT NULL,
  `count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`gacha_draw_m_id`,`day_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gacha_daily`
--

LOCK TABLES `user_gacha_daily` WRITE;
/*!40000 ALTER TABLE `user_gacha_daily` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gacha_daily` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gacha_point`
--

DROP TABLE IF EXISTS `user_gacha_point`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gacha_point` (
  `user_id` int(10) unsigned NOT NULL,
  `point_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`point_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gacha_point`
--

LOCK TABLES `user_gacha_point` WRITE;
/*!40000 ALTER TABLE `user_gacha_point` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gacha_point` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gacha_stepup`
--

DROP TABLE IF EXISTS `user_gacha_stepup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gacha_stepup` (
  `user_id` int(10) unsigned NOT NULL,
  `gacha_m_id` int(10) unsigned NOT NULL,
  `loop_count` int(10) unsigned NOT NULL,
  `step_count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`gacha_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gacha_stepup`
--

LOCK TABLES `user_gacha_stepup` WRITE;
/*!40000 ALTER TABLE `user_gacha_stepup` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gacha_stepup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gacha_ticket`
--

DROP TABLE IF EXISTS `user_gacha_ticket`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gacha_ticket` (
  `user_id` int(10) unsigned NOT NULL,
  `ticket_master_id` int(10) unsigned NOT NULL,
  `normal_amount` int(10) unsigned NOT NULL,
  `apple_amount` int(10) unsigned NOT NULL,
  `google_amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`ticket_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gacha_ticket`
--

LOCK TABLES `user_gacha_ticket` WRITE;
/*!40000 ALTER TABLE `user_gacha_ticket` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gacha_ticket` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_game_settings_push_notification`
--

DROP TABLE IF EXISTS `user_game_settings_push_notification`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_game_settings_push_notification` (
  `user_id` int(10) unsigned NOT NULL,
  `enable_push_notification_setting` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_game_settings_push_notification`
--

LOCK TABLES `user_game_settings_push_notification` WRITE;
/*!40000 ALTER TABLE `user_game_settings_push_notification` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_game_settings_push_notification` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gdpr_consent`
--

DROP TABLE IF EXISTS `user_gdpr_consent`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gdpr_consent` (
  `user_id` int(10) unsigned NOT NULL,
  `gdpr_type` tinyint(3) unsigned NOT NULL,
  `has_consented` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`gdpr_type`),
  KEY `ix_created_at` (`created_at`),
  KEY `ix_user_id_updated_at` (`user_id`,`updated_at`),
  KEY `ix_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gdpr_consent`
--

LOCK TABLES `user_gdpr_consent` WRITE;
/*!40000 ALTER TABLE `user_gdpr_consent` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gdpr_consent` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gift_box`
--

DROP TABLE IF EXISTS `user_gift_box`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gift_box` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `gift_box_m_id` int(10) unsigned NOT NULL,
  `day` int(10) unsigned NOT NULL,
  `shop_product_master_id` int(10) unsigned NOT NULL,
  `is_expired` tinyint(1) NOT NULL,
  `purchased_at` int(10) unsigned NOT NULL,
  `last_received_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`,`gift_box_m_id`,`shop_product_master_id`),
  KEY `ix_user_id_is_expired` (`user_id`,`is_expired`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gift_box`
--

LOCK TABLES `user_gift_box` WRITE;
/*!40000 ALTER TABLE `user_gift_box` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gift_box` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_gps_present_received`
--

DROP TABLE IF EXISTS `user_gps_present_received`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_gps_present_received` (
  `user_id` int(10) unsigned NOT NULL,
  `campaign_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`campaign_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_gps_present_received`
--

LOCK TABLES `user_gps_present_received` WRITE;
/*!40000 ALTER TABLE `user_gps_present_received` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_gps_present_received` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_grade_up_item`
--

DROP TABLE IF EXISTS `user_grade_up_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_grade_up_item` (
  `user_id` int(10) unsigned NOT NULL,
  `item_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`item_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_grade_up_item`
--

LOCK TABLES `user_grade_up_item` WRITE;
/*!40000 ALTER TABLE `user_grade_up_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_grade_up_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info_trigger_basic`
--

DROP TABLE IF EXISTS `user_info_trigger_basic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_info_trigger_basic` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `limit_at` int(10) unsigned DEFAULT NULL,
  `description` text,
  `param_int` int(10) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info_trigger_basic`
--

LOCK TABLES `user_info_trigger_basic` WRITE;
/*!40000 ALTER TABLE `user_info_trigger_basic` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_info_trigger_basic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info_trigger_card_grade_up`
--

DROP TABLE IF EXISTS `user_info_trigger_card_grade_up`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_info_trigger_card_grade_up` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `card_master_id` int(10) unsigned NOT NULL,
  `before_love_level_limit` int(10) unsigned NOT NULL,
  `after_love_level_limit` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info_trigger_card_grade_up`
--

LOCK TABLES `user_info_trigger_card_grade_up` WRITE;
/*!40000 ALTER TABLE `user_info_trigger_card_grade_up` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_info_trigger_card_grade_up` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info_trigger_gacha_point_exchange`
--

DROP TABLE IF EXISTS `user_info_trigger_gacha_point_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_info_trigger_gacha_point_exchange` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `gacha_master_id` int(10) unsigned NOT NULL,
  `point_1_master_id` int(10) unsigned NOT NULL,
  `point_1_before_amount` int(10) unsigned NOT NULL,
  `point_1_after_amount` int(10) unsigned NOT NULL,
  `point_2_master_id` int(10) unsigned NOT NULL,
  `point_2_before_amount` int(10) unsigned NOT NULL,
  `point_2_after_amount` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info_trigger_gacha_point_exchange`
--

LOCK TABLES `user_info_trigger_gacha_point_exchange` WRITE;
/*!40000 ALTER TABLE `user_info_trigger_gacha_point_exchange` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_info_trigger_gacha_point_exchange` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_info_trigger_member_love_level_up`
--

DROP TABLE IF EXISTS `user_info_trigger_member_love_level_up`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_info_trigger_member_love_level_up` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `member_master_id` int(10) unsigned NOT NULL,
  `before_love_level` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_info_trigger_member_love_level_up`
--

LOCK TABLES `user_info_trigger_member_love_level_up` WRITE;
/*!40000 ALTER TABLE `user_info_trigger_member_love_level_up` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_info_trigger_member_love_level_up` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_language`
--

DROP TABLE IF EXISTS `user_language`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_language` (
  `user_id` int(10) unsigned NOT NULL,
  `lang` varchar(191) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_language`
--

LOCK TABLES `user_language` WRITE;
/*!40000 ALTER TABLE `user_language` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_language` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_lesson_deck`
--

DROP TABLE IF EXISTS `user_lesson_deck`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_lesson_deck` (
  `user_id` int(10) unsigned NOT NULL,
  `user_lesson_deck_id` int(10) unsigned NOT NULL,
  `name` varchar(191) NOT NULL,
  `card_master_id_1` int(10) unsigned DEFAULT NULL,
  `card_master_id_2` int(10) unsigned DEFAULT NULL,
  `card_master_id_3` int(10) unsigned DEFAULT NULL,
  `card_master_id_4` int(10) unsigned DEFAULT NULL,
  `card_master_id_5` int(10) unsigned DEFAULT NULL,
  `card_master_id_6` int(10) unsigned DEFAULT NULL,
  `card_master_id_7` int(10) unsigned DEFAULT NULL,
  `card_master_id_8` int(10) unsigned DEFAULT NULL,
  `card_master_id_9` int(10) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`user_lesson_deck_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_lesson_deck`
--

LOCK TABLES `user_lesson_deck` WRITE;
/*!40000 ALTER TABLE `user_lesson_deck` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_lesson_deck` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_lesson_enhancing_item`
--

DROP TABLE IF EXISTS `user_lesson_enhancing_item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_lesson_enhancing_item` (
  `user_id` int(10) unsigned NOT NULL,
  `enhancing_item_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`enhancing_item_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_lesson_enhancing_item`
--

LOCK TABLES `user_lesson_enhancing_item` WRITE;
/*!40000 ALTER TABLE `user_lesson_enhancing_item` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_lesson_enhancing_item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_lesson_result`
--

DROP TABLE IF EXISTS `user_lesson_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_lesson_result` (
  `user_id` int(10) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  `lesson_menu_id1` int(10) unsigned NOT NULL,
  `lesson_menu_id2` int(10) unsigned DEFAULT NULL,
  `lesson_menu_id3` int(10) unsigned DEFAULT NULL,
  `drop_item_json` text NOT NULL,
  `drop_skill_json` text NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_lesson_result`
--

LOCK TABLES `user_lesson_result` WRITE;
/*!40000 ALTER TABLE `user_lesson_result` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_lesson_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live`
--

DROP TABLE IF EXISTS `user_live`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live` (
  `user_id` int(10) unsigned NOT NULL,
  `live_id` bigint(20) unsigned NOT NULL,
  `story_cell_id` int(10) unsigned DEFAULT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  `note_drop` text NOT NULL,
  `is_autoplay` tinyint(1) NOT NULL,
  `autoplay_judgestat` text NOT NULL,
  `magnification` int(10) unsigned NOT NULL,
  `partner_user_id` int(10) unsigned NOT NULL,
  `partner_card_master_id` int(10) unsigned NOT NULL,
  `live_difficulty_master_id` int(10) unsigned NOT NULL,
  `live_started_at` int(10) unsigned NOT NULL,
  `finish_status` tinyint(3) unsigned DEFAULT NULL,
  `is_use_event_marathon_booster` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live`
--

LOCK TABLES `user_live` WRITE;
/*!40000 ALTER TABLE `user_live` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_daily`
--

DROP TABLE IF EXISTS `user_live_daily`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_daily` (
  `user_id` int(10) unsigned NOT NULL,
  `live_daily_id` int(10) unsigned NOT NULL,
  `play_count_per_day` int(10) unsigned NOT NULL,
  `last_play_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_daily_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_daily`
--

LOCK TABLES `user_live_daily` WRITE;
/*!40000 ALTER TABLE `user_live_daily` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_daily` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_deck`
--

DROP TABLE IF EXISTS `user_live_deck`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_deck` (
  `user_id` int(10) unsigned NOT NULL,
  `user_live_deck_id` int(10) unsigned NOT NULL,
  `name` varchar(191) NOT NULL,
  `card_master_id_1` int(10) unsigned DEFAULT NULL,
  `card_master_id_2` int(10) unsigned DEFAULT NULL,
  `card_master_id_3` int(10) unsigned DEFAULT NULL,
  `card_master_id_4` int(10) unsigned DEFAULT NULL,
  `card_master_id_5` int(10) unsigned DEFAULT NULL,
  `card_master_id_6` int(10) unsigned DEFAULT NULL,
  `card_master_id_7` int(10) unsigned DEFAULT NULL,
  `card_master_id_8` int(10) unsigned DEFAULT NULL,
  `card_master_id_9` int(10) unsigned DEFAULT NULL,
  `suit_master_id_1` int(10) unsigned DEFAULT NULL,
  `suit_master_id_2` int(10) unsigned DEFAULT NULL,
  `suit_master_id_3` int(10) unsigned DEFAULT NULL,
  `suit_master_id_4` int(10) unsigned DEFAULT NULL,
  `suit_master_id_5` int(10) unsigned DEFAULT NULL,
  `suit_master_id_6` int(10) unsigned DEFAULT NULL,
  `suit_master_id_7` int(10) unsigned DEFAULT NULL,
  `suit_master_id_8` int(10) unsigned DEFAULT NULL,
  `suit_master_id_9` int(10) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`user_live_deck_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_deck`
--

LOCK TABLES `user_live_deck` WRITE;
/*!40000 ALTER TABLE `user_live_deck` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_deck` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_difficulty`
--

DROP TABLE IF EXISTS `user_live_difficulty`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_difficulty` (
  `user_id` int(10) unsigned NOT NULL,
  `live_difficulty_id` int(10) unsigned NOT NULL,
  `max_score` int(10) unsigned NOT NULL,
  `max_combo` int(10) unsigned NOT NULL,
  `play_count` int(10) unsigned NOT NULL,
  `clear_count` int(10) unsigned NOT NULL,
  `cancel_count` int(10) unsigned NOT NULL,
  `not_cleared_count` int(10) unsigned NOT NULL,
  `is_full_combo` tinyint(1) NOT NULL,
  `cleared_difficulty_achievement_1` int(10) unsigned DEFAULT NULL,
  `cleared_difficulty_achievement_2` int(10) unsigned DEFAULT NULL,
  `cleared_difficulty_achievement_3` int(10) unsigned DEFAULT NULL,
  `enable_autoplay` tinyint(1) NOT NULL,
  `is_autoplay` tinyint(1) NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_difficulty_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_difficulty`
--

LOCK TABLES `user_live_difficulty` WRITE;
/*!40000 ALTER TABLE `user_live_difficulty` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_difficulty` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_mv`
--

DROP TABLE IF EXISTS `user_live_mv`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_mv` (
  `user_id` int(10) unsigned NOT NULL,
  `uniq_id` bigint(20) unsigned NOT NULL,
  `live_master_id` int(10) unsigned NOT NULL,
  `stage_master_id` int(10) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  `is_my_deck` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`uniq_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_mv`
--

LOCK TABLES `user_live_mv` WRITE;
/*!40000 ALTER TABLE `user_live_mv` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_mv` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_mv_deck`
--

DROP TABLE IF EXISTS `user_live_mv_deck`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_mv_deck` (
  `user_id` int(10) unsigned NOT NULL,
  `live_master_id` int(10) unsigned NOT NULL,
  `member_master_id_1` int(10) unsigned DEFAULT NULL,
  `member_master_id_2` int(10) unsigned DEFAULT NULL,
  `member_master_id_3` int(10) unsigned DEFAULT NULL,
  `member_master_id_4` int(10) unsigned DEFAULT NULL,
  `member_master_id_5` int(10) unsigned DEFAULT NULL,
  `member_master_id_6` int(10) unsigned DEFAULT NULL,
  `member_master_id_7` int(10) unsigned DEFAULT NULL,
  `member_master_id_8` int(10) unsigned DEFAULT NULL,
  `member_master_id_9` int(10) unsigned DEFAULT NULL,
  `suit_master_id_1` int(10) unsigned DEFAULT NULL,
  `suit_master_id_2` int(10) unsigned DEFAULT NULL,
  `suit_master_id_3` int(10) unsigned DEFAULT NULL,
  `suit_master_id_4` int(10) unsigned DEFAULT NULL,
  `suit_master_id_5` int(10) unsigned DEFAULT NULL,
  `suit_master_id_6` int(10) unsigned DEFAULT NULL,
  `suit_master_id_7` int(10) unsigned DEFAULT NULL,
  `suit_master_id_8` int(10) unsigned DEFAULT NULL,
  `suit_master_id_9` int(10) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_mv_deck`
--

LOCK TABLES `user_live_mv_deck` WRITE;
/*!40000 ALTER TABLE `user_live_mv_deck` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_mv_deck` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_partner`
--

DROP TABLE IF EXISTS `user_live_partner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_partner` (
  `user_id` int(10) unsigned NOT NULL,
  `time_ns` bigint(20) unsigned NOT NULL,
  `partner_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`time_ns`),
  KEY `ix_user_id_created_at` (`user_id`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_partner`
--

LOCK TABLES `user_live_partner` WRITE;
/*!40000 ALTER TABLE `user_live_partner` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_partner` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_party`
--

DROP TABLE IF EXISTS `user_live_party`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_party` (
  `user_id` int(10) unsigned NOT NULL,
  `party_id` int(10) unsigned NOT NULL,
  `user_live_deck_id` int(10) unsigned NOT NULL,
  `name` varchar(191) NOT NULL,
  `icon_master_id` int(10) unsigned NOT NULL,
  `card_master_id_1` int(10) unsigned DEFAULT NULL,
  `card_master_id_2` int(10) unsigned DEFAULT NULL,
  `card_master_id_3` int(10) unsigned DEFAULT NULL,
  `user_accessory_id_1` bigint(20) unsigned DEFAULT NULL,
  `user_accessory_id_2` bigint(20) unsigned DEFAULT NULL,
  `user_accessory_id_3` bigint(20) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`party_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_party`
--

LOCK TABLES `user_live_party` WRITE;
/*!40000 ALTER TABLE `user_live_party` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_party` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_play_summary`
--

DROP TABLE IF EXISTS `user_live_play_summary`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_play_summary` (
  `user_id` int(10) unsigned NOT NULL,
  `live_id` bigint(20) unsigned NOT NULL,
  `highest_appeal_at_once` int(10) unsigned NOT NULL,
  `highest_heal_at_once` int(10) unsigned NOT NULL,
  `highest_shield_at_once` int(10) unsigned NOT NULL,
  `heal_10` int(10) unsigned NOT NULL,
  `heal_20` int(10) unsigned NOT NULL,
  `heal_30` int(10) unsigned NOT NULL,
  `heal_50` int(10) unsigned NOT NULL,
  `heal_100` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_play_summary`
--

LOCK TABLES `user_live_play_summary` WRITE;
/*!40000 ALTER TABLE `user_live_play_summary` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_play_summary` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_live_skip_ticket`
--

DROP TABLE IF EXISTS `user_live_skip_ticket`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_live_skip_ticket` (
  `user_id` int(10) unsigned NOT NULL,
  `ticket_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`ticket_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_live_skip_ticket`
--

LOCK TABLES `user_live_skip_ticket` WRITE;
/*!40000 ALTER TABLE `user_live_skip_ticket` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_live_skip_ticket` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_bonus`
--

DROP TABLE IF EXISTS `user_login_bonus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_bonus` (
  `user_id` int(10) unsigned NOT NULL,
  `login_bonus_master_id` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `received_at` int(10) unsigned NOT NULL,
  `is_performance` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_bonus_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_bonus`
--

LOCK TABLES `user_login_bonus` WRITE;
/*!40000 ALTER TABLE `user_login_bonus` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_bonus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_bonus_beginner`
--

DROP TABLE IF EXISTS `user_login_bonus_beginner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_bonus_beginner` (
  `user_id` int(10) unsigned NOT NULL,
  `login_bonus_beginner_master_id` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `received_at` int(10) unsigned NOT NULL,
  `is_performance` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_bonus_beginner_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_bonus_beginner`
--

LOCK TABLES `user_login_bonus_beginner` WRITE;
/*!40000 ALTER TABLE `user_login_bonus_beginner` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_bonus_beginner` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_bonus_birthday`
--

DROP TABLE IF EXISTS `user_login_bonus_birthday`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_bonus_birthday` (
  `user_id` int(10) unsigned NOT NULL,
  `login_bonus_birthday_master_id` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `received_at` int(10) unsigned NOT NULL,
  `is_performance` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_bonus_birthday_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_bonus_birthday`
--

LOCK TABLES `user_login_bonus_birthday` WRITE;
/*!40000 ALTER TABLE `user_login_bonus_birthday` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_bonus_birthday` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_bonus_comeback`
--

DROP TABLE IF EXISTS `user_login_bonus_comeback`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_bonus_comeback` (
  `user_id` int(10) unsigned NOT NULL,
  `login_bonus_comeback_master_id` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `received_at` int(10) unsigned NOT NULL,
  `is_performance` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_bonus_comeback_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_bonus_comeback`
--

LOCK TABLES `user_login_bonus_comeback` WRITE;
/*!40000 ALTER TABLE `user_login_bonus_comeback` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_bonus_comeback` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_bonus_event_2d`
--

DROP TABLE IF EXISTS `user_login_bonus_event_2d`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_bonus_event_2d` (
  `user_id` int(10) unsigned NOT NULL,
  `login_bonus_event_2d_master_id` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `received_at` int(10) unsigned NOT NULL,
  `is_performance` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_bonus_event_2d_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_bonus_event_2d`
--

LOCK TABLES `user_login_bonus_event_2d` WRITE;
/*!40000 ALTER TABLE `user_login_bonus_event_2d` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_bonus_event_2d` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_bonus_event_3d`
--

DROP TABLE IF EXISTS `user_login_bonus_event_3d`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_bonus_event_3d` (
  `user_id` int(10) unsigned NOT NULL,
  `login_bonus_event_3d_master_id` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `received_at` int(10) unsigned NOT NULL,
  `is_performance` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_bonus_event_3d_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_bonus_event_3d`
--

LOCK TABLES `user_login_bonus_event_3d` WRITE;
/*!40000 ALTER TABLE `user_login_bonus_event_3d` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_bonus_event_3d` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_login_present`
--

DROP TABLE IF EXISTS `user_login_present`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_login_present` (
  `user_id` int(10) unsigned NOT NULL,
  `login_present_master_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`login_present_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_login_present`
--

LOCK TABLES `user_login_present` WRITE;
/*!40000 ALTER TABLE `user_login_present` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_login_present` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_managed_present`
--

DROP TABLE IF EXISTS `user_managed_present`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_managed_present` (
  `user_id` int(10) unsigned NOT NULL,
  `present_id` int(10) unsigned NOT NULL,
  `manager_id` int(10) unsigned DEFAULT NULL,
  `content_type` tinyint(3) unsigned NOT NULL,
  `content_id` int(10) unsigned DEFAULT NULL,
  `content_amount` int(10) unsigned NOT NULL,
  `route_type` tinyint(3) unsigned NOT NULL,
  `route_id` int(10) unsigned DEFAULT NULL,
  `param_server` text,
  `param_client` text,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`present_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_managed_present`
--

LOCK TABLES `user_managed_present` WRITE;
/*!40000 ALTER TABLE `user_managed_present` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_managed_present` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_member`
--

DROP TABLE IF EXISTS `user_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_member` (
  `user_id` int(10) unsigned NOT NULL,
  `member_master_id` int(10) unsigned NOT NULL,
  `suit_master_id` int(10) unsigned NOT NULL,
  `custom_background_master_id` int(10) unsigned NOT NULL,
  `love_point` int(10) unsigned NOT NULL,
  `love_point_limit` int(10) unsigned NOT NULL,
  `love_level` int(10) unsigned NOT NULL,
  `view_status` tinyint(3) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`member_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_member`
--

LOCK TABLES `user_member` WRITE;
/*!40000 ALTER TABLE `user_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_mission`
--

DROP TABLE IF EXISTS `user_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_mission` (
  `user_id` int(10) unsigned NOT NULL,
  `mission_m_id` int(10) unsigned NOT NULL,
  `mission_count` int(10) unsigned NOT NULL,
  `is_cleared` tinyint(1) NOT NULL,
  `is_received_reward` tinyint(1) NOT NULL,
  `new_expired_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`mission_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_mission`
--

LOCK TABLES `user_mission` WRITE;
/*!40000 ALTER TABLE `user_mission` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_mission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_mission_counter`
--

DROP TABLE IF EXISTS `user_mission_counter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_mission_counter` (
  `user_id` int(10) unsigned NOT NULL,
  `type` tinyint(3) unsigned NOT NULL,
  `count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_mission_counter`
--

LOCK TABLES `user_mission_counter` WRITE;
/*!40000 ALTER TABLE `user_mission_counter` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_mission_counter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_mission_counter_param1`
--

DROP TABLE IF EXISTS `user_mission_counter_param1`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_mission_counter_param1` (
  `user_id` int(10) unsigned NOT NULL,
  `type` tinyint(3) unsigned NOT NULL,
  `param1` int(10) unsigned NOT NULL,
  `count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`type`,`param1`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_mission_counter_param1`
--

LOCK TABLES `user_mission_counter_param1` WRITE;
/*!40000 ALTER TABLE `user_mission_counter_param1` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_mission_counter_param1` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_mission_counter_param2`
--

DROP TABLE IF EXISTS `user_mission_counter_param2`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_mission_counter_param2` (
  `user_id` int(10) unsigned NOT NULL,
  `type` tinyint(3) unsigned NOT NULL,
  `param1` int(10) unsigned NOT NULL,
  `param2` int(10) unsigned NOT NULL,
  `count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`type`,`param1`,`param2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_mission_counter_param2`
--

LOCK TABLES `user_mission_counter_param2` WRITE;
/*!40000 ALTER TABLE `user_mission_counter_param2` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_mission_counter_param2` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_mission_event_start_count`
--

DROP TABLE IF EXISTS `user_mission_event_start_count`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_mission_event_start_count` (
  `user_id` int(10) unsigned NOT NULL,
  `mission_m_id` int(10) unsigned NOT NULL,
  `mission_start_count` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`mission_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_mission_event_start_count`
--

LOCK TABLES `user_mission_event_start_count` WRITE;
/*!40000 ALTER TABLE `user_mission_event_start_count` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_mission_event_start_count` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_notice_last_fetch`
--

DROP TABLE IF EXISTS `user_notice_last_fetch`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_notice_last_fetch` (
  `user_id` int(10) unsigned NOT NULL,
  `last_fetched_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_notice_last_fetch`
--

LOCK TABLES `user_notice_last_fetch` WRITE;
/*!40000 ALTER TABLE `user_notice_last_fetch` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_notice_last_fetch` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_present`
--

DROP TABLE IF EXISTS `user_present`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_present` (
  `user_id` int(10) unsigned NOT NULL,
  `present_id` int(10) unsigned NOT NULL,
  `content_type` tinyint(3) unsigned NOT NULL,
  `content_id` int(10) unsigned DEFAULT NULL,
  `content_amount` int(10) unsigned NOT NULL,
  `route_type` tinyint(3) unsigned NOT NULL,
  `route_id` int(10) unsigned DEFAULT NULL,
  `param_server` text,
  `param_client` text,
  `posted_at` int(10) unsigned NOT NULL,
  `expired_at` int(10) unsigned DEFAULT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`present_id`),
  KEY `ix_user_id_posted_at` (`user_id`,`posted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_present`
--

LOCK TABLES `user_present` WRITE;
/*!40000 ALTER TABLE `user_present` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_present` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_present_received`
--

DROP TABLE IF EXISTS `user_present_received`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_present_received` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `present_id` int(10) unsigned NOT NULL,
  `content_type` tinyint(3) unsigned NOT NULL,
  `content_id` int(10) unsigned DEFAULT NULL,
  `content_amount` int(10) unsigned NOT NULL,
  `route_type` tinyint(3) unsigned NOT NULL,
  `route_id` int(10) unsigned DEFAULT NULL,
  `param_server` text COLLATE utf8mb4_bin,
  `param_client` text COLLATE utf8mb4_bin,
  `posted_at` int(10) unsigned NOT NULL,
  `history_created_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`),
  KEY `unique_idx_user_id_history_created_at` (`user_id`,`history_created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_present_received`
--

LOCK TABLES `user_present_received` WRITE;
/*!40000 ALTER TABLE `user_present_received` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_present_received` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_read_super_notice`
--

DROP TABLE IF EXISTS `user_read_super_notice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_read_super_notice` (
  `user_id` int(10) unsigned NOT NULL,
  `super_notice_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`super_notice_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_read_super_notice`
--

LOCK TABLES `user_read_super_notice` WRITE;
/*!40000 ALTER TABLE `user_read_super_notice` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_read_super_notice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_recovery_ap`
--

DROP TABLE IF EXISTS `user_recovery_ap`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_recovery_ap` (
  `user_id` int(10) unsigned NOT NULL,
  `recovery_ap_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`recovery_ap_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_recovery_ap`
--

LOCK TABLES `user_recovery_ap` WRITE;
/*!40000 ALTER TABLE `user_recovery_ap` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_recovery_ap` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_recovery_lp`
--

DROP TABLE IF EXISTS `user_recovery_lp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_recovery_lp` (
  `user_id` int(10) unsigned NOT NULL,
  `recovery_lp_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`recovery_lp_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_recovery_lp`
--

LOCK TABLES `user_recovery_lp` WRITE;
/*!40000 ALTER TABLE `user_recovery_lp` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_recovery_lp` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_reference_book`
--

DROP TABLE IF EXISTS `user_reference_book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_reference_book` (
  `user_id` int(10) unsigned NOT NULL,
  `reference_book_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`reference_book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_reference_book`
--

LOCK TABLES `user_reference_book` WRITE;
/*!40000 ALTER TABLE `user_reference_book` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_reference_book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_retry_gacha_tmp_result`
--

DROP TABLE IF EXISTS `user_retry_gacha_tmp_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_retry_gacha_tmp_result` (
  `user_id` int(10) unsigned NOT NULL,
  `gacha_draw_m_id` int(10) unsigned NOT NULL,
  `remain_retry_count` int(10) unsigned NOT NULL,
  `expire_at` int(10) unsigned NOT NULL,
  `card_id` int(10) unsigned NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`gacha_draw_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_retry_gacha_tmp_result`
--

LOCK TABLES `user_retry_gacha_tmp_result` WRITE;
/*!40000 ALTER TABLE `user_retry_gacha_tmp_result` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_retry_gacha_tmp_result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_review_request_process_flow`
--

DROP TABLE IF EXISTS `user_review_request_process_flow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_review_request_process_flow` (
  `user_id` int(10) unsigned NOT NULL,
  `review_request_flow_id` bigint(20) unsigned NOT NULL,
  `review_request_trigger_type` tinyint(3) unsigned NOT NULL,
  `review_request_status_type` tinyint(3) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`review_request_flow_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_review_request_process_flow`
--

LOCK TABLES `user_review_request_process_flow` WRITE;
/*!40000 ALTER TABLE `user_review_request_process_flow` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_review_request_process_flow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_rule_description`
--

DROP TABLE IF EXISTS `user_rule_description`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_rule_description` (
  `user_id` int(10) unsigned NOT NULL,
  `rule_description_master_id` int(10) unsigned NOT NULL,
  `display_status` tinyint(3) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`rule_description_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_rule_description`
--

LOCK TABLES `user_rule_description` WRITE;
/*!40000 ALTER TABLE `user_rule_description` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_rule_description` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_scene_tips`
--

DROP TABLE IF EXISTS `user_scene_tips`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_scene_tips` (
  `user_id` int(10) unsigned NOT NULL,
  `scene_tips_type` tinyint(3) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`scene_tips_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_scene_tips`
--

LOCK TABLES `user_scene_tips` WRITE;
/*!40000 ALTER TABLE `user_scene_tips` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_scene_tips` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_school_idol_festival_id_mission_data`
--

DROP TABLE IF EXISTS `user_school_idol_festival_id_mission_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_school_idol_festival_id_mission_data` (
  `user_id` int(10) unsigned NOT NULL,
  `version` int(10) unsigned DEFAULT NULL,
  `ll_user_id` int(10) unsigned DEFAULT NULL,
  `ll_user_json` text NOT NULL,
  `data_send_at` int(10) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_school_idol_festival_id_mission_data`
--

LOCK TABLES `user_school_idol_festival_id_mission_data` WRITE;
/*!40000 ALTER TABLE `user_school_idol_festival_id_mission_data` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_school_idol_festival_id_mission_data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_school_idol_festival_id_reward_dummy_json`
--

DROP TABLE IF EXISTS `user_school_idol_festival_id_reward_dummy_json`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_school_idol_festival_id_reward_dummy_json` (
  `user_id` int(10) unsigned NOT NULL,
  `ll_user_json` text NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_school_idol_festival_id_reward_dummy_json`
--

LOCK TABLES `user_school_idol_festival_id_reward_dummy_json` WRITE;
/*!40000 ALTER TABLE `user_school_idol_festival_id_reward_dummy_json` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_school_idol_festival_id_reward_dummy_json` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_school_idol_festival_id_reward_mission`
--

DROP TABLE IF EXISTS `user_school_idol_festival_id_reward_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_school_idol_festival_id_reward_mission` (
  `user_id` int(10) unsigned NOT NULL,
  `school_idol_festival_id_reward_mission_master_id` int(10) unsigned NOT NULL,
  `is_cleared` tinyint(1) NOT NULL,
  `count` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`school_idol_festival_id_reward_mission_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_school_idol_festival_id_reward_mission`
--

LOCK TABLES `user_school_idol_festival_id_reward_mission` WRITE;
/*!40000 ALTER TABLE `user_school_idol_festival_id_reward_mission` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_school_idol_festival_id_reward_mission` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_shop_event_exchange`
--

DROP TABLE IF EXISTS `user_shop_event_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_shop_event_exchange` (
  `user_id` int(10) unsigned NOT NULL,
  `event_exchange_master_id` int(10) unsigned NOT NULL,
  `event_exchange_content_master_id` int(10) unsigned NOT NULL,
  `exchange_count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_exchange_master_id`,`event_exchange_content_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_shop_event_exchange`
--

LOCK TABLES `user_shop_event_exchange` WRITE;
/*!40000 ALTER TABLE `user_shop_event_exchange` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_shop_event_exchange` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_shop_item_exchange`
--

DROP TABLE IF EXISTS `user_shop_item_exchange`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_shop_item_exchange` (
  `user_id` int(10) unsigned NOT NULL,
  `gacha_point_master_id` int(10) unsigned NOT NULL,
  `item_exchange_content_master_id` int(10) unsigned NOT NULL,
  `exchange_count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`gacha_point_master_id`,`item_exchange_content_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_shop_item_exchange`
--

LOCK TABLES `user_shop_item_exchange` WRITE;
/*!40000 ALTER TABLE `user_shop_item_exchange` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_shop_item_exchange` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_status`
--

DROP TABLE IF EXISTS `user_status`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_status` (
  `user_id` int(10) unsigned NOT NULL,
  `name` varchar(191) NOT NULL,
  `nickname` varchar(191) NOT NULL,
  `rank` int(10) unsigned NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `message` text NOT NULL,
  `common_key` varchar(191) NOT NULL,
  `auth_count` bigint(20) unsigned NOT NULL,
  `session_key` varchar(191) NOT NULL,
  `global_session_key` bigint(20) unsigned NOT NULL,
  `last_request_id` int(10) unsigned NOT NULL,
  `last_timestamp` bigint(20) unsigned NOT NULL,
  `last_login_at` int(10) unsigned NOT NULL,
  `device_token` varchar(191) NOT NULL,
  `device_name` varchar(191) NOT NULL,
  `max_friend_num` int(10) unsigned NOT NULL,
  `live_point_full_at` int(10) unsigned NOT NULL,
  `live_point_broken` int(10) unsigned NOT NULL,
  `activity_point_count` int(10) unsigned NOT NULL,
  `activity_point_reset_at` int(10) unsigned NOT NULL,
  `activity_point_payment_recovery_daily_count` int(10) unsigned NOT NULL,
  `activity_point_payment_recovery_daily_reset_at` int(10) unsigned NOT NULL,
  `game_money` int(10) unsigned NOT NULL,
  `card_exp` int(10) unsigned NOT NULL,
  `free_sns_coin` int(10) unsigned NOT NULL,
  `apple_sns_coin` int(10) unsigned NOT NULL,
  `google_sns_coin` int(10) unsigned NOT NULL,
  `birth_date` int(10) unsigned DEFAULT NULL,
  `birth_month` int(10) unsigned DEFAULT NULL,
  `birth_day` int(10) unsigned DEFAULT NULL,
  `latest_live_deck_id` int(10) unsigned NOT NULL,
  `main_lesson_deck_id` int(10) unsigned NOT NULL,
  `favorite_member_id` int(10) unsigned NOT NULL,
  `recommend_card_master_id` int(10) unsigned NOT NULL,
  `last_live_difficulty_id` int(10) unsigned NOT NULL,
  `emblem_id` int(10) unsigned NOT NULL,
  `tutorial_phase` tinyint(3) unsigned NOT NULL,
  `tutorial_end_at` int(10) unsigned NOT NULL,
  `present_counter` int(10) unsigned NOT NULL,
  `login_days` int(10) unsigned NOT NULL,
  `navi_tap_count` int(10) unsigned NOT NULL,
  `navi_tap_recover_at` int(10) unsigned NOT NULL,
  `is_auto_mode` tinyint(1) NOT NULL,
  `max_score_live_difficulty_master_id` int(10) unsigned DEFAULT NULL,
  `live_max_score` int(10) unsigned NOT NULL,
  `max_combo_live_difficulty_master_id` int(10) unsigned DEFAULT NULL,
  `live_max_combo` int(10) unsigned NOT NULL,
  `lesson_resume_status` tinyint(3) unsigned NOT NULL,
  `lp_magnification` int(10) unsigned NOT NULL,
  `login_bonus_received_at` int(10) unsigned NOT NULL,
  `login_bonus_next_check_at` int(10) unsigned NOT NULL,
  `accessory_box_additional` int(10) unsigned NOT NULL,
  `terms_of_use_version` int(10) unsigned NOT NULL,
  `has_caution` tinyint(1) NOT NULL,
  `last_get_friend_list_at` int(10) unsigned NOT NULL,
  `bootstrap_sifid_check_at` int(10) unsigned NOT NULL,
  `is_personal_notice` tinyint(1) NOT NULL,
  `gdpr_version` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `ix_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_status`
--

LOCK TABLES `user_status` WRITE;
/*!40000 ALTER TABLE `user_status` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_status` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_still_new`
--

DROP TABLE IF EXISTS `user_still_new`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_still_new` (
  `user_id` int(10) unsigned NOT NULL,
  `still_master_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`still_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_still_new`
--

LOCK TABLES `user_still_new` WRITE;
/*!40000 ALTER TABLE `user_still_new` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_still_new` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_story_main`
--

DROP TABLE IF EXISTS `user_story_main`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_story_main` (
  `user_id` int(10) unsigned NOT NULL,
  `story_main_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_main_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_story_main`
--

LOCK TABLES `user_story_main` WRITE;
/*!40000 ALTER TABLE `user_story_main` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_story_main` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_story_main_selected`
--

DROP TABLE IF EXISTS `user_story_main_selected`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_story_main_selected` (
  `user_id` int(10) unsigned NOT NULL,
  `story_main_cell_id` int(10) unsigned NOT NULL,
  `selected_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_main_cell_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_story_main_selected`
--

LOCK TABLES `user_story_main_selected` WRITE;
/*!40000 ALTER TABLE `user_story_main_selected` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_story_main_selected` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_story_member`
--

DROP TABLE IF EXISTS `user_story_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_story_member` (
  `user_id` int(10) unsigned NOT NULL,
  `story_member_master_id` int(10) unsigned NOT NULL,
  `acquired_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_member_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_story_member`
--

LOCK TABLES `user_story_member` WRITE;
/*!40000 ALTER TABLE `user_story_member` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_story_member` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_story_side`
--

DROP TABLE IF EXISTS `user_story_side`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_story_side` (
  `user_id` int(10) unsigned NOT NULL,
  `story_side_master_id` int(10) unsigned NOT NULL,
  `acquired_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_side_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_story_side`
--

LOCK TABLES `user_story_side` WRITE;
/*!40000 ALTER TABLE `user_story_side` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_story_side` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_suit`
--

DROP TABLE IF EXISTS `user_suit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_suit` (
  `user_id` int(10) unsigned NOT NULL,
  `suit_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`suit_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_suit`
--

LOCK TABLES `user_suit` WRITE;
/*!40000 ALTER TABLE `user_suit` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_suit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_time_difference`
--

DROP TABLE IF EXISTS `user_time_difference`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_time_difference` (
  `user_id` int(10) unsigned NOT NULL,
  `time_difference` int(11) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `ix_user_time_difference_time_difference` (`time_difference`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_time_difference`
--

LOCK TABLES `user_time_difference` WRITE;
/*!40000 ALTER TABLE `user_time_difference` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_time_difference` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_training_material`
--

DROP TABLE IF EXISTS `user_training_material`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_training_material` (
  `user_id` int(10) unsigned NOT NULL,
  `training_material_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`training_material_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_training_material`
--

LOCK TABLES `user_training_material` WRITE;
/*!40000 ALTER TABLE `user_training_material` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_training_material` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_unlock_scene`
--

DROP TABLE IF EXISTS `user_unlock_scene`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_unlock_scene` (
  `user_id` int(10) unsigned NOT NULL,
  `unlock_scene_type` tinyint(3) unsigned NOT NULL,
  `status` tinyint(3) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`unlock_scene_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_unlock_scene`
--

LOCK TABLES `user_unlock_scene` WRITE;
/*!40000 ALTER TABLE `user_unlock_scene` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_unlock_scene` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_virtual_access_area`
--

DROP TABLE IF EXISTS `user_virtual_access_area`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_virtual_access_area` (
  `user_id` int(10) unsigned NOT NULL,
  `country` varchar(191) DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_virtual_access_area`
--

LOCK TABLES `user_virtual_access_area` WRITE;
/*!40000 ALTER TABLE `user_virtual_access_area` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_virtual_access_area` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_virtual_gdpr`
--

DROP TABLE IF EXISTS `user_virtual_gdpr`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_virtual_gdpr` (
  `user_id` int(10) unsigned NOT NULL,
  `status` tinyint(3) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_virtual_gdpr`
--

LOCK TABLES `user_virtual_gdpr` WRITE;
/*!40000 ALTER TABLE `user_virtual_gdpr` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_virtual_gdpr` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_virtual_money_balance`
--

DROP TABLE IF EXISTS `user_virtual_money_balance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_virtual_money_balance` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `reason` tinyint(3) unsigned NOT NULL,
  `reason_master_id` int(10) unsigned DEFAULT NULL,
  `platform` tinyint(3) unsigned NOT NULL,
  `virtual_money_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `purchased_at` int(10) unsigned NOT NULL,
  `is_exclude_accounting` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_virtual_money_balance`
--

LOCK TABLES `user_virtual_money_balance` WRITE;
/*!40000 ALTER TABLE `user_virtual_money_balance` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_virtual_money_balance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_virtual_money_consume_history`
--

DROP TABLE IF EXISTS `user_virtual_money_consume_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_virtual_money_consume_history` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `reason` tinyint(3) unsigned NOT NULL,
  `reason_master_id` int(10) unsigned DEFAULT NULL,
  `virtual_money_master_id` int(10) unsigned NOT NULL,
  `free_amount` int(10) unsigned NOT NULL,
  `billing_amount` int(10) unsigned NOT NULL,
  `consume_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`),
  KEY `ix_created_at` (`created_at`),
  KEY `ix_consume_at` (`consume_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_virtual_money_consume_history`
--

LOCK TABLES `user_virtual_money_consume_history` WRITE;
/*!40000 ALTER TABLE `user_virtual_money_consume_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_virtual_money_consume_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_virtual_money_deposit_history`
--

DROP TABLE IF EXISTS `user_virtual_money_deposit_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_virtual_money_deposit_history` (
  `user_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `reason` tinyint(3) unsigned NOT NULL,
  `reason_master_id` int(10) unsigned DEFAULT NULL,
  `price` int(10) unsigned NOT NULL,
  `platform_price` varchar(191) DEFAULT NULL,
  `platform_formatted_price` varchar(191) DEFAULT NULL,
  `deposit_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`created_nsec`),
  KEY `ix_deposit_at_reason` (`deposit_at`,`reason`),
  KEY `ix_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_virtual_money_deposit_history`
--

LOCK TABLES `user_virtual_money_deposit_history` WRITE;
/*!40000 ALTER TABLE `user_virtual_money_deposit_history` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_virtual_money_deposit_history` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_virtual_time`
--

DROP TABLE IF EXISTS `user_virtual_time`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_virtual_time` (
  `user_id` int(10) unsigned NOT NULL,
  `offset` int(11) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_virtual_time`
--

LOCK TABLES `user_virtual_time` WRITE;
/*!40000 ALTER TABLE `user_virtual_time` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_virtual_time` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_voice`
--

DROP TABLE IF EXISTS `user_voice`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_voice` (
  `user_id` int(10) unsigned NOT NULL,
  `navi_voice_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`navi_voice_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_voice`
--

LOCK TABLES `user_voice` WRITE;
/*!40000 ALTER TABLE `user_voice` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_voice` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_weekly_mission`
--

DROP TABLE IF EXISTS `user_weekly_mission`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_weekly_mission` (
  `user_id` int(10) unsigned NOT NULL,
  `mission_m_id` int(10) unsigned NOT NULL,
  `mission_start_count` int(10) unsigned NOT NULL,
  `mission_count` int(10) unsigned NOT NULL,
  `is_cleared` tinyint(1) NOT NULL,
  `is_received_reward` tinyint(1) NOT NULL,
  `cleared_expired_at` int(10) unsigned DEFAULT NULL,
  `new_expired_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`mission_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_weekly_mission`
--

LOCK TABLES `user_weekly_mission` WRITE;
/*!40000 ALTER TABLE `user_weekly_mission` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_weekly_mission` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-05-06  9:52:38
