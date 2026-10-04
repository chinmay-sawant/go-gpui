---
name: catnotify
description: Send a local desktop-cat notification when an agent response is ready, using the HTTP server on port 6969.
---

# Cat response notification

When your response is ready, send one short notification before delivering it.
Use the actual agent name as `source`, such as `Open Code` or `Codex`.
Use `happy` for a ready response. Describe the real result in at most 160
characters. The bubble appears without moving the user's mouse.

```sh
curl --silent --show-error --fail-with-body --max-time 2 \
  http://127.0.0.1:6969/notify \
  -H 'Content-Type: application/json' \
  --data-binary '{"expression":"happy","message":"Response is ready. The work is done.","source":"Open Code"}'
```

Serialize dynamic text as JSON before passing it to curl. Keep shell commands
literal so quotes or dollar signs in a message cannot execute as shell code.
HTTP 202 means the server accepted the notification. If the server is offline,
finish the response normally and mention that the notification was unavailable.
Send once, without retries or launching the cat automatically.

For another emotion, query `GET http://127.0.0.1:6969/expressions`.
Use `worried` for a blocker, `sad` for a failed task, or `sleepy` for a sleep
reminder the user requested. Only report completion when the work is complete.

Deliver notifications through HTTP only. Leave the cursor position and the
user's active application alone.
