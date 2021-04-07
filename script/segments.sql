/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80022
 Source Host           : localhost:3306
 Source Schema         : goleaf

 Target Server Type    : MySQL
 Target Server Version : 80022
 File Encoding         : 65001

 Date: 07/04/2021 15:29:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Database
-- ----------------------------
CREATE DATABASE IF NOT EXISTS goleaf;
CHARACTER SET = utf8mb4;
COLLATE = utf8mb4_general_ci;

-- ----------------------------
-- Table structure for segments
-- ----------------------------
DROP TABLE IF EXISTS `segments`;
CREATE TABLE `segments`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `max_id` bigint(0) UNSIGNED NOT NULL COMMENT '号段起始ID',
  `size` int(0) UNSIGNED NOT NULL COMMENT '号段大小',
  `biz_tag` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '业务标签',
  `description` varchar(1024) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '描述',
  `update_time` datetime(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '号段' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of segments
-- ----------------------------
INSERT INTO `segments` VALUES (1, 0, 10, 'mes', 'mes', '2021-04-07 15:29:00');
INSERT INTO `segments` VALUES (2, 0, 10, 'wms', 'wms', '2021-04-07 15:29:02');

SET FOREIGN_KEY_CHECKS = 1;
