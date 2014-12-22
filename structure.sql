create table if not exists player_character (
  id bigserial primary key
  , name text
  , rolled_str integer
  , rolled_int integer
  , rolled_wis integer
  , rolled_con integer
  , rolled_con integer
  , rolled_cha integer
  , race_id integer
  , class_id integer
  , "level" integer
  , max_hit_points integer
  , base_armor_class integer
);

create table if not exists race (
  id bigserial primary key
  , name text
  , subrace text
  , str_mod integer
  , int_mod integer
  , wis_mod integer
  , dex_mod integer
  , con_mod integer
  , cha_mod integer
  , hit_point_mod integer
);

alter table if exists player_character add column created_at timestamp;

alter table if exists race add column created_at timestamp;
