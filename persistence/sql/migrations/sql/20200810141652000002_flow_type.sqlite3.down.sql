INSERT INTO "_selfservice_verification_requests_tmp" (id, request_url, issued_at, expires_at, csrf_token, created_at, updated_at, nid, messages, form, via, success) SELECT id, request_url, issued_at, expires_at, csrf_token, created_at, updated_at, nid, messages, form, via, success FROM "selfservice_verification_requests";