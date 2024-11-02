select e.event_id,
       e.message_id,
       e.received_at,
       e.subject,
       e.from_name,
       e.from_address,
       e.to_name,
       e.to_address,
       e.body_html,
       e.body_txt,
       e.custom_args,
       e.categories
from emails e
where e.event_id = ?
limit 1
