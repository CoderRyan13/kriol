-- Filename: 000002_add_primary_key_to_entries_table.up.sql

ALTER table entries
ADD id bigserial PRIMARY KEY;