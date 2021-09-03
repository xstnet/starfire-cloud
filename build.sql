/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : 127.0.0.1:3306
 Source Schema         : starfire_cloud

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 02/09/2021 23:25:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sf_file_extends
-- ----------------------------
DROP TABLE IF EXISTS `sf_file_extends`;
CREATE TABLE `sf_file_extends`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `file_id` int UNSIGNED NOT NULL DEFAULT 0,
  `extend` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `file_id`(`file_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件扩展表， 不确定的数据都放在这里， 和file表1对1' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sf_files
-- ----------------------------
DROP TABLE IF EXISTS `sf_files`;
CREATE TABLE `sf_files`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `short_md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件的内容分别从开始-中间-结尾截取一定的长度，生成一个简短的摘要， 用于尽可能快的匹配到是否已有该文件',
  `path` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件保存路径, 如 /path/日期/用户ID/文件名',
  `size` bigint UNSIGNED NOT NULL DEFAULT 0,
  `extend` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '扩展名，windows下最长可达200多个',
  `mime_type` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `own_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '拥有者ID，也就是第一个上传该文件的用户ID',
  `kind` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件类别：0：其他，1：文档，2：图片，3：音频，4：视频',
  `ref_count` mediumint UNSIGNED NOT NULL DEFAULT 0 COMMENT '引用计数',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `md5`(`md5`) USING BTREE,
  INDEX `short_md5`(`short_md5`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '上传文件表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sf_user_files
-- ----------------------------
DROP TABLE IF EXISTS `sf_user_files`;
CREATE TABLE `sf_user_files`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件层级， 上级ID',
  `user_id` int UNSIGNED NOT NULL DEFAULT 0,
  `file_id` int UNSIGNED NOT NULL DEFAULT 0,
  `is_dir` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否文件夹，0: 文件，1：文件夹',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件名称或文件夹名称',
  `is_delete` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '1：已删除，回收站数据读取=1',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间，可用做上传时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  INDEX `user_id_parentid_fileid`(`user_id`, `parent_id`, `file_id`) USING BTREE COMMENT '不能是unique,一个文件可以被上传多次'
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户关联文件表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sf_users
-- ----------------------------
DROP TABLE IF EXISTS `sf_users`;
CREATE TABLE `sf_users`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `used_space` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '已用存储空间（Unit: Byte）',
  `total_space` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '总存储空间（Unit: Byte）, 值为0时代表不限制',
  `created_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;

