alter table articles
  add body mediumtext NOT NULL,
  add created datetime,
  add updated datetime;

update articles set created = CURRENT_TIMESTAMP where created is null;
update articles set updated = CURRENT_TIMESTAMP where updated is null;