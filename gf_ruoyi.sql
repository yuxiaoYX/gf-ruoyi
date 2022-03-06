/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : localhost:3306
 Source Schema         : gf_ruoyi

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 06/03/2022 18:02:54
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `role_key` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限字符',
  `role_sort` int(3) NULL DEFAULT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色状态；0:禁用，1:正常',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 'admin', 1, '1', NULL, '2021-04-05 12:20:56', '2021-04-12 11:22:09', NULL);
INSERT INTO `sys_role` VALUES (2, '普通用户', 'common', 2, '1', '1', '2021-04-05 13:52:54', '2021-04-29 16:09:26', NULL);
INSERT INTO `sys_role` VALUES (3, '一般用户', 'yiban', 3, '1', NULL, '2021-04-05 20:44:19', '2021-04-29 16:07:10', NULL);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户账号',
  `nick_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户昵称',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录密码',
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号码',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像地址',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户状态；0:禁用,1:正常',
  `dept_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '张三', 'admin', '13306955996', NULL, '1', NULL, NULL, '', NULL, '2016-12-10 18:20:54', '2001-05-07 07:48:15', NULL);
INSERT INTO `sys_user` VALUES (2, 'admin1', '熊子韬', 'admin1', '13383097147', NULL, '1', NULL, NULL, '', NULL, '2001-05-22 16:40:18', '2019-12-26 04:10:40', NULL);
INSERT INTO `sys_user` VALUES (3, 'Song Ziyi', '宋子异', 'N57icHZE8H', '13380294485', NULL, '1', NULL, NULL, '', NULL, '2011-04-29 23:10:39', '2014-07-10 11:58:56', NULL);
INSERT INTO `sys_user` VALUES (4, 'Jiang Xiuying', '蒋秀英', 'TEOWJ9MQWw', '13308985072', NULL, '1', NULL, NULL, '', NULL, '2006-04-27 18:26:31', '2010-10-28 07:19:29', NULL);
INSERT INTO `sys_user` VALUES (5, 'Yu Lu', '于璐', 'nKfJlwcUWP', '13377340669', NULL, '1', NULL, NULL, '', NULL, '2016-05-02 11:19:09', '2008-07-15 21:30:40', NULL);
INSERT INTO `sys_user` VALUES (6, 'Zhu Yunxi', '朱云熙', 'J7hOluztMq', '13374414913', NULL, '1', NULL, NULL, '', NULL, '2008-02-18 00:16:40', '2018-12-06 12:56:27', NULL);
INSERT INTO `sys_user` VALUES (7, 'He Shihan', '贺詩涵', 'ZDy2TpkJzA', '13346790787', NULL, '1', NULL, NULL, '', NULL, '2018-01-27 22:41:00', '2018-04-21 17:59:59', NULL);
INSERT INTO `sys_user` VALUES (8, 'Meng Rui', '孟睿', 'gyG1kFURfE', '13385848456', NULL, '1', NULL, NULL, '', NULL, '2003-10-28 03:01:26', '2011-08-08 22:08:15', NULL);
INSERT INTO `sys_user` VALUES (9, 'Yao Lan', '姚岚', 'FH8NvNHZym', '13391029254', NULL, '1', NULL, NULL, '', NULL, '2007-09-06 00:04:00', '2001-01-05 21:15:59', NULL);
INSERT INTO `sys_user` VALUES (10, 'Shi Shihan', '石詩涵', 'NbYJFTzvA2', '13337380665', NULL, '1', NULL, NULL, '', NULL, '2021-07-07 14:28:18', '2001-04-15 00:27:39', NULL);

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `user_id` int(10) NULL DEFAULT NULL COMMENT '用户id',
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  `created_at` datetime NULL DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (26, '46kkum0cgzcu2bn8ewk1004euypfa6u4', 3, 'admin', '127.0.0.1', 'Chrome', 'Windows 7', '2022-01-07 18:03:08');
INSERT INTO `sys_user_online` VALUES (30, '1uzr24fch3l6vmtamck100m13ax9edyb', 2, 'admin1', '127.0.0.1', 'Chrome', 'Windows 7', '2022-01-12 17:26:41');
INSERT INTO `sys_user_online` VALUES (34, 'sj60UMnaZHRDKfREWq5JXn1TB+l6RWZH0HV8ME5KQ89WNbXYEZ4SBIB0wt3lU6PK', 0, 'admin1', '127.0.0.1', 'Chrome', 'Windows 7', '2022-01-14 17:44:56');
INSERT INTO `sys_user_online` VALUES (35, 'iw0FIPDuR1RJQVt9q9RvcMHgp9KVi3C6D517bevaVni4sVboNP5YYnWAQRrsGqHp', 0, 'admin1', '127.0.0.1', 'Chrome', 'Windows 7', '2022-01-14 17:47:14');

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` int(10) NOT NULL COMMENT '用户ID',
  `role_id` int(10) NOT NULL COMMENT '角色ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (2, 3);
INSERT INTO `sys_user_role` VALUES (2, 2);

SET FOREIGN_KEY_CHECKS = 1;
