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

 Date: 24/04/2022 18:01:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_aa
-- ----------------------------
DROP TABLE IF EXISTS `sys_aa`;
CREATE TABLE `sys_aa`  (
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
-- Records of sys_aa
-- ----------------------------
INSERT INTO `sys_aa` VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);
INSERT INTO `sys_aa` VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', '初始化密码 123456', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);
INSERT INTO `sys_aa` VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', '深色主题theme-dark，浅色主题theme-light', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);
INSERT INTO `sys_aa` VALUES (4, '账号自助-验证码开关', 'sys.account.captchaOnOff', 'true', 'Y', '是否开启验证码功能（true开启，false关闭）', '2022-04-13 15:06:52', '2022-04-14 10:23:32', NULL);
INSERT INTO `sys_aa` VALUES (5, '账号自助-是否开启用户注册功能', 'sys.account.registerUser', 'false', 'Y', '是否开启注册用户功能（true开启，false关闭）', '2022-04-13 15:06:52', '2022-04-13 15:06:52', NULL);

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
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'info', '1', '0', '新增操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (19, 2, '删除', '2', 'sys_oper_type', '', 'danger', '1', '0', '删除操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (20, 3, '修改', '3', 'sys_oper_type', '', 'info', '1', '0', '修改操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (21, 4, '查询', '4', 'sys_oper_type', '', 'primary', '1', '0', '查询操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'warning', '1', '0', '导出操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', '1', '0', '导入操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', '1', '0', '强退操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', '1', '0', '生成操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', '1', '0', '清空操作', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
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
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', '操作类型列表', '2022-04-07 13:54:52', '2022-04-18 12:04:00', NULL);
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
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用目标字符串',
  `job_params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '3' COMMENT '计划执行错误策略（1立即执行 2执行一次 3放弃执行）',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '状态（0正常 1暂停）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`, `job_name`, `job_group`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------
INSERT INTO `sys_job` VALUES (1, '系统默认（无参）', 'DEFAULT', 'ryTask.ryNoParams', '', '0/10 * * * * ?', '3', '1', '', NULL, NULL, NULL);
INSERT INTO `sys_job` VALUES (2, '系统默认（有参）', 'DEFAULT', 'ryTask.ryParams(\'ry\')', '', '0/15 * * * * ?', '3', '1', '', NULL, NULL, NULL);
INSERT INTO `sys_job` VALUES (3, '系统默认（多参）', 'DEFAULT', 'ryTask.ryMultipleParams(\'ry\', true, 2000L, 316.50D, 100)', '', '0/20 * * * * ?', '3', '1', '', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log`  (
  `job_log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务日志ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务名称',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '任务组名',
  `invoke_target` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用目标字符串',
  `job_params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数',
  `job_message` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日志信息',
  `status` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '执行状态（0正常 1失败）',
  `exception_info` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '异常信息',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`job_log_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务调度日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (1, 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '0', '登录成功', '2022-04-22 10:55:59');
INSERT INTO `sys_login_log` VALUES (2, 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '0', '登录成功', '2022-04-23 10:20:13');
INSERT INTO `sys_login_log` VALUES (3, 'admin', '127.0.0.1', '内网IP', 'Edge', 'Windows 10', '0', '登录成功', '2022-04-24 09:04:11');

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
  `is_auth` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '是否验证（0是 1否）',
  `is_log` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '是否记录操作日志（0是 1否）',
  `icon` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1077 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单表----id有问题，没有办法新建菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (2, '系统管理', 0, 2, 'system', NULL, '1', '0', NULL, 'M', '0', '0', NULL, '0', '1', 'system', '2021-04-05 14:05:04', NULL, NULL);
INSERT INTO `sys_menu` VALUES (3, '系统监控', 0, 3, 'monitor', NULL, '1', '0', NULL, 'M', '0', '0', NULL, '0', '1', 'monitor', '2021-04-05 14:05:04', '2021-05-28 14:58:44', NULL);
INSERT INTO `sys_menu` VALUES (4, '系统工具', 0, 4, 'tool', NULL, '1', '0', NULL, 'M', '0', '0', NULL, '0', '1', 'tool', '2021-04-05 14:05:04', '2021-07-09 16:36:28', NULL);
INSERT INTO `sys_menu` VALUES (5, '淘宝官网', 0, 5, 'https://www.taobao.com/', NULL, '0', '0', NULL, 'M', '0', '0', NULL, '0', '1', 'guide', '2021-04-05 14:05:04', '2022-04-18 15:27:27', NULL);
INSERT INTO `sys_menu` VALUES (100, '用户管理', 2, 1, 'user', 'system/user/index', '1', '0', NULL, 'C', '0', '0', 'system/user/getList', '0', '1', 'user', '2021-04-05 14:05:04', '2021-05-17 11:55:55', NULL);
INSERT INTO `sys_menu` VALUES (101, '角色管理', 2, 2, 'role', 'system/role/index', '1', '0', NULL, 'C', '0', '0', 'system/role/getList', '0', '1', 'peoples', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 2, 3, 'menu', 'system/menu/index', '1', '0', NULL, 'C', '0', '0', 'system/menu/getList', '0', '1', 'tree-table', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (103, '部门管理', 2, 4, 'dept', 'system/dept/index', '1', '0', NULL, 'C', '0', '0', 'system/dept/getList', '0', '1', 'tree', '2021-04-06 21:43:09', '2021-04-16 18:18:24', NULL);
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 2, 5, 'post', 'system/post/index', '1', '0', NULL, 'C', '0', '0', 'system/post/getList', '0', '1', 'post', '2021-04-06 21:43:09', '2021-04-16 18:16:14', '2022-04-20 16:41:32');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 2, 6, 'dict', 'system/dict/index', '1', '0', NULL, 'C', '0', '0', 'system/dictType/getList', '0', '1', 'dict', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (106, '参数设置', 2, 7, 'config', 'system/config/index', '1', '0', NULL, 'C', '0', '0', 'system/config/getList', '0', '1', 'edit', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (107, '通知公告', 2, 8, 'notice', 'system/notice/index', '1', '0', NULL, 'C', '0', '0', 'system/notice/getList', '0', '1', 'message', '2021-04-06 21:43:09', '2021-07-09 16:34:53', '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (108, '日志管理', 3, 1, 'log', 'monitor/logininfor/index', '1', '0', NULL, 'M', '0', '0', NULL, '0', '1', 'log', '2021-04-06 21:43:09', '2022-04-14 15:39:03', NULL);
INSERT INTO `sys_menu` VALUES (109, '在线用户', 3, 2, 'online', 'monitor/online/index', '1', '0', NULL, 'C', '0', '0', 'monitor/online/getList', '0', '1', 'online', '2021-04-06 21:43:09', '2021-07-09 16:37:25', NULL);
INSERT INTO `sys_menu` VALUES (110, '定时任务', 3, 3, 'job', 'monitor/job/index', '1', '0', NULL, 'C', '1', '1', 'monitor/job/getList', '0', '1', 'job', '2021-04-06 21:43:09', '2022-04-20 16:40:58', NULL);
INSERT INTO `sys_menu` VALUES (111, '数据监控', 3, 4, 'druid', 'monitor/druid/index', '1', '0', NULL, 'C', '1', '1', 'monitor/druid/getList', '0', '1', 'druid', '2021-04-06 21:43:09', '2022-04-20 16:41:04', NULL);
INSERT INTO `sys_menu` VALUES (112, '服务监控', 3, 5, 'server', 'monitor/server/index', '1', '0', NULL, 'C', '1', '1', 'monitor/server/getList', '0', '1', 'server', '2021-04-06 21:43:09', '2022-04-20 16:41:09', NULL);
INSERT INTO `sys_menu` VALUES (113, '缓存监控', 3, 6, 'cache', 'monitor/cache/index', '1', '0', NULL, 'C', '1', '1', 'monitor/cache/getList', '0', '1', 'redis', '2021-04-06 21:43:09', '2022-04-20 16:41:12', NULL);
INSERT INTO `sys_menu` VALUES (114, '表单构建', 4, 1, 'build', 'tool/build/index', '1', '0', NULL, 'C', '0', '0', 'tool/build/getList', '0', '1', 'build', '2021-04-06 21:43:09', '2021-07-09 16:35:43', NULL);
INSERT INTO `sys_menu` VALUES (115, '代码生成', 4, 2, 'gen', 'tool/gen/index', '1', '0', NULL, 'C', '0', '0', 'tool/gen/getList', '0', '1', 'code', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (116, '系统接口', 4, 3, 'swagger', 'tool/swagger/index', '1', '0', NULL, 'C', '0', '0', 'tool/swagger/getList', '0', '1', 'swagger', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (500, '操作日志', 108, 1, 'operlog', 'monitor/operlog/index', '1', '0', NULL, 'C', '0', '0', 'monitor/operLog/getList', '0', '1', 'form', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (501, '登录日志', 108, 2, 'logininfor', 'monitor/logininfor/index', '1', '0', NULL, 'C', '0', '0', 'monitor/loginLog/getList', '0', '1', 'logininfor', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1001, '用户查询', 100, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1002, '用户新增', 100, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1003, '用户修改', 100, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1004, '用户删除', 100, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1005, '用户分配角色', 100, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/userSelectRole', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1006, '重置密码', 100, 6, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/resetPwd', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1007, '用户状态修改', 100, 7, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/changeStatus', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1008, '用户导出', 100, 8, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1009, '用户导入', 100, 9, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/import', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1010, '用户查询个人信息', 100, 10, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/profile', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1011, '用户修改个人信息', 100, 11, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/updateProfile', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1012, '用户修改个人密码', 100, 12, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/profile/updatePwd', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1013, '用户修改个人头像', 100, 13, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/user/profile/avatar', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1014, '角色查询', 101, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1015, '角色新增', 101, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1016, '角色修改', 101, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1017, '角色删除', 101, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1018, '角色导出', 101, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1019, '角色授权用户列表', 101, 6, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/allocatedList', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1020, '角色分配用户', 101, 7, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/roleSelectUser', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1021, '角色取消用户', 101, 8, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/role/cancel', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1022, '菜单查询', 102, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1023, '菜单新增', 102, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1024, '菜单修改', 102, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1025, '菜单删除', 102, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1026, '菜单下拉树', 102, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/menu/treeselect', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1027, '部门查询', 103, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1028, '部门新增', 103, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1029, '部门修改', 103, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1030, '部门删除', 103, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1031, '部门树结构', 103, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/dept/treeselect', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1032, '岗位查询', 104, 1, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1033, '岗位新增', 104, 2, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1034, '岗位修改', 104, 3, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1035, '岗位删除', 104, 4, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1036, '岗位导出', 104, 5, NULL, NULL, '1', '0', NULL, 'F', '0', '0', 'system/post/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1037, '字典查询', 105, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictType/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1038, '字典新增', 105, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictType/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1039, '字典修改', 105, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictType/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1040, '字典删除', 105, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictType/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1041, '字典导出', 105, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictType/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1042, '字典数据查询', 105, 6, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictData/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1043, '字典数据新增', 105, 7, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictData/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1044, '字典数据修改', 105, 8, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictData/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1045, '字典数据删除', 105, 9, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictData/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1046, '字典数据导出', 105, 10, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictData/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1047, '字典类型查询数据', 105, 11, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/dictData/type', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1048, '参数查询', 106, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1049, '参数新增', 106, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1050, '参数修改', 106, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1051, '参数删除', 106, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1052, '参数导出', 106, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1053, '参数键查询值', 106, 6, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/config/configKey', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1054, '公告查询', 107, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1055, '公告新增', 107, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1056, '公告修改', 107, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1057, '公告删除', 107, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'system/notice/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1058, '操作查询', 500, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/operLog/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1059, '操作删除', 500, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/operLog/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1060, '操作清空', 500, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/operLog/clean', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1061, '登录查询', 501, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/loginLog/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1062, '登录删除', 501, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/loginLog/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1063, '登录清空', 501, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/loginLog/clean', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1064, '批量强退', 109, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/online/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1065, '任务查询', 110, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/getOne', '0', '1', '#', '2021-04-06 21:43:09', '2021-04-12 14:30:29', NULL);
INSERT INTO `sys_menu` VALUES (1066, '任务新增', 110, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/create', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1067, '任务修改', 110, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1068, '任务删除', 110, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1069, '状态修改', 110, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/changeStatus', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1070, '任务导出', 110, 7, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'monitor/job/export', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1071, '生成查询', 115, 1, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/getOne', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1072, '生成修改', 115, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/update', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1073, '生成删除', 115, 3, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/delete', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1074, '导入代码', 115, 2, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/import', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1075, '预览代码', 115, 4, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/preview', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1076, '生成代码', 115, 5, '#', NULL, '1', '0', NULL, 'F', '0', '0', 'tool/gen/code', '0', '1', '#', '2021-04-06 21:43:09', NULL, NULL);

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2删除 3修改 4查询）',
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
) ENGINE = MyISAM AUTO_INCREMENT = 151 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (142, '46kkum0cjan6su35wmk1000mxmqmxsq1', 1, 'admin', 'Windows 7', '127.0.0.1', '内网IP', 'Chrome', '2022-04-15 15:41:14');
INSERT INTO `sys_user_online` VALUES (141, '1uzr24fcjaltu1m06h8200hy43pb5kc1', 2, 'admin1', 'Windows 7', '127.0.0.1', '内网IP', 'Chrome', '2022-04-15 14:37:17');
INSERT INTO `sys_user_online` VALUES (140, '46kkum0cjaltbgklcng1007338hc68xb', 1, 'admin', 'Windows 7', '127.0.0.1', '内网IP', 'Chrome', '2022-04-15 14:36:36');
INSERT INTO `sys_user_online` VALUES (139, '46kkum0cjaieirf2jn8100dldouewy7c', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-15 11:56:13');
INSERT INTO `sys_user_online` VALUES (143, '46kkum0cjd1ko9ykluw1005fe88i1ns8', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-18 11:22:57');
INSERT INTO `sys_user_online` VALUES (144, '46kkum0cjd1krtcltxo200dj7il7r96u', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-18 11:23:04');
INSERT INTO `sys_user_online` VALUES (145, '46kkum0cjd1meh79zd8100502resbtag', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-18 11:25:12');
INSERT INTO `sys_user_online` VALUES (146, '46kkum0cje0ejzukea8100j6pwadlscg', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-19 14:40:36');
INSERT INTO `sys_user_online` VALUES (147, '46kkum0cjes014k3sxs100nhd1wikfsy', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-20 12:18:09');
INSERT INTO `sys_user_online` VALUES (148, '46kkum0cjgfi7n8d7o8100786d69dibp', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-22 10:55:59');
INSERT INTO `sys_user_online` VALUES (149, '46kkum0cjh9dd7a9fjc100zs92hcz47n', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-23 10:20:13');
INSERT INTO `sys_user_online` VALUES (150, '46kkum0cji2dp5wuioc1006ftfo9k2fl', 1, 'admin', 'Windows 10', '127.0.0.1', '内网IP', 'Edge', '2022-04-24 09:04:11');

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

-- ----------------------------
-- Table structure for tools_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table`;
CREATE TABLE `tools_gen_table`  (
  `table_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表名称',
  `table_comment` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '表描述',
  `tpl_category` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'crud' COMMENT '使用的模板（crud单表操作 tree树表操作）',
  `object_name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据对象名称',
  `api_file` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'api文件夹',
  `controller_file` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'controller文件夹',
  `server_file` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'server文件夹',
  `function_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能描述',
  `function_author` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '生成功能作者',
  `options` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '其它生成选项',
  `remark` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`table_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 70 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of tools_gen_table
-- ----------------------------
INSERT INTO `tools_gen_table` VALUES (65, 'demo_gen_class', '代码生成关联测试表', 'crud', NULL, NULL, NULL, NULL, '代码生成关联测试', 'gfast', '', '', NULL, NULL, NULL);
INSERT INTO `tools_gen_table` VALUES (68, 'demo_gen_tree', '代码生成树形结构测试表', 'tree', NULL, NULL, NULL, NULL, '代码生成树形结构测试', 'gfast', '{\"treeCode\":\"id\",\"treeName\":\"demoName\",\"treeParentCode\":\"parentId\"}', '', NULL, NULL, NULL);
INSERT INTO `tools_gen_table` VALUES (69, 'demo_gen', '代码生成测试表', 'crud', NULL, NULL, NULL, NULL, '代码生成测试', 'gfast', '', '', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for tools_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `tools_gen_table_column`;
CREATE TABLE `tools_gen_table_column`  (
  `column_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `table_id` bigint(20) NULL DEFAULT NULL COMMENT '归属表编号',
  `column_name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列名称',
  `column_comment` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列描述',
  `column_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '列类型',
  `go_type` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Go类型',
  `go_field` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'Go字段名',
  `html_field` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'html字段名',
  `is_pk` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否主键（1是）',
  `is_increment` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否自增（1是）',
  `is_required` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否必填（1是）',
  `is_insert` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否为插入字段（1是）',
  `is_edit` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否编辑字段（1是）',
  `is_list` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否列表字段（1是）',
  `is_query` char(1) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '是否查询字段（1是）',
  `query_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'EQ' COMMENT '查询方式（等于、不等于、大于、小于、范围）',
  `html_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）',
  `dict_type` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `link_table_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表名',
  `link_table_class` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表类名',
  `link_table_package` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表包名',
  `link_label_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表键名',
  `link_label_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '关联表字段值',
  PRIMARY KEY (`column_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 720 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '代码生成业务表字段' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of tools_gen_table_column
-- ----------------------------
INSERT INTO `tools_gen_table_column` VALUES (666, 65, 'id', '分类id', 'int(10) unsigned', 'uint', 'Id', 'id', '1', '1', '0', '0', '0', '1', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (667, 65, 'class_name', '分类名', 'varchar(30)', 'string', 'ClassName', 'className', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (693, 68, 'id', 'ID', 'int(11) unsigned', 'uint', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (694, 68, 'parent_id', '父级ID', 'int(10) unsigned', 'uint', 'ParentId', 'parentId', '0', '0', '1', '1', '1', '0', '0', 'EQ', 'select', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (695, 68, 'demo_name', '姓名', 'varchar(20)', 'string', 'DemoName', 'demoName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (696, 68, 'demo_age', '年龄', 'int(10) unsigned', 'uint', 'DemoAge', 'demoAge', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 4, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (697, 68, 'classes', '班级', 'varchar(30)', 'string', 'Classes', 'classes', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 5, 'demo_gen_class', 'DemoGenClass', 'gfast/app/system', 'id', 'className');
INSERT INTO `tools_gen_table_column` VALUES (698, 68, 'demo_born', '出生年月', 'datetime', 'Time', 'DemoBorn', 'demoBorn', '0', '0', '', '1', '1', '1', '', 'EQ', 'datetime', '', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (699, 68, 'demo_gender', '性别', 'tinyint(3) unsigned', 'uint', 'DemoGender', 'demoGender', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_user_sex', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (700, 68, 'created_at', '创建日期', 'datetime', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '1', '1', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (701, 68, 'updated_at', '修改日期', 'datetime', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (702, 68, 'deleted_at', '删除日期', 'datetime', 'Time', 'DeletedAt', 'deletedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (703, 68, 'created_by', '创建人', 'bigint(20) unsigned', 'uint64', 'CreatedBy', 'createdBy', '0', '0', '0', '0', '0', '1', '0', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (704, 68, 'updated_by', '修改人', 'bigint(20) unsigned', 'uint64', 'UpdatedBy', 'updatedBy', '0', '0', '0', '0', '0', '1', '0', 'EQ', 'input', '', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (705, 68, 'demo_status', '状态', 'tinyint(4)', 'int', 'DemoStatus', 'demoStatus', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', 'sys_normal_disable', 13, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (706, 68, 'demo_cate', '分类', 'varchar(30)', 'string', 'DemoCate', 'demoCate', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'checkbox', 'cms_news_type', 14, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (707, 69, 'id', '', 'int(11) unsigned', 'uint', 'Id', 'id', '1', '1', '0', '0', '0', '0', '0', 'EQ', 'input', '', 1, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (708, 69, 'demo_name', '姓名', 'varchar(20)', 'string', 'DemoName', 'demoName', '0', '0', '1', '1', '1', '1', '1', 'LIKE', 'input', '', 2, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (709, 69, 'demo_age', '年龄', 'int(10) unsigned', 'uint', 'DemoAge', 'demoAge', '0', '0', '1', '1', '1', '1', '', 'EQ', 'input', '', 3, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (710, 69, 'classes', '班级', 'varchar(30)', 'string', 'Classes', 'classes', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', '', 4, 'demo_gen_class', 'DemoGenClass', 'gfast/app/system', 'id', 'className');
INSERT INTO `tools_gen_table_column` VALUES (711, 69, 'demo_born', '出生年月', 'datetime', 'Time', 'DemoBorn', 'demoBorn', '0', '0', '', '1', '1', '1', '', 'EQ', 'datetime', '', 5, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (712, 69, 'demo_gender', '性别', 'tinyint(3) unsigned', 'uint', 'DemoGender', 'demoGender', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'select', 'sys_user_sex', 6, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (713, 69, 'created_at', '创建日期', 'datetime', 'Time', 'CreatedAt', 'createdAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 7, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (714, 69, 'updated_at', '修改日期', 'datetime', 'Time', 'UpdatedAt', 'updatedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 8, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (715, 69, 'deleted_at', '删除日期', 'datetime', 'Time', 'DeletedAt', 'deletedAt', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'datetime', '', 9, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (716, 69, 'created_by', '创建人', 'bigint(20) unsigned', 'uint64', 'CreatedBy', 'createdBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 10, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (717, 69, 'updated_by', '修改人', 'bigint(20) unsigned', 'uint64', 'UpdatedBy', 'updatedBy', '0', '0', '0', '0', '0', '0', '0', 'EQ', 'input', '', 11, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (718, 69, 'demo_status', '状态', 'tinyint(4)', 'int', 'DemoStatus', 'demoStatus', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'radio', 'sys_normal_disable', 12, '', '', '', '', '');
INSERT INTO `tools_gen_table_column` VALUES (719, 69, 'demo_cate', '分类', 'varchar(30)', 'string', 'DemoCate', 'demoCate', '0', '0', '1', '1', '1', '1', '1', 'EQ', 'checkbox', 'cms_news_type', 13, '', '', '', '', '');

SET FOREIGN_KEY_CHECKS = 1;
