with params as (select ? as to_address,
                       ? as subject),
     parsed as (select case
                         when p.to_address is null then '%'
                         else '%' || p.to_address || '%'
                         end as to_address,
                       case
                         when p.subject is null then '%'
                         else '%' || p.subject || '%'
                         end as subject
                from params p),
     results as (select e.*
                 from emails e
                        join parsed p on e.to_address like p.to_address and e.subject like p.subject
                 order by e.received_at desc)
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
from results e
limit (? * ?), ?
