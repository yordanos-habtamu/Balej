#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${BASE_URL:-http://localhost:8000}"
QUERY_URL="$BASE_URL/query"
CV_PATH="${CV_PATH:-}"

log() { printf "\n==> %s\n" "$*"; }
post() {
  local payload="$1"
  curl -s -X POST "$QUERY_URL" -H 'Content-Type: application/json' -d "$payload"
}

log "Health: users"
post '{"query":"query { users { id email role } }"}'

log "Register"
post '{"query":"mutation($input: NewUser!){ Register(input: $input){ id firstName lastName email role } }","variables":{"input":{"id":"temp-id-1","firstName":"Test","lastName":"User","email":"testuser@example.com","password":"Passw0rd!","phoneNumber":"1234567890","address":"Addis Ababa","role":"client","plan":"basic"}}}'

log "Login"
post '{"query":"mutation($input: LoginInput!){ LoginUser(input: $input){ token user { id email role } } }","variables":{"input":{"email":"testuser@example.com","password":"Passw0rd!"}}}'

log "CreateUser"
post '{"query":"mutation($input: NewUser!){ createUser(input: $input){ id email role } }","variables":{"input":{"id":"temp-id-2","firstName":"Another","lastName":"User","email":"another@example.com","password":"Passw0rd!","phoneNumber":"111222333","address":"Addis Ababa","role":"client","plan":"basic"}}}'

log "AddJobOffer"
post '{"query":"mutation($input: NewJobOffer!){ AddJobOffer(input: $input){ job_id title location category applicationCount } }","variables":{"input":{"job_type":"full_time","title":"Senior Designer","description":"Design leadership role","location":"Remote","offeredByuser":1,"category":"design","applicationCount":0}}}'

log "Jobs"
post '{"query":"query { jobs { job_id title location category applicationCount } }"}'

log "JobsWithFilter"
post '{"query":"query($filter: JobFilter){ jobsWithFilter(filter: $filter){ job_id title category } }","variables":{"filter":{"category":"design"}}}'

log "Applicants"
post '{"query":"query { applicants { userId job_id status resume } }"}'

log "ApplicantsWithFilter"
post '{"query":"query($filter: JobApplicantFilter){ applicantsWithfilter(filter: $filter){ userId job_id status } }","variables":{"filter":{"status":"APPROVED"}}}'

log "ApproveApplication"
post '{"query":"mutation($input: Approval!){ ApproveApplication(input: $input){ response } }","variables":{"input":{"userId":1,"job_id":1}}}'

log "RejectApplication"
post '{"query":"mutation($input: Approval!){ RejectApplication(input: $input){ response } }","variables":{"input":{"userId":1,"job_id":1}}}'

log "SendMessage"
post '{"query":"mutation { SendMessage(chatID:\"chat-1\", content:\"Hello\", senderId:1, recieverId:2){ id message senderId chatID sentAt isRead isFile fileUrl fileName } }"}'

if [[ -n "$CV_PATH" && -f "$CV_PATH" ]]; then
  log "ApplyToJob (file upload)"
  curl -s -X POST "$QUERY_URL" \
    -H 'Apollo-Require-Preflight: true' \
    -F 'operations={"query":"mutation($input: JobApplication!){ ApplyToJob(input: $input){ response user { id email } } }","variables":{"input":{"jobApplicationId":1,"job_id":1,"userId":1,"proposal":"I can do this.","cv":null}}}' \
    -F 'map={"0":["variables.input.cv"]}' \
    -F "0=@$CV_PATH"
else
  log "ApplyToJob skipped (set CV_PATH=/path/to/cv.pdf)"
fi

log "Done"
