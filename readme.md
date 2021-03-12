# Laravel Go

Laravel 的 Golang 开源框架，包含了 GORM、Echo 等，当前配置里的 APP_POLICE_URL 为企业微信机器人地址

### 数据表依赖

```
# 演示
CREATE TABLE `demos` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '标题',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='演示';

# 系统日志
CREATE TABLE `system_logs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `request` text COMMENT '请求参数',
  `message` mediumtext COMMENT '日志内容',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统日志';
```