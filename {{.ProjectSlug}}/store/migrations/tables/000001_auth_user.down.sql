DROP TRIGGER IF EXISTS "trg-auth_user_token-disable_old_tokens" ON auth_user;
DROP FUNCTION IF EXISTS fn_auth_user_token_disable_old_tokens;

drop table if exists auth_user_token;
drop table if exists auth_user;
