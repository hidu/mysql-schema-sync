
-- Table : user_billing_error_receipt
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_billing_error_receipt` (
  `user_id` int(10) unsigned NOT NULL,
  `product_id` varchar(255) NOT NULL,
  `receipt_id` varchar(255) NOT NULL,
  `signature` text,
  `receipt` mediumtext NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`product_id`,`receipt_id`),
  KEY `ix_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



-- Table : user_challenge_beginner
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_challenge_beginner` (
  `user_id` int(10) unsigned NOT NULL,
  `challenge_m_id` int(10) unsigned NOT NULL,
  `reward_received_count` int(10) unsigned NOT NULL,
  `is_completed` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`challenge_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_challenge_cell
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_challenge_cell` (
  `user_id` int(10) unsigned NOT NULL,
  `cell_m_id` int(10) unsigned NOT NULL,
  `challenge_start_count` int(10) unsigned NOT NULL,
  `is_received_reward` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`cell_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_challenge_subscription
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_challenge_subscription` (
  `user_id` int(10) unsigned NOT NULL,
  `challenge_set_m_id` int(10) unsigned NOT NULL,
  `end_at` int(10) unsigned NOT NULL,
  `is_completed` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`challenge_set_m_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_communication_member_detail_badge
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_communication_member_detail_badge` (
  `user_id` int(10) unsigned NOT NULL,
  `member_master_id` int(10) unsigned NOT NULL,
  `is_story_member_badge` tinyint(1) NOT NULL,
  `is_story_side_badge` tinyint(1) NOT NULL,
  `is_voice_badge` tinyint(1) NOT NULL,
  `is_theme_badge` tinyint(1) NOT NULL,
  `is_card_badge` tinyint(1) NOT NULL,
  `is_music_badge` tinyint(1) NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`member_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_coop_room_accessory
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_coop_room_accessory` (
  `user_id` int(10) unsigned NOT NULL,
  `room_id` int(10) unsigned NOT NULL,
  `user_accessory_id` bigint(20) unsigned NOT NULL,
  `accessory_master_id` int(10) unsigned NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `grade` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`room_id`,`user_accessory_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_coop_room_card
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_coop_room_card` (
  `user_id` int(10) unsigned NOT NULL,
  `room_id` int(10) unsigned NOT NULL,
  `card_master_id` int(10) unsigned NOT NULL,
  `deck_position` int(10) unsigned NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `grade` int(10) unsigned NOT NULL,
  `is_awakening_image` tinyint(1) NOT NULL,
  `is_all_training_activated` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`room_id`,`card_master_id`,`deck_position`,`level`,`grade`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_coop_room_squad
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_coop_room_squad` (
  `user_id` int(10) unsigned NOT NULL,
  `room_id` int(10) unsigned NOT NULL,
  `squad_id` int(10) unsigned NOT NULL,
  `card_master_id_1` int(10) unsigned NOT NULL,
  `card_master_id_2` int(10) unsigned NOT NULL,
  `card_master_id_3` int(10) unsigned NOT NULL,
  `user_accessory_id_1` bigint(20) unsigned DEFAULT NULL,
  `user_accessory_id_2` bigint(20) unsigned DEFAULT NULL,
  `user_accessory_id_3` bigint(20) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`room_id`,`squad_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_event_coop
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_event_coop` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `current_room_id` int(10) unsigned NOT NULL,
  `event_point` int(10) unsigned NOT NULL,
  `recent_award_id` int(10) unsigned NOT NULL,
  `event_voltage_point` int(10) unsigned NOT NULL,
  `coop_point` int(10) unsigned NOT NULL,
  `coop_point_reset_at` int(10) unsigned NOT NULL,
  `coop_point_broken` int(10) unsigned NOT NULL,
  `playable_at` int(10) unsigned NOT NULL,
  `penalty_count` int(10) unsigned NOT NULL,
  `penalty_reset_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_event_coop_daily_ranking_reward_received
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_event_coop_daily_ranking_reward_received` (
  `user_id` int(10) unsigned NOT NULL,
  `event_id` int(10) unsigned NOT NULL,
  `day_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_id`,`day_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_event_coop_global_reward_received
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_event_coop_global_reward_received` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `reward_master_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`,`reward_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_event_coop_live_play_result
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_event_coop_live_play_result` (
  `user_id` int(10) unsigned NOT NULL,
  `event_master_id` int(10) unsigned NOT NULL,
  `day_id` int(10) unsigned NOT NULL,
  `live_master_id` int(10) unsigned NOT NULL,
  `play_count` int(10) unsigned NOT NULL,
  `highest_voltage` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`event_master_id`,`day_id`,`live_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_external_movie
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_external_movie` (
  `user_id` int(10) unsigned NOT NULL,
  `external_movie_id` int(10) unsigned NOT NULL,
  `reward_received` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`external_movie_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_gacha_stepup
-- Type  : alter
-- RelationTables : 
-- SQL   : 
ALTER TABLE `user_gacha_stepup`
ADD `recover_at` int(10) unsigned NOT NULL AFTER step_count;



-- Table : user_info_trigger_subscription
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_info_trigger_subscription` (
  `user_id` int(10) unsigned NOT NULL,
  `trigger_id` bigint(20) unsigned NOT NULL,
  `info_trigger_type` tinyint(3) unsigned NOT NULL,
  `start_at` int(10) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trigger_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_live_daily
-- Type  : alter
-- RelationTables : 
-- SQL   : 
ALTER TABLE `user_live_daily`
ADD `refill_count_per_day` int(10) unsigned NOT NULL AFTER play_count_per_day;



-- Table : user_live_last_play_deck
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_live_last_play_deck` (
  `user_id` int(10) unsigned NOT NULL,
  `live_difficulty_id` int(10) unsigned NOT NULL,
  `score` int(10) unsigned NOT NULL,
  `is_cleared` tinyint(1) NOT NULL,
  `recorded_at` int(10) unsigned NOT NULL,
  `card_master_id_1` int(10) unsigned NOT NULL,
  `card_master_id_2` int(10) unsigned NOT NULL,
  `card_master_id_3` int(10) unsigned NOT NULL,
  `card_master_id_4` int(10) unsigned NOT NULL,
  `card_master_id_5` int(10) unsigned NOT NULL,
  `card_master_id_6` int(10) unsigned NOT NULL,
  `card_master_id_7` int(10) unsigned NOT NULL,
  `card_master_id_8` int(10) unsigned NOT NULL,
  `card_master_id_9` int(10) unsigned NOT NULL,
  `suit_master_id_1` int(10) unsigned NOT NULL,
  `suit_master_id_2` int(10) unsigned NOT NULL,
  `suit_master_id_3` int(10) unsigned NOT NULL,
  `suit_master_id_4` int(10) unsigned NOT NULL,
  `suit_master_id_5` int(10) unsigned NOT NULL,
  `suit_master_id_6` int(10) unsigned NOT NULL,
  `suit_master_id_7` int(10) unsigned NOT NULL,
  `suit_master_id_8` int(10) unsigned NOT NULL,
  `suit_master_id_9` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_difficulty_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_live_last_play_squad
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_live_last_play_squad` (
  `user_id` int(10) unsigned NOT NULL,
  `live_difficulty_id` int(10) unsigned NOT NULL,
  `squad1_card_master_id_1` int(10) unsigned NOT NULL,
  `squad1_card_master_id_2` int(10) unsigned NOT NULL,
  `squad1_card_master_id_3` int(10) unsigned NOT NULL,
  `squad2_card_master_id_1` int(10) unsigned NOT NULL,
  `squad2_card_master_id_2` int(10) unsigned NOT NULL,
  `squad2_card_master_id_3` int(10) unsigned NOT NULL,
  `squad3_card_master_id_1` int(10) unsigned NOT NULL,
  `squad3_card_master_id_2` int(10) unsigned NOT NULL,
  `squad3_card_master_id_3` int(10) unsigned NOT NULL,
  `squad1_user_accessory_id_1` bigint(20) unsigned DEFAULT NULL,
  `squad1_user_accessory_id_2` bigint(20) unsigned DEFAULT NULL,
  `squad1_user_accessory_id_3` bigint(20) unsigned DEFAULT NULL,
  `squad2_user_accessory_id_1` bigint(20) unsigned DEFAULT NULL,
  `squad2_user_accessory_id_2` bigint(20) unsigned DEFAULT NULL,
  `squad2_user_accessory_id_3` bigint(20) unsigned DEFAULT NULL,
  `squad3_user_accessory_id_1` bigint(20) unsigned DEFAULT NULL,
  `squad3_user_accessory_id_2` bigint(20) unsigned DEFAULT NULL,
  `squad3_user_accessory_id_3` bigint(20) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_difficulty_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_live_play_summary
-- Type  : alter
-- RelationTables : 
-- SQL   : 
ALTER TABLE `user_live_play_summary`
ADD `highest_sp_skill_at_once` int(10) unsigned NOT NULL AFTER heal_50,
drop `heal_100`;



-- Table : user_live_tower
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_live_tower` (
  `user_id` int(10) unsigned NOT NULL,
  `live_id` bigint(20) unsigned NOT NULL,
  `tower_id` int(10) unsigned NOT NULL,
  `tower_floor_no` int(10) unsigned NOT NULL,
  `deck_id` int(10) unsigned NOT NULL,
  `note_drop` text NOT NULL,
  `is_autoplay` tinyint(1) NOT NULL,
  `autoplay_judgestat` text NOT NULL,
  `magnification` int(10) unsigned NOT NULL,
  `live_difficulty_master_id` int(10) unsigned NOT NULL,
  `live_started_at` int(10) unsigned NOT NULL,
  `finish_status` tinyint(3) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`live_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_member_love_panel
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_member_love_panel` (
  `user_id` int(10) unsigned NOT NULL,
  `member_master_id` int(10) unsigned NOT NULL,
  `member_love_panel_cell_master_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`member_master_id`,`member_love_panel_cell_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_read_daily_theater
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_read_daily_theater` (
  `user_id` int(10) unsigned NOT NULL,
  `daily_theater_id` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`daily_theater_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_recovery_tower_card_used_count_item
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_recovery_tower_card_used_count_item` (
  `user_id` int(10) unsigned NOT NULL,
  `recovery_tower_card_used_count_item_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`recovery_tower_card_used_count_item_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_repro
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_repro` (
  `user_id` int(10) unsigned NOT NULL,
  `group_no` int(10) unsigned NOT NULL,
  `is_enabled_billing_pack` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`),
  KEY `ix_user_repro_group_no` (`group_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_shop_item_exchange_monthly_limit
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_shop_item_exchange_monthly_limit` (
  `user_id` int(10) unsigned NOT NULL,
  `monthly_id` int(10) unsigned NOT NULL,
  `item_exchange_content_master_id` int(10) unsigned NOT NULL,
  `exchange_count` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`monthly_id`,`item_exchange_content_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_status
-- Type  : alter
-- RelationTables : 
-- SQL   : 
ALTER TABLE `user_status`
ADD `live_point_subscription_recovery_daily_count` int(10) unsigned NOT NULL AFTER live_point_broken,
ADD `live_point_subscription_recovery_daily_reset_at` int(10) unsigned NOT NULL AFTER live_point_subscription_recovery_daily_count,
ADD `subscription_coin` int(10) unsigned NOT NULL AFTER google_sns_coin,
ADD `last_user_mission_reset_at` int(10) unsigned NOT NULL AFTER gdpr_version,
ADD `create_user_at` int(10) unsigned NOT NULL AFTER last_user_mission_reset_at,
ADD KEY `ix_create_user_at` (`create_user_at`);



-- Table : user_story_event_history
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_story_event_history` (
  `user_id` int(10) unsigned NOT NULL,
  `story_event_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_story_event_unlock_item
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_story_event_unlock_item` (
  `user_id` int(10) unsigned NOT NULL,
  `story_event_unlock_item_master_id` int(10) unsigned NOT NULL,
  `amount` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_event_unlock_item_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_story_linkage
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_story_linkage` (
  `user_id` int(10) unsigned NOT NULL,
  `story_linkage_cell_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_linkage_cell_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_story_main_part_digest_movie
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_story_main_part_digest_movie` (
  `user_id` int(10) unsigned NOT NULL,
  `story_main_part_master_id` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`story_main_part_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_subscription_cancel_history
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_subscription_cancel_history` (
  `user_id` int(10) unsigned NOT NULL,
  `subscription_history_nsec` bigint(20) unsigned NOT NULL,
  `cancel_date` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`subscription_history_nsec`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_subscription_end_period_history
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_subscription_end_period_history` (
  `user_id` int(10) unsigned NOT NULL,
  `subscription_history_nsec` bigint(20) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`subscription_history_nsec`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_subscription_received
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_subscription_received` (
  `user_id` int(10) unsigned NOT NULL,
  `subscription_master_id` int(10) unsigned NOT NULL,
  `created_nsec` bigint(20) unsigned NOT NULL,
  `continue_count` int(10) unsigned NOT NULL,
  `start_at` int(10) unsigned NOT NULL,
  `end_at` int(10) unsigned NOT NULL,
  `subscription_coin_last_received_at` int(10) unsigned NOT NULL,
  `continue_reward_status` tinyint(3) unsigned NOT NULL,
  `is_read_subscription_pass_grade_up` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`subscription_master_id`,`created_nsec`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_subscription_status
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_subscription_status` (
  `user_id` int(10) unsigned NOT NULL,
  `subscription_master_id` int(10) unsigned NOT NULL,
  `start_date` int(10) unsigned NOT NULL,
  `expire_date` int(10) unsigned NOT NULL,
  `platform_expire_date` int(10) unsigned NOT NULL,
  `renewal_count` int(10) unsigned NOT NULL,
  `continue_count` int(10) unsigned NOT NULL,
  `subscription_pass_id` bigint(20) unsigned NOT NULL,
  `attach_id` varchar(255) NOT NULL,
  `is_auto_renew` tinyint(1) NOT NULL,
  `is_done_trial` tinyint(1) NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`subscription_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_tower
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_tower` (
  `user_id` int(10) unsigned NOT NULL,
  `tower_id` int(10) unsigned NOT NULL,
  `cleared_floor` int(10) unsigned NOT NULL,
  `read_floor` int(10) unsigned NOT NULL,
  `voltage` int(10) unsigned NOT NULL,
  `recovery_point_full_at` int(10) unsigned NOT NULL,
  `recovery_point_last_consumed_at` int(10) unsigned NOT NULL,
  `is_new` tinyint(1) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`tower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_tower_card_used_count
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_tower_card_used_count` (
  `user_id` int(10) unsigned NOT NULL,
  `tower_id` int(10) unsigned NOT NULL,
  `card_master_id` int(10) unsigned NOT NULL,
  `used_count` int(10) unsigned NOT NULL,
  `recovered_count` int(10) unsigned NOT NULL,
  `last_used_at` int(10) unsigned NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`tower_id`,`card_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_trade
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_trade` (
  `user_id` int(10) unsigned NOT NULL,
  `trade_product_master_id` int(11) NOT NULL,
  `trade_count` int(10) unsigned NOT NULL,
  `monthly_id` int(10) unsigned DEFAULT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`trade_product_master_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



-- Table : user_wsnet
-- Type  : create
-- RelationTables : 
-- SQL   : 
CREATE TABLE `user_wsnet` (
  `user_id` int(10) unsigned NOT NULL,
  `auth_string` varchar(255) NOT NULL,
  `created_at` int(10) unsigned NOT NULL,
  `updated_at` int(10) unsigned NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


