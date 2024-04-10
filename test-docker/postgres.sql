CREATE DATABASE b2_committer WITH ENCODING ='UTF8';
-- Switch to the newly created database
\c b2_committer;

-- Create rollbacks table
CREATE TABLE IF NOT EXISTS rollbacks
(
    id         SERIAL PRIMARY KEY,
    created_at timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    blockchain varchar(32) NOT NULL,
    event_id   bigint      NOT NULL
);
CREATE INDEX if NOT EXISTS event_id_index ON rollbacks (event_id);

-- Create sync_blocks table
CREATE TABLE IF NOT EXISTS sync_blocks
(
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    blockchain   VARCHAR(32) NOT NULL,
    miner        VARCHAR(42) NOT NULL,
    block_time   BIGINT      NOT NULL,
    block_number BIGINT      NOT NULL,
    block_hash   VARCHAR(66) NOT NULL,
    tx_count     BIGINT      NOT NULL,
    event_count  BIGINT      NOT NULL,
    parent_hash  VARCHAR(66) NOT NULL,
    status       VARCHAR(32) NOT NULL,
    check_count  BIGINT      NOT NULL
);
CREATE INDEX if not exists status_index ON sync_blocks (status);
CREATE INDEX if not exists tx_count_index ON sync_blocks (tx_count);
CREATE INDEX if not exists check_count_index ON sync_blocks (check_count);

-- Create sync_blocks_history table
CREATE TABLE IF NOT EXISTS sync_blocks_history
(
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    blockchain   VARCHAR(32) NOT NULL,
    miner        VARCHAR(42) NOT NULL,
    block_time   BIGINT      NOT NULL,
    block_number BIGINT      NOT NULL,
    block_hash   VARCHAR(66) NOT NULL,
    tx_count     BIGINT      NOT NULL,
    event_count  BIGINT      NOT NULL,
    parent_hash  VARCHAR(66) NOT NULL,
    status       VARCHAR(32) NOT NULL,
    check_count  BIGINT      NOT NULL
);

-- Create sync_events table
CREATE TABLE IF NOT EXISTS sync_events
(
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sync_block_id     BIGINT      NOT NULL,
    blockchain        VARCHAR(32) NOT NULL,
    block_time        BIGINT      NOT NULL,
    block_number      BIGINT      NOT NULL,
    block_hash        VARCHAR(66) NOT NULL,
    block_log_indexed BIGINT      NOT NULL,
    tx_index          BIGINT      NOT NULL,
    tx_hash           VARCHAR(66) NOT NULL,
    event_name        VARCHAR(32) NOT NULL,
    event_hash        VARCHAR(66) NOT NULL,
    contract_address  VARCHAR(42) NOT NULL,
    data              JSONB       NOT NULL,
    status            VARCHAR(32) NOT NULL,
    retry_count       BIGINT               DEFAULT 0
);

-- Create sync_events_history table
CREATE TABLE IF NOT EXISTS sync_events_history
(
    id                SERIAL PRIMARY KEY,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    sync_block_id     BIGINT      NOT NULL,
    blockchain        VARCHAR(32) NOT NULL,
    block_time        BIGINT      NOT NULL,
    block_number      BIGINT      NOT NULL,
    block_hash        VARCHAR(66) NOT NULL,
    block_log_indexed BIGINT      NOT NULL,
    tx_index          BIGINT      NOT NULL,
    tx_hash           VARCHAR(66) NOT NULL,
    event_name        VARCHAR(32) NOT NULL,
    event_hash        VARCHAR(66) NOT NULL,
    contract_address  VARCHAR(42) NOT NULL,
    data              JSONB       NOT NULL,
    status            VARCHAR(32) NOT NULL
);

-- Create sync_tasks table
CREATE TABLE IF NOT EXISTS sync_tasks
(
    id           SERIAL PRIMARY KEY,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    blockchain   VARCHAR(32) NOT NULL,
    latest_block BIGINT      NOT NULL,
    start_block  BIGINT      NOT NULL,
    end_block    BIGINT      NOT NULL,
    handle_num   BIGINT      NOT NULL,
    contracts    TEXT        NOT NULL,
    status       VARCHAR(32) NOT NULL
);

drop table if exists proposal;
-- Create proposal table
CREATE TABLE IF NOT EXISTS proposal
(
    id              bigserial PRIMARY KEY,
    created_at      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    proposal_id     bigint       NOT NULL,
    state_root_hash varchar(128) NOT NULL,
    proof_root_hash varchar(128) NOT NULL,
    start_batch_num bigint       NOT NULL,
    end_batch_num   bigint       NOT NULL,
    btc_tx_hash     varchar(128),
    winner          varchar(128),
    status          bigint       NOT NULL DEFAULT 0
);
CREATE INDEX if not exists proposal_id_index ON proposal (proposal_id);
