create table if not exists emails
(
  event_id     text not null primary key,
  message_id   text not null,
  received_at  datetime,
  subject      text,
  from_name    text,
  from_address text,
  to_name      text,
  to_address   text,
  body_html    text,
  body_txt     text,
  custom_args  json,
  categories   json
);

create index if not exists ix_received_at on emails (received_at);
create index if not exists ix_to_address on emails (to_address);
create index if not exists ix_to_subject on emails (subject);
