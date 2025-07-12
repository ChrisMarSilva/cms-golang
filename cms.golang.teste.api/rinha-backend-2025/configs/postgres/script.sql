
CREATE TABLE "TbUser" (
  "id" SERIAL PRIMARY KEY, 
  "name" VARCHAR(50) NOT NULL,
  "email" VARCHAR(255) UNIQUE NOT null
);

CREATE INDEX CONCURRENTLY IF NOT EXISTS "idx_user_email" ON "TbUser"("email");

/*
CREATE TABLE IF NOT EXISTS summary_snapshot (
    ts TIMESTAMPTZ PRIMARY KEY,
    default_count BIGINT NOT NULL,
    fallback_count BIGINT NOT NULL,
    default_cents BIGINT NOT NULL,
    fallback_cents BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_snapshot_ts ON summary_snapshot(ts DESC);


*/