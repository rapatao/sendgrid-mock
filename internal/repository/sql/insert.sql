insert into emails(event_id,
                   message_id,
                   received_at,
                   subject,
                   from_name,
                   from_address,
                   to_name,
                   to_address,
                   body_html,
                   body_txt)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
