Create Database If Not Exists b2_committer_op Character Set UTF8;
USE b2_committer_op;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `blob_info`;
CREATE TABLE `blob_info` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `block_number` bigint NOT NULL,
    `block_hash_hex` varchar(128) NOT NULL,
    `block_time` bigint NOT NULL COMMENT ' block_time',
    `blob_versioned_hash` varchar(128) NOT NULL,
    `blob_hashes_index` bigint NOT NULL,
    `blob_side_car_index` bigint NOT NULL,
    `blob_side_car_commitment` varchar(128) NOT NULL,
    `blob` MEDIUMTEXT NOT NULL,
    PRIMARY KEY (`id`),
    KEY `op_blob_index` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
