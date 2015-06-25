-- CREATE DATABASE go;
-- CREATE USER 'go'@'localhost' IDENTIFIED BY 'go';
-- GRANT ALL PRIVILEGES ON `go` . * TO 'go'@'localhost';

USE `go`;

DROP TABLE IF EXISTS `words`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `words` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `word` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX words_ind1 ON words (word);
/*!40101 SET character_set_client = @saved_cs_client */;
