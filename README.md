# Hello World on Go

Using Go Lang for testing rest api with Gorilla, querying to mysql database using Gorm.

### Prerequisites

Install mysql table with below DDL

```
CREATE DATABASE `test`;
DROP TABLE IF EXISTS `test`;
CREATE TABLE `test` (
  `field1` varchar(40) NOT NULL,
  PRIMARY KEY (`field1`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of test
-- ----------------------------
INSERT INTO `test` VALUES ('aaaaaa');
INSERT INTO `test` VALUES ('bbbbb');
INSERT INTO `test` VALUES ('ccccccc');
```