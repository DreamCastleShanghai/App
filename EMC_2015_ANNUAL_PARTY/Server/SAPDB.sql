CREATE DATABASE  IF NOT EXISTS `SAP` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `SAP`;
-- MySQL dump 10.13  Distrib 5.7.9, for osx10.9 (x86_64)
--
-- Host: 127.0.0.1    Database: SAP
-- ------------------------------------------------------
-- Server version	5.6.24

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
-- Table structure for table `Session`
--

DROP TABLE IF EXISTS `Session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Session` (
  `SessionId` int(11) NOT NULL AUTO_INCREMENT,
  `SpeakerId` int(11) DEFAULT NULL,
  `SessionTitle` varchar(45) DEFAULT NULL,
  `Format` varchar(45) DEFAULT NULL,
  `Track` varchar(45) DEFAULT NULL,
  `Location` varchar(45) DEFAULT NULL,
  `StarTime` int(64) DEFAULT NULL,
  `EndTime` int(64) DEFAULT NULL,
  `SessionDescription` varchar(45) DEFAULT NULL,
  `Point` int(11) DEFAULT NULL,
  PRIMARY KEY (`SessionId`),
  KEY `SpeakerId_idx` (`SpeakerId`),
  CONSTRAINT `SpeakerId` FOREIGN KEY (`SpeakerId`) REFERENCES `Speaker` (`SpeakerId`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Session`
--

LOCK TABLES `Session` WRITE;
/*!40000 ALTER TABLE `Session` DISABLE KEYS */;
INSERT INTO `Session` VALUES (1,1,'session 001','format1','track1','loc1',1454144400,1454148000,'des for se1',10),(2,2,'session 002','format1','track2','loc3',1454151600,1454113200,'des for se2',20),(3,3,'session 003','format2','track3','loc2',1454151600,1454114750,'des for se3',30),(4,1,'session 004','format4','track1','loc2',1454144400,1454113200,'des for se4',25),(5,2,'session 005','format3','track2','loc3',1454151600,1454114750,'des for sp5',5),(6,3,'session 006','format2','track3','loc1',1454144400,1454113200,'des for sp6',15);
/*!40000 ALTER TABLE `Session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Speaker`
--

DROP TABLE IF EXISTS `Speaker`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Speaker` (
  `SpeakerId` int(11) NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(45) DEFAULT NULL,
  `LastName` varchar(45) DEFAULT NULL,
  `SpeakerTitle` varchar(45) DEFAULT NULL,
  `Company` varchar(45) DEFAULT NULL,
  `Conuntry` varchar(45) DEFAULT NULL,
  `Email` varchar(45) DEFAULT NULL,
  `SpeakerIcon` varchar(45) DEFAULT NULL,
  `SpeakerDescription` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`SpeakerId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Speaker`
--

LOCK TABLES `Speaker` WRITE;
/*!40000 ALTER TABLE `Speaker` DISABLE KEYS */;
INSERT INTO `Speaker` VALUES (1,'tang','taizong','dr','mec','shanghai','123@gmail.com',NULL,'des for sp1'),(2,'zhao','kuangyin','student','sony','ny','333@sina.com',NULL,'des for sp2'),(3,'li','shimin','teacher','google','beijing','zhins@emc.com',NULL,'des for sp3');
/*!40000 ALTER TABLE `Speaker` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Survey`
--

DROP TABLE IF EXISTS `Survey`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Survey` (
  `SurveyId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `SpeakerId` int(11) DEFAULT NULL,
  `SessionId` int(11) DEFAULT NULL,
  `SpeakerRank` int(11) DEFAULT NULL,
  `SessionRank` int(11) DEFAULT NULL,
  PRIMARY KEY (`SurveyId`),
  KEY `UserId_idx` (`UserId`),
  KEY `SpeakerId_idx` (`SpeakerId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Survey`
--

LOCK TABLES `Survey` WRITE;
/*!40000 ALTER TABLE `Survey` DISABLE KEYS */;
/*!40000 ALTER TABLE `Survey` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `UserId` int(11) NOT NULL AUTO_INCREMENT,
  `LoginName` varchar(45) DEFAULT NULL,
  `PassWord` varchar(45) DEFAULT NULL,
  `FirstName` varchar(45) DEFAULT NULL,
  `LastName` varchar(45) DEFAULT NULL,
  `Icon` varchar(45) DEFAULT NULL,
  `Rank` int(11) DEFAULT NULL,
  `Authority` int(11) DEFAULT NULL,
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
INSERT INTO `User` VALUES (1,'test001','001','zheng','min','icon1',10,3),(2,'test002','002','ma','li','icon2',15,2),(3,'test003','003','ding','junhui','icon3',20,1),(4,'test004','004','cao','shuai','icon4',25,0);
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User_Session_Relation`
--

DROP TABLE IF EXISTS `User_Session_Relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User_Session_Relation` (
  `relationid` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `SessionId` int(11) DEFAULT NULL,
  `LikeFlag` varchar(45) DEFAULT NULL,
  `CollectionFlag` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`relationid`),
  KEY `SessionId_idx` (`SessionId`),
  KEY `UserId_idx` (`UserId`),
  CONSTRAINT `SessionId` FOREIGN KEY (`SessionId`) REFERENCES `Session` (`SessionId`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User_Session_Relation`
--

LOCK TABLES `User_Session_Relation` WRITE;
/*!40000 ALTER TABLE `User_Session_Relation` DISABLE KEYS */;
INSERT INTO `User_Session_Relation` VALUES (1,1,1,'0','0'),(4,1,2,'0','1');
/*!40000 ALTER TABLE `User_Session_Relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Vote`
--

DROP TABLE IF EXISTS `Vote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Vote` (
  `VoteId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `VoteItemId` int(11) DEFAULT NULL,
  PRIMARY KEY (`VoteId`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Vote`
--

LOCK TABLES `Vote` WRITE;
/*!40000 ALTER TABLE `Vote` DISABLE KEYS */;
INSERT INTO `Vote` VALUES (1,1,1),(2,1,3),(3,3,2),(4,2,4),(5,2,1),(6,3,1),(7,4,1),(8,4,2),(9,4,5),(14,2,2);
/*!40000 ALTER TABLE `Vote` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `VoteItem`
--

DROP TABLE IF EXISTS `VoteItem`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `VoteItem` (
  `VoteItemId` int(11) NOT NULL,
  `VoteItemName` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`VoteItemId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `VoteItem`
--

LOCK TABLES `VoteItem` WRITE;
/*!40000 ALTER TABLE `VoteItem` DISABLE KEYS */;
INSERT INTO `VoteItem` VALUES (1,'vote item 001'),(2,'vote Item 002'),(3,'vote item 003'),(4,'vote item 004'),(5,'vote item 005');
/*!40000 ALTER TABLE `VoteItem` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tests`
--

DROP TABLE IF EXISTS `tests`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tests` (
  `id_tests` int(11) NOT NULL,
  `temp` int(11) DEFAULT NULL,
  PRIMARY KEY (`id_tests`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tests`
--

LOCK TABLES `tests` WRITE;
/*!40000 ALTER TABLE `tests` DISABLE KEYS */;
INSERT INTO `tests` VALUES (100,120),(200,NULL);
/*!40000 ALTER TABLE `tests` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-01-11  7:54:56
