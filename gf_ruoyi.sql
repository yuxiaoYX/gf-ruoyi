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

 Date: 15/04/2022 18:07:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(5) NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`config_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '参数配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);
INSERT INTO `sys_config` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', '初始化密码 123456', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);
INSERT INTO `sys_config` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', '深色主题theme-dark，浅色主题theme-light', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);
INSERT INTO `sys_config` VALUES (4, '账号自助-验证码开关', 'sys.account.captchaOnOff', 'true', 'Y', '是否开启验证码功能（true开启，false关闭）', '2022-04-13 15:06:52', '2022-04-14 10:23:32', NULL);
INSERT INTO `sys_config` VALUES (5, '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'false', 'Y', '是否开启注册用户功能（true开启，false关闭）', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '部门状态（0正常 1停用）',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 111 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '若依科技', 0, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '深圳总公司', 1, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (102, 100, '长沙分公司', 2, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (103, 101, '研发部门', 1, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (104, 101, '市场部门', 2, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (105, 101, '测试部门', 3, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (106, 101, '财务部门', 4, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (107, 101, '运维部门', 5, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (108, 102, '市场部门', 1, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (109, 102, '财务部门', 2, '若依', '15888888888', 'ry@qq.com', '0', '2022-04-10 17:49:50', '2022-04-10 17:49:50', NULL);
INSERT INTO `sys_dept` VALUES (110, 100, '测试', 1, '', '', '', '0', '2022-04-11 09:35:18', '2022-04-11 09:35:18', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'N' COMMENT '是否默认（0是 1否）',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', '0', '0', '显示菜单', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', '1', '0', '隐藏菜单', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', '0', '0', '正常状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', '1', '0', '停用状态', '2022-04-07 13:54:52', '2022-04-09 14:20:03', NULL);
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type123', '', 'info', '1', '0', '新增操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type123', '', 'info', '1', '0', '修改操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type123', '', 'danger', '1', '0', '删除操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type123', '', 'primary', '1', '0', '授权操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type123', '', 'warning', '1', '0', '导出操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type123', '', 'warning', '1', '0', '导入操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type123', '', 'danger', '1', '0', '强退操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type123', '', 'warning', '1', '0', '生成操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type123', '', 'danger', '1', '0', '清空操作', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_data` VALUES (27, 1, '是', 'Y', 'sys_yes_no', '', 'primary', '', '0', '系统默认是', '2022-04-13 16:00:13', '2022-04-13 16:00:24', NULL);
INSERT INTO `sys_dict_data` VALUES (28, 2, '否', 'N', 'sys_yes_no', '', 'danger', '', '0', '系统默认否', '2022-04-13 16:00:48', '2022-04-13 16:00:55', NULL);
INSERT INTO `sys_dict_data` VALUES (29, 1, '成功', '0', 'sys_common_status', '', 'primary', '', '0', '正常状态', '2022-04-14 16:02:03', '2022-04-14 16:02:03', NULL);
INSERT INTO `sys_dict_data` VALUES (30, 2, '失败', '1', 'sys_common_status', '', 'danger', '', '0', '停用状态', '2022-04-14 16:02:27', '2022-04-14 16:02:27', NULL);

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', '菜单状态列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', '系统开关列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type123', '0', '操作类型列表', '2022-04-07 13:54:52', '2022-04-09 14:44:54', NULL);
INSERT INTO `sys_dict_type` VALUES (10, '系统内置', 'sys_yes_no', '0', '', '2022-04-13 15:59:20', '2022-04-13 15:59:20', NULL);
INSERT INTO `sys_dict_type` VALUES (11, '系统状态', 'sys_common_status', '0', '日志，成功或失败', '2022-04-14 16:01:38', '2022-04-14 16:01:38', NULL);

-- ----------------------------
-- Table structure for sys_file
-- ----------------------------
DROP TABLE IF EXISTS `sys_file`;
CREATE TABLE `sys_file`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名称',
  `src` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '本地文件存储路径',
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'URL地址，可能为空',
  `file_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件类型',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '操作用户',
  `created_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文件列表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_file
-- ----------------------------
INSERT INTO `sys_file` VALUES (1, 'c6znge18hcbkv0i8ts.png', 'upload/20201110/c6znge18hcbkv0i8ts.png', '/upload/20201110/c6znge18hcbkv0i8ts.png', NULL, 1, '2020-11-10 22:22:45');
INSERT INTO `sys_file` VALUES (2, 'c6znjl77xng89hmlhg.png', 'upload/20201110/c6znjl77xng89hmlhg.png', '/upload/20201110/c6znjl77xng89hmlhg.png', NULL, 1, '2020-11-10 22:26:55');
INSERT INTO `sys_file` VALUES (3, 'c73j3oajspswdgqb1g.png', 'upload/20201115/c73j3oajspswdgqb1g.png', '/upload/20201115/c73j3oajspswdgqb1g.png', NULL, 10, '2020-11-15 11:48:44');
INSERT INTO `sys_file` VALUES (4, 'c73xron93lmguw4xrr.png', 'upload/20201115/c73xron93lmguw4xrr.png', '/upload/20201115/c73xron93lmguw4xrr.png', NULL, 12, '2020-11-15 23:18:21');
INSERT INTO `sys_file` VALUES (5, 'c75n1svi4mlsam1cdd.png', 'upload/20201117/c75n1svi4mlsam1cdd.png', '/upload/20201117/c75n1svi4mlsam1cdd.png', NULL, 15, '2020-11-17 23:19:41');
INSERT INTO `sys_file` VALUES (6, 'c7bhxgi4xbvc2owc1v.jpg', 'upload/20201124/c7bhxgi4xbvc2owc1v.jpg', '/upload/20201124/c7bhxgi4xbvc2owc1v.jpg', NULL, 16, '2020-11-24 20:34:55');
INSERT INTO `sys_file` VALUES (14, 'cj8wfq8pp6ngssc44t.jpg', 'upload\\20220413\\cj8wfq8pp6ngssc44t.jpg', NULL, '用户头像', 1, '2022-04-13 14:30:44');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (10, '', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '1', '用户名或密码错误！', '2022-04-15 10:22:45');
INSERT INTO `sys_login_log` VALUES (11, '', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '1', '用户名或密码错误！', '2022-04-15 10:23:13');
INSERT INTO `sys_login_log` VALUES (12, '', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '1', '用户名或密码错误！', '2022-04-15 10:24:07');
INSERT INTO `sys_login_log` VALUES (13, 'admin1', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '1', '用户名或密码错误！', '2022-04-15 10:29:53');
INSERT INTO `sys_login_log` VALUES (14, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '0', '登录成功', '2022-04-15 10:57:23');
INSERT INTO `sys_login_log` VALUES (15, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '0', '登录成功', '2022-04-15 10:59:27');
INSERT INTO `sys_login_log` VALUES (16, 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '0', '登录成功', '2022-04-15 11:07:05');
INSERT INTO `sys_login_log` VALUES (17, 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '0', '登录成功', '2022-04-15 11:56:13');
INSERT INTO `sys_login_log` VALUES (18, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '0', '登录成功', '2022-04-15 14:36:36');
INSERT INTO `sys_login_log` VALUES (19, '123456', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '1', '用户名或密码错误！', '2022-04-15 14:36:59');
INSERT INTO `sys_login_log` VALUES (20, 'admin1', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '0', '登录成功', '2022-04-15 14:37:17');
INSERT INTO `sys_login_log` VALUES (21, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Windows 7', '0', '登录成功', '2022-04-15 15:41:14');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
  `parent_id` int(10) NULL DEFAULT 0 COMMENT '父菜单ID',
  `sort` int(3) NULL DEFAULT NULL COMMENT '排序标记',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由地址',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '组件路径',
  `is_frame` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '是否为外链（0是 1否）',
  `is_cache` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '是否缓存（0缓存 1不缓存）',
  `query` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由参数',
  `menu_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单类型（M目录 C菜单 F按钮）',
  `visible` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '显示状态（0显示 1隐藏）',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单状态（0正常 1停用）',
  `perms` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `is_auth` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '是否验证（0是 1否）',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1062 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单表----id有问题，没有办法新建菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (2, '系统管理', 0, 2, 'system', NULL, '1', '0', NULL, 'M', '0', '0', NULL, '', 'system', '2021-04-05 14:05:04', NULL, NULL);
INSERT INTO `sys_menu` VALUES (3, '系统监控', 0, 3, 'monitor', NULL, '1', '0', NULL, 'M', '0', '0', NULL, '', 'monitor', '2021-04-05 14:05:04', '2021-05-28 14:58:44', NULL);
INSERT INTO `sys_menu` VALUES (4, '系统工具', 0, 4, 'tool', NULL, '1', '0', NULL, 'M', '0', '0', NULL, '', 'tool', '2021-04-05 14:05:04', '2021-07-09 16:36:28', NULL);
INSERT INTO `sys_menu` VALUES (5, '淘宝官网', 0, 5, 'https://www.taobao.com/', NULL, '0', '0', NULL, 'M', '0', '0', NULL, '', 'guide', '2021-04-05 14:05:04', '2021-08-21 09:30:37', NULL);
INSERT INTO `sys_menu` VALUES (100, '用户管理', 2, 1, 'user', 'system/user/index', '1', '0', NULL, 'C', '0', '0', 'system/user/getList', '', 'user', '2021-04-05 14:05:04', '2021-05-17 11:55:55', NULL);
INSERT INTO `sys_menu` VALUES (101, '角色管理', 2, 2, 'role', 'system/role/index', '1', '0', NULL, 'C', '0', '0', 'system/role/getList', '', 'peoples', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 2, 3, 'menu', 'system/menu/index', '1', '0', NULL, 'C', '0', '0', 'system/menu/getList', '', 'tree-table', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (103, '部门管理', 2, 10, 'dept', 'system/dept/index', '1', '0', NULL, 'C', '0', '0', 'system/dept/getList', '', 'tree', '2021-04-06 21:43:09', '2021-04-16 18:18:24', NULL);
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 2, 5, 'post', 'system/post/index', '1', '0', NULL, 'C', '0', '0', 'system/post/getList', '', 'post', '2021-04-06 21:43:09', '2021-04-16 18:16:14', '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 2, 6, 'dict', 'system/dict/index', '1', '0', NULL, 'C', '0', '0', 'system/dict/getList', '', 'dict', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (106, '参数设置', 2, 7, 'config', 'system/config/index', '1', '0', NULL, 'C', '0', '0', 'system/config/getList', '', 'edit', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (107, '通知公告', 2, 8, 'notice', 'system/notice/index', '1', '0', NULL, 'C', '0', '0', 'system/notice/getList', '', 'message', '2021-04-06 21:43:09', '2021-07-09 16:34:53', NULL);
INSERT INTO `sys_menu` VALUES (108, '日志管理', 2, 9, 'log', 'monitor/logininfor/index', '1', '0', NULL, 'M', '0', '0', NULL, '', 'log', '2021-04-06 21:43:09', '2022-04-14 15:39:03', NULL);
INSERT INTO `sys_menu` VALUES (109, '在线用户', 3, 1, 'online', 'monitor/online/index', '1', '0', NULL, 'C', '0', '0', 'monitor/online/getList', '', 'online', '2021-04-06 21:43:09', '2021-07-09 16:37:25', NULL);
INSERT INTO `sys_menu` VALUES (110, '定时任务', 3, 2, 'job', 'monitor/job/index', '1', '0', NULL, 'C', '0', '0', 'monitor/job/getList', '', 'job', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (111, '数据监控', 3, 3, 'druid', 'monitor/druid/index', '1', '0', NULL, 'C', '0', '0', 'monitor/druid/getList', '', 'druid', '2021-04-06 21:43:09', '2021-05-28 14:59:11', NULL);
INSERT INTO `sys_menu` VALUES (112, '服务监控', 3, 4, 'server', 'monitor/server/index', '1', '0', NULL, 'C', '0', '0', 'monitor/server/getList', '', 'server', '2021-04-06 21:43:09', '2021-05-28 14:59:21', NULL);
INSERT INTO `sys_menu` VALUES (113, '缓存监控', 3, 5, 'cache', 'monitor/cache/index', '1', '0', NULL, 'C', '0', '0', 'monitor/cache/getList', '', 'redis', '2021-04-06 21:43:09', '2021-05-28 14:59:25', NULL);
INSERT INTO `sys_menu` VALUES (114, '表单构建', 4, 1, 'build', 'tool/build/index', '1', '0', NULL, 'C', '0', '0', 'tool/build/getList', '', 'build', '2021-04-06 21:43:09', '2021-07-09 16:35:43', NULL);
INSERT INTO `sys_menu` VALUES (115, '代码生成', 4, 2, 'gen', 'tool/gen/index', '1', '0', NULL, 'C', '0', '0', 'tool/gen/getList', '', 'code', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (116, '系统接口', 4, 3, 'swagger', 'tool/swagger/index', '1', '0', NULL, 'C', '0', '0', 'tool/swagger/getList', '', 'swagger', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (500, '操作日志', 108, 1, 'operlog', 'monitor/operlog/index', '1', '0', NULL, 'C', '0', '0', 'monitor/operlog/getList', '', 'form', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (501, '登录日志', 108, 2, 'logininfor', 'monitor/logininfor/index', '1', '0', NULL, 'C', '0', '0', 'monitor/logininfor/getList', '', 'logininfor', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1001, '用户查询', 100, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1002, '用户新增', 100, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1003, '用户修改', 100, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1004, '用户删除', 100, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1005, '用户导出', 100, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1006, '用户导入', 100, 6, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/import', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1007, '重置密码', 100, 7, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/resetPwd', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1008, '角色查询', 101, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1009, '角色新增', 101, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1010, '角色修改', 101, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1011, '角色删除', 101, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1012, '角色导出', 101, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1013, '菜单查询', 102, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1014, '菜单新增', 102, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1015, '菜单修改', 102, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1016, '菜单删除', 102, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1017, '部门查询', 103, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1018, '部门新增', 103, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1019, '部门修改', 103, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1020, '部门删除', 103, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1021, '岗位查询', 104, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/getOne', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1022, '岗位新增', 104, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/create', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1023, '岗位修改', 104, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/update', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1024, '岗位删除', 104, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/delete', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1025, '岗位导出', 104, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/export', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1026, '字典查询', 105, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dict/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1027, '字典新增', 105, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dict/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1028, '字典修改', 105, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dict/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1029, '字典删除', 105, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dict/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1030, '字典导出', 105, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dict/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1031, '参数查询', 106, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1032, '参数新增', 106, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1033, '参数修改', 106, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1034, '参数删除', 106, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1035, '参数导出', 106, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1036, '公告查询', 107, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1037, '公告新增', 107, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1038, '公告修改', 107, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1039, '公告删除', 107, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1040, '操作查询', 500, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/operlog/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1041, '操作删除', 500, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/operlog/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1042, '日志导出', 500, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/operlog/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1043, '登录查询', 501, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/logininfor/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1044, '登录删除', 501, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/logininfor/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1045, '日志导出', 501, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/logininfor/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1046, '在线查询', 109, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/online/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1047, '批量强退', 109, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/online/batchLogout', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1048, '单条强退', 109, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/online/forceLogout', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1049, '任务查询', 110, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/getOne', '', '#', '2021-04-06 21:43:09', '2021-04-12 14:30:29', NULL);
INSERT INTO `sys_menu` VALUES (1050, '任务新增', 110, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/create', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1051, '任务修改', 110, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1052, '任务删除', 110, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1053, '状态修改', 110, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/changeStatus', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1054, '任务导出', 110, 7, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1055, '生成查询', 115, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/getOne', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1056, '生成修改', 115, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/update', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1057, '生成删除', 115, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/delete', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1058, '导入代码', 115, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/import', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1059, '预览代码', 115, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/preview', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1060, '生成代码', 115, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/code', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1061, '11111', 4, 6, '11111111', '', '1', '0', '', 'M', '0', '0', '', '', 'button', '2022-04-11 18:18:49', '2022-04-11 18:18:49', NULL);

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
  `json_result` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '返回参数',
  `status` int(1) NULL DEFAULT 0 COMMENT '操作状态（0正常 1异常）',
  `error_msg` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '错误消息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色名称',
  `role_sort` int(3) NULL DEFAULT NULL COMMENT '显示顺序',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色状态（0正常 1停用）',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', 1, '0', NULL, '2021-04-05 12:20:56', '2021-04-12 11:22:09', NULL);
INSERT INTO `sys_role` VALUES (2, '普通用户', 2, '0', '1', '2021-04-05 13:52:54', '2022-04-10 12:26:14', NULL);
INSERT INTO `sys_role` VALUES (3, '游客', 3, '0', '1111111111', '2021-04-05 13:52:54', '2022-04-10 12:26:14', NULL);
INSERT INTO `sys_role` VALUES (4, '回种清转', 7, '0', 'dolore elit pariatur veniam', '2022-04-04 11:21:07', '2022-04-10 12:26:03', NULL);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` int(10) NOT NULL COMMENT '角色ID',
  `menu_id` int(10) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和菜单关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (2, 2);
INSERT INTO `sys_role_menu` VALUES (2, 100);
INSERT INTO `sys_role_menu` VALUES (2, 101);
INSERT INTO `sys_role_menu` VALUES (2, 102);
INSERT INTO `sys_role_menu` VALUES (2, 105);
INSERT INTO `sys_role_menu` VALUES (2, 106);
INSERT INTO `sys_role_menu` VALUES (2, 107);
INSERT INTO `sys_role_menu` VALUES (2, 108);
INSERT INTO `sys_role_menu` VALUES (2, 500);
INSERT INTO `sys_role_menu` VALUES (2, 501);
INSERT INTO `sys_role_menu` VALUES (2, 1001);
INSERT INTO `sys_role_menu` VALUES (2, 1002);
INSERT INTO `sys_role_menu` VALUES (2, 1003);
INSERT INTO `sys_role_menu` VALUES (2, 1004);
INSERT INTO `sys_role_menu` VALUES (2, 1005);
INSERT INTO `sys_role_menu` VALUES (2, 1006);
INSERT INTO `sys_role_menu` VALUES (2, 1007);
INSERT INTO `sys_role_menu` VALUES (2, 1008);
INSERT INTO `sys_role_menu` VALUES (2, 1009);
INSERT INTO `sys_role_menu` VALUES (2, 1010);
INSERT INTO `sys_role_menu` VALUES (2, 1011);
INSERT INTO `sys_role_menu` VALUES (2, 1012);
INSERT INTO `sys_role_menu` VALUES (2, 1013);
INSERT INTO `sys_role_menu` VALUES (2, 1014);
INSERT INTO `sys_role_menu` VALUES (2, 1015);
INSERT INTO `sys_role_menu` VALUES (2, 1016);
INSERT INTO `sys_role_menu` VALUES (2, 1026);
INSERT INTO `sys_role_menu` VALUES (2, 1027);
INSERT INTO `sys_role_menu` VALUES (2, 1028);
INSERT INTO `sys_role_menu` VALUES (2, 1029);
INSERT INTO `sys_role_menu` VALUES (2, 1030);
INSERT INTO `sys_role_menu` VALUES (2, 1031);
INSERT INTO `sys_role_menu` VALUES (2, 1032);
INSERT INTO `sys_role_menu` VALUES (2, 1033);
INSERT INTO `sys_role_menu` VALUES (2, 1034);
INSERT INTO `sys_role_menu` VALUES (2, 1035);
INSERT INTO `sys_role_menu` VALUES (2, 1036);
INSERT INTO `sys_role_menu` VALUES (2, 1037);
INSERT INTO `sys_role_menu` VALUES (2, 1038);
INSERT INTO `sys_role_menu` VALUES (2, 1039);
INSERT INTO `sys_role_menu` VALUES (2, 1040);
INSERT INTO `sys_role_menu` VALUES (2, 1041);
INSERT INTO `sys_role_menu` VALUES (2, 1042);
INSERT INTO `sys_role_menu` VALUES (2, 1043);
INSERT INTO `sys_role_menu` VALUES (2, 1044);
INSERT INTO `sys_role_menu` VALUES (2, 1045);

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
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户状态（0正常 1停用）',
  `dept_id` bigint(20) NULL DEFAULT NULL COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '张三2', 'admin', '13306955991', NULL, '0', 100, NULL, '', NULL, '2016-12-10 18:20:54', '2022-04-13 14:30:44', NULL);
INSERT INTO `sys_user` VALUES (2, 'admin1', '熊子韬', '123456', '13383097147', NULL, '0', 103, NULL, '', NULL, '2001-05-22 16:40:18', '2022-04-11 17:34:29', NULL);
INSERT INTO `sys_user` VALUES (3, 'Song Ziyi', '宋子异', 'N57icHZE8H', '13380294485', NULL, '0', 100, NULL, '', NULL, '2011-04-29 23:10:39', '2014-07-10 11:58:56', NULL);
INSERT INTO `sys_user` VALUES (4, 'Jiang Xiuying', '蒋秀英', 'TEOWJ9MQWw', '13308985072', NULL, '0', 100, NULL, '', NULL, '2006-04-27 18:26:31', '2022-04-13 14:51:24', NULL);
INSERT INTO `sys_user` VALUES (5, 'Yu Lu', '于璐', 'nKfJlwcUWP', '13377340669', NULL, '0', 100, NULL, '', NULL, '2016-05-02 11:19:09', '2008-07-15 21:30:40', NULL);
INSERT INTO `sys_user` VALUES (6, 'Zhu Yunxi', '朱云熙', 'J7hOluztMq', '13374414913', NULL, '0', 100, NULL, '', NULL, '2008-02-18 00:16:40', '2018-12-06 12:56:27', NULL);
INSERT INTO `sys_user` VALUES (7, 'He Shihan', '贺詩涵', 'ZDy2TpkJzA', '13346790787', NULL, '0', 100, NULL, '', NULL, '2018-01-27 22:41:00', '2018-04-21 17:59:59', NULL);
INSERT INTO `sys_user` VALUES (8, 'Meng Rui', '孟睿', 'gyG1kFURfE', '13385848456', NULL, '0', 100, NULL, '', NULL, '2003-10-28 03:01:26', '2011-08-08 22:08:15', NULL);
INSERT INTO `sys_user` VALUES (9, 'Yao Lan', '姚岚', 'FH8NvNHZym', '13391029254', NULL, '0', 100, NULL, '', NULL, '2007-09-06 00:04:00', '2001-01-05 21:15:59', NULL);
INSERT INTO `sys_user` VALUES (10, 'Shi Shihan', '石詩涵', 'NbYJFTzvA2', '13337380665', NULL, '0', 100, NULL, '', NULL, '2021-07-07 14:28:18', '2001-04-15 00:27:39', NULL);
INSERT INTO `sys_user` VALUES (13, '高芳', '龙娟', '6V^7^nOEH', '18676426764', '', '0', 100, NULL, '', NULL, '2022-04-02 17:55:57', '2022-04-04 09:49:13', NULL);
INSERT INTO `sys_user` VALUES (14, '3333', '111', '123456', '15093278800', '', '0', 100, '', '', NULL, '2022-04-11 16:41:15', '2022-04-11 16:41:15', '2022-04-11 16:41:28');
INSERT INTO `sys_user` VALUES (15, '333', '111', '123456', '15093278800', '', '0', 100, '', '', NULL, '2022-04-11 16:43:10', '2022-04-11 16:43:10', '2022-04-11 16:43:16');

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `user_id` bigint(20) NULL DEFAULT NULL COMMENT '用户id',
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `os` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `ipaddr` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `created_at` datetime NULL DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`, `token`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 143 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (142, '46kkum0cjan6su35wmk1000mxmqmxsq1', 1, 'admin', 'Windows 7', '127.0.0.1', '内网IP', 'Chrome', '2022-04-15 15:41:14');
INSERT INTO `sys_user_online` VALUES (141, '1uzr24fcjaltu1m06h8200hy43pb5kc1', 2, 'admin1', 'Windows 7', '127.0.0.1', '内网IP', 'Chrome', '2022-04-15 14:37:17');
INSERT INTO `sys_user_online` VALUES (140, '46kkum0cjaltbgklcng1007338hc68xb', 1, 'admin', 'Windows 7', '127.0.0.1', '内网IP', 'Chrome', '2022-04-15 14:36:36');
INSERT INTO `sys_user_online` VALUES (139, '46kkum0cjaieirf2jn8100dldouewy7c', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-15 11:56:13');

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` int(10) NOT NULL COMMENT '用户ID',
  `role_id` int(10) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户和角色关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1);
INSERT INTO `sys_user_role` VALUES (2, 2);
INSERT INTO `sys_user_role` VALUES (2, 3);
INSERT INTO `sys_user_role` VALUES (2, 4);
INSERT INTO `sys_user_role` VALUES (5, 2);

SET FOREIGN_KEY_CHECKS = 1;
