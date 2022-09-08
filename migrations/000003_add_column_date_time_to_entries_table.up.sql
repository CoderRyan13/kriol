-- Filename: 000003_add_column_date_time_to_entries_table.up.sql

ALTER TABLE entries
ADD COLUMN date_time timestamp(0) with time zone NOT NULL DEFAULT NOW();