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
-- Table structure for table `Demo_Jam_Item`
--

DROP TABLE IF EXISTS `Demo_Jam_Item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Demo_Jam_Item` (
  `DemoJamItemId` int(11) NOT NULL AUTO_INCREMENT,
  `TeamName` varchar(45) DEFAULT NULL,
  `Resource` varchar(45) DEFAULT NULL,
  `Department` varchar(45) DEFAULT NULL,
  `Introduction` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`DemoJamItemId`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Demo_Jam_Item`
--

LOCK TABLES `Demo_Jam_Item` WRITE;
/*!40000 ALTER TABLE `Demo_Jam_Item` DISABLE KEYS */;
INSERT INTO `Demo_Jam_Item` VALUES (0,'Fully Automated Distribution Center','dj/dj1.jpg','department 001',NULL),(1,'Connected Cargo Inspection','dj/dj2.jpg','department 014',''),(2,'Intelligent Plaza','dj/dj3.jpg','lab of computer',''),(3,'WeChat Hiring as a Service','dj/dj4.jpg','agriculture group',''),(4,'Fault Diagnosis & Monitoring for Spindle on H','dj/dj5.jpg','engineering',''),(5,'Cyber Rumor Hunter','dj/dj6.jpg','office of hr',''),(6,'A Full Lifecycle IOT Exercise','dj/dj7.jpg','office of hr',''),(7,'EASI - Embedded Analytics','dj/dj8.jpg','office of hr','');
/*!40000 ALTER TABLE `Demo_Jam_Item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Demo_Jam_Vote`
--

DROP TABLE IF EXISTS `Demo_Jam_Vote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Demo_Jam_Vote` (
  `DemoJamVoteId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `DemoJamItemId` int(11) DEFAULT NULL,
  `VoteTime` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`DemoJamVoteId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Demo_Jam_Vote`
--

LOCK TABLES `Demo_Jam_Vote` WRITE;
/*!40000 ALTER TABLE `Demo_Jam_Vote` DISABLE KEYS */;
/*!40000 ALTER TABLE `Demo_Jam_Vote` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Egg_Hiking_Item`
--

DROP TABLE IF EXISTS `Egg_Hiking_Item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Egg_Hiking_Item` (
  `EggHikingItemId` int(11) NOT NULL AUTO_INCREMENT,
  `EggHikingTitle` varchar(45) DEFAULT NULL,
  `EggHikingDetail` varchar(256) DEFAULT NULL,
  `Resource` varchar(45) DEFAULT NULL,
  `EggHikingBG` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`EggHikingItemId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Egg_Hiking_Item`
--

LOCK TABLES `Egg_Hiking_Item` WRITE;
/*!40000 ALTER TABLE `Egg_Hiking_Item` DISABLE KEYS */;
INSERT INTO `Egg_Hiking_Item` VALUES (1,'Egg Hiking','We love our boss and support egg hiking.','eh/eh.png',NULL);
/*!40000 ALTER TABLE `Egg_Hiking_Item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Hiking_Vote`
--

DROP TABLE IF EXISTS `Hiking_Vote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Hiking_Vote` (
  `HikingId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `VoteFlag` char(1) DEFAULT NULL,
  `SubTime` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`HikingId`)
) ENGINE=InnoDB AUTO_INCREMENT=4003 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Hiking_Vote`
--

LOCK TABLES `Hiking_Vote` WRITE;
/*!40000 ALTER TABLE `Hiking_Vote` DISABLE KEYS */;
/*!40000 ALTER TABLE `Hiking_Vote` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Message`
--

DROP TABLE IF EXISTS `Message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Message` (
  `MessageId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `MessageDetail` varchar(256) DEFAULT NULL,
  `MessageTitle` varchar(45) DEFAULT NULL,
  `MessageTime` int(64) DEFAULT NULL,
  `MessageRealTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `MessageType` int(11) DEFAULT NULL,
  PRIMARY KEY (`MessageId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Message`
--

LOCK TABLES `Message` WRITE;
/*!40000 ALTER TABLE `Message` DISABLE KEYS */;
/*!40000 ALTER TABLE `Message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Picture_Wall`
--

DROP TABLE IF EXISTS `Picture_Wall`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Picture_Wall` (
  `PictureWallId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `Picture` varchar(45) DEFAULT NULL,
  `Category` varchar(45) DEFAULT NULL,
  `Comment` varchar(45) DEFAULT NULL,
  `SubTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `IsDelete` char(1) DEFAULT '0',
  PRIMARY KEY (`PictureWallId`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Picture_Wall`
--

LOCK TABLES `Picture_Wall` WRITE;
/*!40000 ALTER TABLE `Picture_Wall` DISABLE KEYS */;
/*!40000 ALTER TABLE `Picture_Wall` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Score_History`
--

DROP TABLE IF EXISTS `Score_History`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Score_History` (
  `ScoreHistoryId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `ScoreType` int(11) DEFAULT NULL,
  `Score` int(11) DEFAULT NULL,
  `ScoreDetail` varchar(45) DEFAULT NULL,
  `SubTime` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`ScoreHistoryId`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Score_History`
--

LOCK TABLES `Score_History` WRITE;
/*!40000 ALTER TABLE `Score_History` DISABLE KEYS */;
/*!40000 ALTER TABLE `Score_History` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Session`
--

DROP TABLE IF EXISTS `Session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Session` (
  `SessionId` int(11) NOT NULL AUTO_INCREMENT,
  `Title` varchar(512) DEFAULT NULL,
  `Format` varchar(512) DEFAULT NULL,
  `Track` varchar(512) DEFAULT NULL,
  `Location` varchar(45) DEFAULT NULL,
  `StartTime` int(64) DEFAULT NULL,
  `EndTime` int(64) DEFAULT NULL,
  `Description` varchar(2048) DEFAULT NULL,
  `Point` int(11) DEFAULT NULL,
  `Logo` varchar(45) DEFAULT NULL,
  `RealTime` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`SessionId`)
) ENGINE=InnoDB AUTO_INCREMENT=30825 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Session`
--

LOCK TABLES `Session` WRITE;
/*!40000 ALTER TABLE `Session` DISABLE KEYS */;
/*!40000 ALTER TABLE `Session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Session_Survey_Result`
--

DROP TABLE IF EXISTS `Session_Survey_Result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Session_Survey_Result` (
  `SurveyId` int(11) NOT NULL AUTO_INCREMENT,
  `SessionId` int(11) DEFAULT NULL,
  `UserId` int(11) DEFAULT NULL,
  `A1` int(11) DEFAULT NULL,
  `A2` int(11) DEFAULT NULL,
  `A3` int(11) DEFAULT NULL,
  `SubTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `IsCorrect` char(1) DEFAULT '0',
  PRIMARY KEY (`SurveyId`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Session_Survey_Result`
--

LOCK TABLES `Session_Survey_Result` WRITE;
/*!40000 ALTER TABLE `Session_Survey_Result` DISABLE KEYS */;
/*!40000 ALTER TABLE `Session_Survey_Result` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Speaker_Session_Relation`
--

DROP TABLE IF EXISTS `Speaker_Session_Relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Speaker_Session_Relation` (
  `RelationId` int(11) NOT NULL AUTO_INCREMENT,
  `SessionId` int(11) DEFAULT NULL,
  `SpeakerId` int(11) DEFAULT NULL,
  `Role` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`RelationId`)
) ENGINE=InnoDB AUTO_INCREMENT=190 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Speaker_Session_Relation`
--

LOCK TABLES `Speaker_Session_Relation` WRITE;
/*!40000 ALTER TABLE `Speaker_Session_Relation` DISABLE KEYS */;
/*!40000 ALTER TABLE `Speaker_Session_Relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Static_Res`
--

DROP TABLE IF EXISTS `Static_Res`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Static_Res` (
  `resId` int(11) NOT NULL AUTO_INCREMENT,
  `Resource` varchar(45) DEFAULT NULL,
  `ResLable` varchar(45) DEFAULT NULL,
  `ResType` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`resId`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Static_Res`
--

LOCK TABLES `Static_Res` WRITE;
/*!40000 ALTER TABLE `Static_Res` DISABLE KEYS */;
INSERT INTO `Static_Res` VALUES (1,'home/bar1.jpg',NULL,'bar'),(2,'home/bar2.jpg','ds','bar'),(3,'home/bar3.jpg','sc','bar'),(5,'home/bar4.jpg','','bar'),(6,'home/bar5.jpg','','bar'),(7,'home/bar6.jpg','dj','bar'),(8,'home/bar7.jpg','up','bar'),(9,'home/bar8.jpg','cpr','bar'),(10,'map/map1.jpg','1','map'),(11,'map/map2.jpg','2','map1'),(12,'map/map3.jpg','3','map1'),(13,'map/map4.jpg','4','map1'),(14,'map/map5.jpg','5','map1'),(15,'map/map6.jpg','6','map1'),(16,'map/map7.jpg','7','map1'),(17,'map/map8.jpg','8','map1'),(18,'me/info.jpg',NULL,'me');
/*!40000 ALTER TABLE `Static_Res` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Survey_Info`
--

DROP TABLE IF EXISTS `Survey_Info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Survey_Info` (
  `SurveyInfoId` int(11) NOT NULL AUTO_INCREMENT,
  `SessionId` int(11) DEFAULT NULL,
  `QContent1` varchar(512) DEFAULT NULL,
  `Q11` varchar(512) DEFAULT NULL,
  `Q12` varchar(512) DEFAULT NULL,
  `Q13` varchar(512) DEFAULT NULL,
  `Q14` varchar(512) DEFAULT NULL,
  `Answer1` int(11) DEFAULT NULL,
  `QContent2` varchar(512) DEFAULT NULL,
  `Q21` varchar(512) DEFAULT NULL,
  `Q22` varchar(512) DEFAULT NULL,
  `Q23` varchar(512) DEFAULT NULL,
  `Q24` varchar(512) DEFAULT NULL,
  `Answer2` int(11) DEFAULT NULL,
  PRIMARY KEY (`SurveyInfoId`),
  KEY `SpeakerId_idx` (`SessionId`)
) ENGINE=InnoDB AUTO_INCREMENT=99 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Survey_Info`
--

LOCK TABLES `Survey_Info` WRITE;
/*!40000 ALTER TABLE `Survey_Info` DISABLE KEYS */;
/*!40000 ALTER TABLE `Survey_Info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User`
--

DROP TABLE IF EXISTS `User`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User` (
  `UserId` int(11) NOT NULL AUTO_INCREMENT,
  `LoginName` varchar(256) DEFAULT NULL,
  `PassWord` varchar(45) DEFAULT NULL,
  `FirstName` varchar(45) DEFAULT NULL,
  `LastName` varchar(45) DEFAULT NULL,
  `Icon` varchar(45) DEFAULT NULL,
  `Department` varchar(45) DEFAULT NULL,
  `Title` varchar(45) DEFAULT NULL,
  `Status` varchar(45) DEFAULT NULL,
  `Category` varchar(45) DEFAULT NULL,
  `Company` varchar(45) DEFAULT NULL,
  `Email` varchar(45) DEFAULT NULL,
  `Conuntry` varchar(45) DEFAULT NULL,
  `Description` varchar(45) DEFAULT NULL,
  `Score` int(11) DEFAULT '0',
  `Authority` int(11) DEFAULT '-1',
  `DemoJamId1` int(11) DEFAULT '-1',
  `DemoJamId2` int(11) DEFAULT '-1',
  `VoiceVoteId1` int(11) DEFAULT '-1',
  `VoiceVoteId2` int(11) DEFAULT '-1',
  `EggVoted` char(1) DEFAULT '0',
  `GreenAmb` varchar(45) DEFAULT '0',
  `SubTime` int(64) DEFAULT NULL,
  `DeviceToken` varchar(64) DEFAULT NULL,
  `CanServeyTime` int(64) DEFAULT '0',
  `IsPrized` char(1) DEFAULT '0',
  PRIMARY KEY (`UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=812487 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User`
--

LOCK TABLES `User` WRITE;
/*!40000 ALTER TABLE `User` DISABLE KEYS */;
/*!40000 ALTER TABLE `User` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `User_Picture_Relation`
--

DROP TABLE IF EXISTS `User_Picture_Relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `User_Picture_Relation` (
  `RelationId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `PictureWallId` int(11) DEFAULT NULL,
  `LikeFlag` char(1) DEFAULT NULL,
  PRIMARY KEY (`RelationId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User_Picture_Relation`
--

LOCK TABLES `User_Picture_Relation` WRITE;
/*!40000 ALTER TABLE `User_Picture_Relation` DISABLE KEYS */;
/*!40000 ALTER TABLE `User_Picture_Relation` ENABLE KEYS */;
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
  `LikeFlag` char(1) DEFAULT '0',
  `CollectionFlag` char(1) DEFAULT '0',
  PRIMARY KEY (`relationid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `User_Session_Relation`
--

LOCK TABLES `User_Session_Relation` WRITE;
/*!40000 ALTER TABLE `User_Session_Relation` DISABLE KEYS */;
/*!40000 ALTER TABLE `User_Session_Relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Voice_Item`
--

DROP TABLE IF EXISTS `Voice_Item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Voice_Item` (
  `VoiceItemId` int(11) NOT NULL AUTO_INCREMENT,
  `VoicerName` varchar(45) DEFAULT NULL,
  `SongName` varchar(45) DEFAULT NULL,
  `VoicerPic` varchar(45) DEFAULT NULL,
  `VoicerDes` varchar(256) DEFAULT NULL,
  `VoicerBG` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`VoiceItemId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Voice_Item`
--

LOCK TABLES `Voice_Item` WRITE;
/*!40000 ALTER TABLE `Voice_Item` DISABLE KEYS */;
INSERT INTO `Voice_Item` VALUES (0,'哎呦不错哦','改变自己','sv/sv1.png','Shen Sharon; Ping Ursula; Chen lan','sv/bg1.jpg'),(1,'Wang Josie','Let it go','sv/sv2.png','Wang Josie','sv/bg2.jpg'),(2,'萨瓦滴咖','萨瓦迪卡','sv/sv3.png','Zhang Destiny; Song Samuel; Zhao Kun','sv/bg3.jpg');
/*!40000 ALTER TABLE `Voice_Item` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Voice_Vote`
--

DROP TABLE IF EXISTS `Voice_Vote`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Voice_Vote` (
  `VoiceVoteId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `VoiceItemId` int(11) DEFAULT NULL,
  `SubTime` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`VoiceVoteId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Voice_Vote`
--

LOCK TABLES `Voice_Vote` WRITE;
/*!40000 ALTER TABLE `Voice_Vote` DISABLE KEYS */;
/*!40000 ALTER TABLE `Voice_Vote` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Winner`
--

DROP TABLE IF EXISTS `Winner`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `Winner` (
  `WinnerId` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` int(11) DEFAULT NULL,
  `SubmitTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `WinnerType` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`WinnerId`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Winner`
--

LOCK TABLES `Winner` WRITE;
/*!40000 ALTER TABLE `Winner` DISABLE KEYS */;
/*!40000 ALTER TABLE `Winner` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-02-29  2:55:19
