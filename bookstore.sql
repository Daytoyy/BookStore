
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` float(100, 0) NOT NULL,
  `sales` int(100) NOT NULL,
  `stock` int(100) NOT NULL,
  `img_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 44 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO `books` VALUES (1, '解忧杂货店', '东野圭吾', 27, 100, 100, 'static/img/解忧杂货铺.png');
INSERT INTO `books` VALUES (2, '边城', '沈从文', 23, 101, 99, 'static/img/边城.png');
INSERT INTO `books` VALUES (3, '中国哲学史', '冯友兰', 44, 101, 99, 'static/img/中国哲学史.jpg');
INSERT INTO `books` VALUES (4, '忽然七日', ' 劳伦', 19, 104, 96, 'static/img/忽然七日.png');
INSERT INTO `books` VALUES (5, '苏东坡传', '林语堂', 19, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (6, '百年孤独', '马尔克斯', 29, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (7, '扶桑', '严歌苓', 20, 102, 98, 'static/img/default.jpg');
INSERT INTO `books` VALUES (8, '给孩子的诗', '北岛', 22, 102, 98, 'static/img/default.jpg');
INSERT INTO `books` VALUES (9, '为奴十二年', '所罗门', 16, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (10, '平凡的世界', '路遥', 55, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (11, '悟空传', '今何在', 14, 103, 97, 'static/img/default.jpg');
INSERT INTO `books` VALUES (12, '硬派健身', '斌卡', 31, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (13, '从晚清到民国', '唐德刚', 40, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (14, '三体', '刘慈欣', 56, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (15, '看见', '柴静', 19, 102, 98, 'static/img/default.jpg');
INSERT INTO `books` VALUES (16, '活着', '余华', 11, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (17, '小王子', '安托万', 19, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (18, '我们仨', '杨绛', 11, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (19, '生命不息,折腾不止', '罗永浩', 25, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (20, '皮囊', '蔡崇达', 24, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (21, '恰到好处的幸福', '毕淑敏', 16, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (22, '大数据预测', '埃里克', 37, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (23, '人月神话', '布鲁克斯', 56, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (24, 'C语言入门经典', '霍尔顿', 45, 101, 99, 'static/img/default.jpg');
INSERT INTO `books` VALUES (25, '数学之美', '吴军', 30, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (26, 'Java编程思想', '埃史尔', 70, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (27, '设计模式之禅', '秦小波', 20, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (28, '图解机器学习', '杉山将', 34, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (29, '艾伦图灵传', '安德鲁', 47, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (30, '教父', '马里奥普佐', 29, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (40, 'Go语言学习笔记', '雨痕', 51, 100, 33, 'static/img/default.jpg');
INSERT INTO `books` VALUES (43, 'go语言基础', 'l1ng14', 30, 10, 111, 'static/img/default.jpg');
INSERT INTO `books` VALUES (44, '毕业', '东野奎吾', 15.8, 0, 200, 'static/img/default.jpg');
INSERT INTO `books` VALUES (45, '沉睡的森林', '东野奎吾', 15, 0, 200, 'static/img/default.jpg');


-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `PASSWORD` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`, `username`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE,
  INDEX `id`(`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions`  (
  `session_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`session_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `sessions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for carts
-- ----------------------------
DROP TABLE IF EXISTS `carts`;
CREATE TABLE `carts`  (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11, 2) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `carts_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for cart_items
-- ----------------------------
DROP TABLE IF EXISTS `cart_items`;
CREATE TABLE `cart_items`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `book_id` int(11) NOT NULL,
  `cart_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `book_id`(`book_id`) USING BTREE,
  INDEX `cart_id`(`cart_id`) USING BTREE,
  CONSTRAINT `cart_items_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `cart_items_ibfk_2` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 77 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime(0) NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11, 2) NOT NULL,
  `state` int(11) NOT NULL,
  `user_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;


-- ----------------------------
-- Table structure for order_items
-- ----------------------------
DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `COUNT` int(11) NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `author` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` double(11, 2) NOT NULL,
  `img_path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `order_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE,
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;






