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
                from params p)
select count(*)
from emails e
       join parsed p on e.to_address like p.to_address and e.subject like p.subject
