# SendGrid Mock

![GitHub Release](https://img.shields.io/github/v/release/rapatao/sendgrid-mock)

Provides a service that simulates the sending mail endpoint provided by SendGrid API that may only be used for
development purposes.

It contains a simple UI where you can:

* list received mails, filtering by recipient and subject, with pagination;
* view the HTML, text or raw HTML content of a mail;
* list and download attachments;
* delete a single mail or all of them;
* switch between light, dark or automatic theme.

![UI Screenshot](screenshot.png)

## API

* `POST /v3/mail/send`: SendGrid like sending mail endpoint
  * Requires the `Authorization: Bearer <API_KEY>` header
  * The request body is validated against the SendGrid mail send schema
  * Returns `202 Accepted` with the `X-Message-Id` header
* `GET /messages`: search persisted messages
  * Query parameters:
    * `to`: filter by to address (accepts `%` wildcard)
    * `subject`: filter by subject (accepts `%` wildcard)
    * `page`: paginate results starting from `0`
    * `rows`: number of rows returned on each page (default `10`)
* `GET /messages/:event_id`: return message content
  * Path parameter:
    * `event_id`: event identifier
  * Query parameter:
    * `format`: must be one of:
      * `html`: rendered HTML, with links rewritten to track clicks
      * `text`: plain text content
      * `raw`: original HTML source, returned as plain text
* `GET /messages/:event_id/attachments/:filename`: download an attachment
  * Path parameters:
    * `event_id`: event identifier
    * `filename`: attachment file name
* `GET /messages/:event_id/links/:encoded_link`: track a click and redirect to the decoded link (used by the rendered
  HTML content)
* `DELETE /messages/:event_id`: delete message
  * Path parameter:
    * `event_id`: event identifier
* `DELETE /messages`: delete all messages
* `GET /health`: health check

Errors are returned as JSON describing the failed field, the reason and a hint on how to fix it.

## Events

There is a basic support for events, currently sending (when enabled) one of these events:

* **delivered**: when the message was persisted in the database;
* **dropped**: when some failure occurs when persisting in the database;
* **open**: when mail content is retrieved through the `GET /messages/:event_id` endpoint;
* **click**: when a link is clicked in the rendered content;

Each event includes the message `custom_args` and categories.

## Configuration

There are some customization possibles, all of them are defined by using the system environment.

* **API_KEY**: Used to authenticate the sending mail endpoint, defaults to `mock.luPzMYLzMTWJqMLCO37ZJRmPllQ7ct78`;
* **EVENT_DELIVERY_URL**: when defined, it used for event sending;
* **MAIL_HISTORY_DURATION**: duration, using Golang pattern, which is used to delete messages older than that duration,
  defaults to `24h`;
* **STORAGE_FILE**: when defined, instead of creating a randomized temporary file, the given file name will be used;
* **BLOCK_DELETE_ALL**: when value is `true`, deleting all messages will not be allowed;
* **EVENT_TIMESTAMP_DELAY**: duration, using Golang pattern, which is used to add extra time to the events;
* **WEB_STATIC_FILES**: directory with the UI static files, defaults to `./web/`;

## Running

The service listens on port `3000`.

### Container (Docker)

```shell
docker run -p 3000:3000 --rm ghcr.io/rapatao/sendgrid-mock:${version:-latest}
```

### From source

Requires Go and a C compiler (CGO is used by SQLite).

```shell
make run
```

### Sending a message

```shell
curl -X POST "http://localhost:3000/v3/mail/send" \
  -H "Authorization: Bearer mock.luPzMYLzMTWJqMLCO37ZJRmPllQ7ct78" \
  -H "Content-Type: application/json" \
  --data '{"personalizations": [{"to": [{"email": "alex@example.com"}]}], "from": {"email": "orders@example.com"}, "subject": "Hello", "content": [{"type": "text/html", "value": "<p>Hello!</p>"}]}'
```

More examples are available in the [scripts](scripts) directory.

## Notes

This service/repository should not be used as a studying reference since it isn't using good development practices
and has limited test coverage. It was built with the intention to provide a service that should be used for development
purposes only.
