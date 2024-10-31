import http from 'k6/http';
import {check} from 'k6';

export const options = {
  vus: 50,
  duration: "10s",
};

export default function () {
  let params = {
    headers: {
      "Authorization": "Bearer mock.luPzMYLzMTWJqMLCO37ZJRmPllQ7ct78",
      "Content-Type": "application/json",
    },
  }

  let body = {
    "personalizations": [
      {
        "to": [
          {
            "email": "alex@example.com",
            "name": "Alex"
          },
          {
            "email": "bola@example.com",
            "name": "Bola"
          }
        ],
        "cc": [
          {
            "email": "charlie@example.com",
            "name": "Charlie"
          }
        ],
        "bcc": [
          {
            "email": "dana@example.com",
            "name": "Dana"
          }
        ]
      },
      {
        "from": {
          "email": "sales@example.com",
          "name": "Example Sales Team"
        },
        "to": [
          {
            "email": "ira@example.com",
            "name": "Ira"
          }
        ],
        "bcc": [
          {
            "email": "lee@example.com",
            "name": "Lee"
          }
        ]
      }
    ],
    "from": {
      "email": "orders@example.com",
      "name": "Example Order Confirmation"
    },
    "reply_to": {
      "email": "customer_service@example.com",
      "name": "Example Customer Service Team"
    },
    "subject": "Your Example Order Confirmation",
    "content": [
      {
        "type": "text/html",
        "value": "<p>Hello from Twilio SendGrid!</p><p>Sending with the email service trusted by developers and marketers for <strong>time-savings</strong>, <strong>scalability</strong>, and <strong>delivery expertise</strong>.</p><p>%open-track%</p>"
      }
    ],
    "attachments": [
      {
        "content": "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KCiAgICA8aGVhZD4KICAgICAgICA8bWV0YSBjaGFyc2V0PSJVVEYtOCI+CiAgICAgICAgPG1ldGEgaHR0cC1lcXVpdj0iWC1VQS1Db21wYXRpYmxlIiBjb250ZW50PSJJRT1lZGdlIj4KICAgICAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoLCBpbml0aWFsLXNjYWxlPTEuMCI+CiAgICAgICAgPHRpdGxlPkRvY3VtZW50PC90aXRsZT4KICAgIDwvaGVhZD4KCiAgICA8Ym9keT4KCiAgICA8L2JvZHk+Cgo8L2h0bWw+Cg==",
        "filename": "index.html",
        "type": "text/html",
        "disposition": "attachment"
      }
    ],
    "categories": [
      "cake",
      "pie",
      "baking"
    ],
    "send_at": 1617260400,
    "batch_id": "AsdFgHjklQweRTYuIopzXcVBNm0aSDfGHjklmZcVbNMqWert1znmOP2asDFjkl",
    "asm": {
      "group_id": 12345,
      "groups_to_display": [
        12345
      ]
    },
    "ip_pool_name": "transactional email",
    "mail_settings": {
      "bypass_list_management": {
        "enable": false
      },
      "footer": {
        "enable": false
      },
      "sandbox_mode": {
        "enable": false
      }
    },
    "tracking_settings": {
      "click_tracking": {
        "enable": true,
        "enable_text": false
      },
      "open_tracking": {
        "enable": true,
        "substitution_tag": "%open-track%"
      },
      "subscription_tracking": {
        "enable": false
      }
    }
  }

  const response = http.post('http://localhost:3000/v3/mail/send', JSON.stringify(body), params)

  check(response, {
    'is status 202': (r) => r.status === 202,
    'has X-Message-Id header': (r) => r.headers['X-Message-Id'] !== undefined,
  })

}
