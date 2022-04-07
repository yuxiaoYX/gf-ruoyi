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

 Date: 07/04/2022 18:05:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', '0', '0', '性别男', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', '1', '0', '性别女', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', '1', '0', '性别未知', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', '0', '0', '显示菜单', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', '1', '0', '隐藏菜单', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', '0', '0', '正常状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', '1', '0', '停用状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', '0', '0', '正常状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', '1', '0', '停用状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', '0', '0', '默认分组', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', '1', '0', '系统分组', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', '0', '0', '系统默认是', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', '1', '0', '系统默认否', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', '0', '0', '通知', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', '1', '0', '公告', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', '0', '0', '正常状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', '1', '0', '关闭状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'info', '1', '0', '新增操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (19, 2, '修改', '2', 'sys_oper_type', '', 'info', '1', '0', '修改操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (20, 3, '删除', '3', 'sys_oper_type', '', 'danger', '1', '0', '删除操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (21, 4, '授权', '4', 'sys_oper_type', '', 'primary', '1', '0', '授权操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'warning', '1', '0', '导出操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', '1', '0', '导入操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', '1', '0', '强退操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', '1', '0', '生成操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', '1', '0', '清空操作', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (27, 1, '成功', '0', 'sys_common_status', '', 'primary', '1', '0', '正常状态', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (28, 2, '失败', '1', 'sys_common_status', '', 'danger', '1', '0', '停用状态', '2022-04-07 13:54:52', NULL, NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', '0', '用户性别列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (2, '菜单状态', 'sys_show_hide', '0', '菜单状态列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (3, '系统开关', 'sys_normal_disable', '0', '系统开关列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (4, '任务状态', 'sys_job_status', '0', '任务状态列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (5, '任务分组', 'sys_job_group', '0', '任务分组列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (6, '系统是否', 'sys_yes_no', '0', '系统是否列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (7, '通知类型', 'sys_notice_type', '0', '通知类型列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (8, '通知状态', 'sys_notice_status', '0', '通知状态列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (9, '操作类型', 'sys_oper_type', '0', '操作类型列表', '2022-04-07 13:54:52', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (10, '系统状态', 'sys_common_status', '0', '登录状态列表', '2022-04-07 13:54:52', NULL, NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1116 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单表----id有问题，没有办法新建菜单' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (2, '系统管理', 0, 2, 'system', '', '1', '0', NULL, 'M', '0', '0', NULL, '', 'system', '2021-04-05 14:05:04', NULL, NULL);
INSERT INTO `sys_menu` VALUES (3, '系统监控', 0, 3, 'monitor', '', '1', '0', NULL, 'M', '0', '0', '', '', 'monitor', '2021-04-05 14:05:04', '2021-05-28 14:58:44', NULL);
INSERT INTO `sys_menu` VALUES (4, '系统工具', 0, 4, 'tool', '', '1', '0', NULL, 'M', '0', '0', '', '', 'tool', '2021-04-05 14:05:04', '2021-07-09 16:36:28', NULL);
INSERT INTO `sys_menu` VALUES (5, '淘宝官网', 0, 5, 'https://www.taobao.com/', '', '0', '0', NULL, 'M', '0', '0', '', '', 'guide', '2021-04-05 14:05:04', '2021-08-21 09:30:37', NULL);
INSERT INTO `sys_menu` VALUES (100, '用户管理', 2, 1, 'user', 'system/user/index', '1', '0', NULL, 'C', '0', '0', 'api/user/getList', '', 'user', '2021-04-05 14:05:04', '2021-05-17 11:55:55', NULL);
INSERT INTO `sys_menu` VALUES (101, '角色管理', 2, 2, 'role', 'system/role/index', '1', '0', NULL, 'C', '0', '0', 'system:role:list', '', 'peoples', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (102, '菜单管理', 2, 3, 'menu', 'system/menu/index', '1', '0', NULL, 'C', '0', '0', 'system:menu:list', '', 'tree-table', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (103, '部门管理', 2, 10, 'dept', 'system/dept/index', '1', '0', NULL, 'C', '0', '0', 'system:dept:list', '', 'tree', '2021-04-06 21:43:09', '2021-04-16 18:18:24', '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (104, '岗位管理', 2, 5, 'post', 'system/post/index', '1', '0', NULL, 'C', '0', '0', 'system:post:list', '', 'post', '2021-04-06 21:43:09', '2021-04-16 18:16:14', '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (105, '字典管理', 2, 6, 'dict', 'system/dict/index', '1', '0', NULL, 'C', '0', '0', 'system:dict:list', '', 'dict', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (106, '参数设置', 2, 7, 'config', 'system/config/index', '1', '0', NULL, 'C', '0', '0', 'system:config:list', '', 'edit', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (107, '通知公告', 2, 8, 'notice', 'system/notice/index', '1', '0', NULL, 'C', '0', '0', 'system:notice:list', '', 'message', '2021-04-06 21:43:09', '2021-07-09 16:34:53', NULL);
INSERT INTO `sys_menu` VALUES (108, '日志管理', 2, 9, 'log', '', '1', '0', NULL, 'M', '0', '0', '', '', 'log', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (109, '在线用户', 3, 1, 'online', 'monitor/online/index', '1', '0', NULL, 'C', '0', '0', 'monitor:online:list', '', 'online', '2021-04-06 21:43:09', '2021-07-09 16:37:25', NULL);
INSERT INTO `sys_menu` VALUES (110, '定时任务', 3, 2, 'job', 'monitor/job/index', '1', '0', NULL, 'C', '0', '0', 'monitor:job:list', '', 'job', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (111, '数据监控', 3, 3, 'druid', 'monitor/druid/index', '1', '0', NULL, 'C', '0', '0', 'monitor:druid:list', '', 'druid', '2021-04-06 21:43:09', '2021-05-28 14:59:11', NULL);
INSERT INTO `sys_menu` VALUES (112, '服务监控', 3, 4, 'server', 'monitor/server/index', '1', '0', NULL, 'C', '0', '0', 'monitor:server:list', '', 'server', '2021-04-06 21:43:09', '2021-05-28 14:59:21', NULL);
INSERT INTO `sys_menu` VALUES (113, '缓存监控', 3, 5, 'cache', 'monitor/cache/index', '1', '0', NULL, 'C', '0', '0', 'monitor:cache:list', '', 'redis', '2021-04-06 21:43:09', '2021-05-28 14:59:25', NULL);
INSERT INTO `sys_menu` VALUES (114, '表单构建', 4, 1, 'build', 'tool/build/index', '1', '0', NULL, 'C', '0', '0', 'tool:build:list', '', 'build', '2021-04-06 21:43:09', '2021-07-09 16:35:43', NULL);
INSERT INTO `sys_menu` VALUES (115, '代码生成', 4, 2, 'gen', 'tool/gen/index', '1', '0', NULL, 'C', '0', '0', 'tool:gen:list', '', 'code', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (116, '系统接口', 4, 3, 'swagger', 'tool/swagger/index', '1', '0', NULL, 'C', '0', '0', 'tool:swagger:list', '', 'swagger', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (500, '操作日志', 108, 1, 'operlog', 'monitor/operlog/index', '1', '0', NULL, 'C', '0', '0', 'monitor:operlog:list', '', 'form', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (501, '登录日志', 108, 2, 'logininfor', 'monitor/logininfor/index', '1', '0', NULL, 'C', '0', '0', 'monitor:logininfor:list', '', 'logininfor', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1001, '用户查询', 100, 1, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1002, '用户新增', 100, 2, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1003, '用户修改', 100, 3, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1004, '用户删除', 100, 4, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1005, '用户导出', 100, 5, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1006, '用户导入', 100, 6, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:import', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1007, '重置密码', 100, 7, '', '', '1', '0', NULL, 'F', '0', '0', 'system:user:resetPwd', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1008, '角色查询', 101, 1, '', '', '1', '0', NULL, 'F', '0', '0', 'system:role:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1009, '角色新增', 101, 2, '', '', '1', '0', NULL, 'F', '0', '0', 'system:role:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1010, '角色修改', 101, 3, '', '', '1', '0', NULL, 'F', '0', '0', 'system:role:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1011, '角色删除', 101, 4, '', '', '1', '0', NULL, 'F', '0', '0', 'system:role:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1012, '角色导出', 101, 5, '', '', '1', '0', NULL, 'F', '0', '0', 'system:role:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1013, '菜单查询', 102, 1, '', '', '1', '0', NULL, 'F', '0', '0', 'system:menu:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1014, '菜单新增', 102, 2, '', '', '1', '0', NULL, 'F', '0', '0', 'system:menu:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1015, '菜单修改', 102, 3, '', '', '1', '0', NULL, 'F', '0', '0', 'system:menu:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1016, '菜单删除', 102, 4, '', '', '1', '0', NULL, 'F', '0', '0', 'system:menu:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1017, '部门查询', 103, 1, '', '', '1', '0', NULL, 'F', '0', '0', 'system:dept:query', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1018, '部门新增', 103, 2, '', '', '1', '0', NULL, 'F', '0', '0', 'system:dept:add', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1019, '部门修改', 103, 3, '', '', '1', '0', NULL, 'F', '0', '0', 'system:dept:edit', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1020, '部门删除', 103, 4, '', '', '1', '0', NULL, 'F', '0', '0', 'system:dept:remove', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1021, '岗位查询', 104, 1, '', '', '1', '0', NULL, 'F', '0', '0', 'system:post:query', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1022, '岗位新增', 104, 2, '', '', '1', '0', NULL, 'F', '0', '0', 'system:post:add', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1023, '岗位修改', 104, 3, '', '', '1', '0', NULL, 'F', '0', '0', 'system:post:edit', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1024, '岗位删除', 104, 4, '', '', '1', '0', NULL, 'F', '0', '0', 'system:post:remove', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1025, '岗位导出', 104, 5, '', '', '1', '0', NULL, 'F', '0', '0', 'system:post:export', '', '#', '2021-04-06 21:43:09', NULL, '2021-05-14 18:09:39');
INSERT INTO `sys_menu` VALUES (1026, '字典查询', 105, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:dict:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1027, '字典新增', 105, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:dict:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1028, '字典修改', 105, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:dict:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1029, '字典删除', 105, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:dict:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1030, '字典导出', 105, 5, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:dict:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1031, '参数查询', 106, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:config:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1032, '参数新增', 106, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:config:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1033, '参数修改', 106, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:config:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1034, '参数删除', 106, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:config:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1035, '参数导出', 106, 5, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:config:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1036, '公告查询', 107, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:notice:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1037, '公告新增', 107, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:notice:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1038, '公告修改', 107, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:notice:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1039, '公告删除', 107, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:notice:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1040, '操作查询', 500, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:operlog:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1041, '操作删除', 500, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:operlog:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1042, '日志导出', 500, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:operlog:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1043, '登录查询', 501, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:logininfor:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1044, '登录删除', 501, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:logininfor:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1045, '日志导出', 501, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:logininfor:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1046, '在线查询', 109, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:online:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1047, '批量强退', 109, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:online:batchLogout', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1048, '单条强退', 109, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:online:forceLogout', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1049, '任务查询', 110, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:job:query', '', '#', '2021-04-06 21:43:09', '2021-04-12 14:30:29', NULL);
INSERT INTO `sys_menu` VALUES (1050, '任务新增', 110, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:job:add', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1051, '任务修改', 110, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:job:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1052, '任务删除', 110, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:job:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1053, '状态修改', 110, 5, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:job:changeStatus', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1054, '任务导出', 110, 7, '#', '', '1', '0', NULL, 'F', '0', '0', 'monitor:job:export', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1055, '生成查询', 115, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'tool:gen:query', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1056, '生成修改', 115, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'tool:gen:edit', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1057, '生成删除', 115, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'tool:gen:remove', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1058, '导入代码', 115, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'tool:gen:import', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1059, '预览代码', 115, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'tool:gen:preview', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1060, '生成代码', 115, 5, '#', '', '1', '0', NULL, 'F', '0', '0', 'tool:gen:code', '', '#', '2021-04-06 21:43:09', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1066, 'api查询', 118, 1, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:apis:query', '', '#', '2021-04-18 17:31:58', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1067, 'api新增', 118, 2, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:apis:add', '', '#', '2021-04-18 17:31:58', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1068, 'api修改', 118, 3, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:apis:edit', '', '#', '2021-04-18 17:31:58', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1069, 'api删除', 118, 4, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:apis:remove', '', '#', '2021-04-18 17:31:58', NULL, NULL);
INSERT INTO `sys_menu` VALUES (1070, 'api导出', 118, 5, '#', '', '1', '0', NULL, 'F', '0', '0', 'system:apis:export', '', '#', '2021-04-19 18:05:26', NULL, NULL);

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
INSERT INTO `sys_role` VALUES (2, '普通用户', 2, '0', '1', '2021-04-05 13:52:54', '2021-04-29 16:09:26', NULL);
INSERT INTO `sys_role` VALUES (3, '一般用户', 3, '0', NULL, '2021-04-05 20:44:19', '2021-04-29 16:07:10', NULL);
INSERT INTO `sys_role` VALUES (4, '回种清转', 7, '0', 'dolore elit pariatur veniam', '2022-04-04 11:21:07', '2022-04-04 11:23:00', NULL);

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
INSERT INTO `sys_role_menu` VALUES (2, 3);
INSERT INTO `sys_role_menu` VALUES (2, 4);
INSERT INTO `sys_role_menu` VALUES (2, 100);
INSERT INTO `sys_role_menu` VALUES (2, 101);

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
  `dept_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `login_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '最后登录IP',
  `login_date` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '后台用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '张三', 'admin', '13306955996', NULL, '0', NULL, NULL, '', NULL, '2016-12-10 18:20:54', '2001-05-07 07:48:15', NULL);
INSERT INTO `sys_user` VALUES (2, 'admin1', '熊子韬', 'admin1', '13383097147', NULL, '0', NULL, NULL, '', NULL, '2001-05-22 16:40:18', '2019-12-26 04:10:40', NULL);
INSERT INTO `sys_user` VALUES (3, 'Song Ziyi', '宋子异', 'N57icHZE8H', '13380294485', NULL, '0', NULL, NULL, '', NULL, '2011-04-29 23:10:39', '2014-07-10 11:58:56', NULL);
INSERT INTO `sys_user` VALUES (4, 'Jiang Xiuying', '蒋秀英', 'TEOWJ9MQWw', '13308985072', NULL, '0', NULL, NULL, '', NULL, '2006-04-27 18:26:31', '2010-10-28 07:19:29', NULL);
INSERT INTO `sys_user` VALUES (5, 'Yu Lu', '于璐', 'nKfJlwcUWP', '13377340669', NULL, '0', NULL, NULL, '', NULL, '2016-05-02 11:19:09', '2008-07-15 21:30:40', NULL);
INSERT INTO `sys_user` VALUES (6, 'Zhu Yunxi', '朱云熙', 'J7hOluztMq', '13374414913', NULL, '0', NULL, NULL, '', NULL, '2008-02-18 00:16:40', '2018-12-06 12:56:27', NULL);
INSERT INTO `sys_user` VALUES (7, 'He Shihan', '贺詩涵', 'ZDy2TpkJzA', '13346790787', NULL, '0', NULL, NULL, '', NULL, '2018-01-27 22:41:00', '2018-04-21 17:59:59', NULL);
INSERT INTO `sys_user` VALUES (8, 'Meng Rui', '孟睿', 'gyG1kFURfE', '13385848456', NULL, '0', NULL, NULL, '', NULL, '2003-10-28 03:01:26', '2011-08-08 22:08:15', NULL);
INSERT INTO `sys_user` VALUES (9, 'Yao Lan', '姚岚', 'FH8NvNHZym', '13391029254', NULL, '0', NULL, NULL, '', NULL, '2007-09-06 00:04:00', '2001-01-05 21:15:59', NULL);
INSERT INTO `sys_user` VALUES (10, 'Shi Shihan', '石詩涵', 'NbYJFTzvA2', '13337380665', NULL, '0', NULL, NULL, '', NULL, '2021-07-07 14:28:18', '2001-04-15 00:27:39', NULL);
INSERT INTO `sys_user` VALUES (13, '高芳', '龙娟', '6V^7^nOEH', '18676426764', '', '0', '', '', '', NULL, '2022-04-02 17:55:57', '2022-04-04 09:49:13', NULL);

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
) ENGINE = MyISAM AUTO_INCREMENT = 67 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
INSERT INTO `sys_user_online` VALUES (26, '46kkum0cgzcu2bn8ewk1004euypfa6u4', 3, 'admin', '127.0.0.1', 'Chrome', 'Windows 7', '2022-01-07 18:03:08');
INSERT INTO `sys_user_online` VALUES (66, '46kkum0cj3rmnzzcao8100apt94cttci', 1, 'admin', '127.0.0.1', 'Edge', 'Windows 10', '2022-04-07 13:41:22');
INSERT INTO `sys_user_online` VALUES (54, '1uzr24fcj3mcr2yflqw200l1u7x70fa9', 2, 'admin1', '127.0.0.1', 'Chrome', 'Windows 7', '2022-04-07 09:33:20');

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
INSERT INTO `sys_user_role` VALUES (1, 2);
INSERT INTO `sys_user_role` VALUES (2, 2);
INSERT INTO `sys_user_role` VALUES (2, 3);

SET FOREIGN_KEY_CHECKS = 1;
