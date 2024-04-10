Create Database If Not Exists b2_committer2 Character Set UTF8;
USE b2_committer2;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rollbacks
-- ----------------------------
DROP TABLE IF EXISTS `rollbacks`;
CREATE TABLE `rollbacks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `blockchain` varchar(32) NOT NULL COMMENT ' chain_name, eth',
  `event_id` bigint NOT NULL COMMENT ' event id',
  PRIMARY KEY (`id`),
  KEY `event_id_index` (`event_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sync_blocks
-- ----------------------------
DROP TABLE IF EXISTS `sync_blocks`;
CREATE TABLE `sync_blocks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `blockchain` varchar(32) NOT NULL COMMENT ' chain name',
  `miner` varchar(42) NOT NULL COMMENT ' miner',
  `block_time` bigint NOT NULL COMMENT ' block_time',
  `block_number` bigint NOT NULL COMMENT ' block_number',
  `block_hash` varchar(66) NOT NULL COMMENT ' block hash',
  `tx_count` bigint NOT NULL COMMENT ' tx count',
  `event_count` bigint NOT NULL COMMENT ' event count',
  `parent_hash` varchar(66) NOT NULL COMMENT ' parent hash',
  `status` varchar(32) NOT NULL COMMENT ' status',
  `check_count` bigint NOT NULL COMMENT ' check count',
  PRIMARY KEY (`id`),
  KEY `status_index` (`status`),
  KEY `tx_count_index` (`tx_count`),
  KEY `check_count_index` (`check_count`)
) ENGINE=InnoDB AUTO_INCREMENT=2923365 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sync_blocks_history
-- ----------------------------
DROP TABLE IF EXISTS `sync_blocks_history`;
CREATE TABLE `sync_blocks_history` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `blockchain` varchar(32) NOT NULL COMMENT ' block chain',
  `miner` varchar(42) NOT NULL COMMENT ' miner ',
  `block_time` bigint NOT NULL COMMENT ' block time',
  `block_number` bigint NOT NULL COMMENT ' block number',
  `block_hash` varchar(66) NOT NULL COMMENT ' block Hash',
  `tx_count` bigint NOT NULL COMMENT ' tx count ',
  `event_count` bigint NOT NULL COMMENT ' event count ',
  `parent_hash` varchar(66) NOT NULL COMMENT ' parent hash',
  `status` varchar(32) NOT NULL COMMENT ' status',
  `check_count` bigint NOT NULL COMMENT ' check count',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2792531 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sync_events
-- ----------------------------
DROP TABLE IF EXISTS `sync_events`;
CREATE TABLE `sync_events` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sync_block_id` bigint NOT NULL COMMENT ' sync_block_id',
  `blockchain` varchar(32) NOT NULL COMMENT ' blockchain',
  `block_time` bigint NOT NULL COMMENT ' block_time',
  `block_number` bigint NOT NULL COMMENT ' block_number',
  `block_hash` varchar(66) NOT NULL COMMENT ' block_hash',
  `block_log_indexed` bigint NOT NULL COMMENT ' block_log_indexed',
  `tx_index` bigint NOT NULL COMMENT ' tx_index',
  `tx_hash` varchar(66) NOT NULL COMMENT ' tx_hash',
  `event_name` varchar(32) NOT NULL COMMENT ' event_name',
  `event_hash` varchar(66) NOT NULL COMMENT ' event_hash',
  `contract_address` varchar(42) NOT NULL COMMENT ' contract_address',
  `data` json NOT NULL COMMENT ' data',
  `status` varchar(32) NOT NULL COMMENT ' status',
  `retry_count` bigint DEFAULT '0' COMMENT 'retry_count',
  PRIMARY KEY (`id`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1011299 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sync_events_history
-- ----------------------------
DROP TABLE IF EXISTS `sync_events_history`;
CREATE TABLE `sync_events_history` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sync_block_id` bigint NOT NULL COMMENT ' 同步区块ID',
  `blockchain` varchar(32) NOT NULL COMMENT ' 链名',
  `block_time` bigint NOT NULL COMMENT ' 区块时间',
  `block_number` bigint NOT NULL COMMENT ' 区块高度',
  `block_hash` varchar(66) NOT NULL COMMENT ' 区块Hash',
  `block_log_indexed` bigint NOT NULL COMMENT ' 日志index',
  `tx_index` bigint NOT NULL COMMENT ' 交易index',
  `tx_hash` varchar(66) NOT NULL COMMENT ' 交易Hash',
  `event_name` varchar(32) NOT NULL COMMENT ' 事件名称',
  `event_hash` varchar(66) NOT NULL COMMENT ' 事件Hash',
  `contract_address` varchar(42) NOT NULL COMMENT ' 合约地址',
  `data` json NOT NULL COMMENT ' 数据内容',
  `status` varchar(32) NOT NULL COMMENT ' 状态',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sync_tasks
-- ----------------------------
DROP TABLE IF EXISTS `sync_tasks`;
CREATE TABLE `sync_tasks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `blockchain` varchar(32) NOT NULL COMMENT ' 链名',
  `latest_block` bigint NOT NULL COMMENT ' 同步高度',
  `start_block` bigint NOT NULL COMMENT ' 同步开始区块高度',
  `end_block` bigint NOT NULL COMMENT ' 同步结束区块高度',
  `handle_num` bigint NOT NULL COMMENT ' 处理数',
  `contracts` text NOT NULL COMMENT ' 合约地址，多个用,分割',
  `status` varchar(32) NOT NULL COMMENT ' 状态',
  PRIMARY KEY (`id`),
  KEY `status_index` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=1000010 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

DROP TABLE IF EXISTS `proposal`;
CREATE TABLE `proposal` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `proposal_id` bigint NOT NULL,
    `state_root_hash` varchar(128) NOT NULL,
    `proof_root_hash` varchar(128) NOT NULL,
    `start_batch_num` bigint NOT NULL,
    `end_batch_num` bigint NOT NULL,
    `btc_tx_hash` varchar(128) ,
    `winner` varchar(128),
    `status` bigint NOT NULL DEFAULT 0,
    `generate_details_file` tinyint default 0 comment '0:no generate details file 1:already generate details file for uploading',
    `generate_details_file_time` datetime,
    `ar_tx_hash`  varchar(128),
    PRIMARY KEY (`id`),
    KEY `proposal_id_index` (`proposal_id`)
)
