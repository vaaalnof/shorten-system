-- ==================================================================================================================
--                                               AUTH SERVICE
-- ==================================================================================================================

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =========================================================
-- USERS
-- =========================================================

CREATE TABLE users (
    id UUID PRIMARY KEY
        DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    avatar_url TEXT,
    is_active BOOLEAN NOT NULL
        DEFAULT TRUE,
    email_verified BOOLEAN NOT NULL
        DEFAULT FALSE,
    created_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT,
    updated_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- =========================================================
-- USER INDEXES
-- =========================================================

CREATE UNIQUE INDEX idx_users_email
    ON users(email);

-- =========================================================
-- AUTH PROVIDERS
-- SUPPORT:
-- - LOCAL LOGIN
-- - GOOGLE LOGIN
-- - GITHUB LOGIN
-- =========================================================

CREATE TABLE user_auth_providers (
    id UUID PRIMARY KEY
        DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL
        REFERENCES users(id)
            ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL,
    provider_user_id VARCHAR(255),
    password_hash TEXT,
    created_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT,
    CONSTRAINT chk_provider
        CHECK (
            provider IN ('local', 'google', 'github')
            ),
    CONSTRAINT uq_provider_user
        UNIQUE (provider,provider_user_id)
);

-- =========================================================
-- AUTH PROVIDER INDEXES
-- =========================================================

CREATE INDEX idx_auth_user_id
    ON user_auth_providers(user_id);

CREATE INDEX idx_auth_provider
    ON user_auth_providers(provider);

CREATE INDEX idx_auth_provider_user_id
    ON user_auth_providers(provider,user_id);

-- =========================================================
-- USER SESSIONS
-- JWT REFRESH TOKENS
-- =========================================================

CREATE TABLE user_sessions (
    id UUID PRIMARY KEY
        DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL
        REFERENCES users(id)
            ON DELETE CASCADE,
    refresh_token TEXT NOT NULL UNIQUE,
    ip_address INET,
    user_agent TEXT,
    expired_at BIGINT NOT NULL,
    revoked_at BIGINT,
    created_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- =========================================================

-- SESSION INDEXES

-- =========================================================

CREATE INDEX idx_sessions_user_id
    ON user_sessions(user_id);

CREATE INDEX idx_sessions_expired_at
    ON user_sessions(expired_at);

CREATE INDEX idx_sessions_revoked_at
    ON user_sessions(revoked_at);

CREATE INDEX idx_sessions_user_revoked
    ON user_sessions(user_id, revoked_at);

-- =========================================================

-- OAUTH STATES

-- OAUTH SECURITY

-- =========================================================

CREATE TABLE oauth_states (
    id UUID PRIMARY KEY
        DEFAULT gen_random_uuid(),
    state VARCHAR(255) NOT NULL UNIQUE,
    provider VARCHAR(50) NOT NULL,
    expired_at BIGINT NOT NULL,
    created_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- =========================================================
-- OAUTH STATE INDEXES
-- =========================================================

CREATE INDEX idx_oauth_states_expired_at
    ON oauth_states(expired_at);



-- ==================================================================================================================
--                                               URL SHORTENER SERVICE
-- ==================================================================================================================


-- =========================================================
-- RESERVED ALIASES
-- =========================================================

CREATE TABLE reserved_aliases (
    id BIGSERIAL PRIMARY KEY,
    keyword VARCHAR(100) NOT NULL UNIQUE,
    created_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- =========================================================
-- RESERVED ALIAS INDEXES
-- =========================================================

CREATE UNIQUE INDEX idx_reserved_aliases_keyword
    ON reserved_aliases(keyword);

-- =========================================================
-- RESERVED ALIAS SEED
-- =========================================================

INSERT INTO reserved_aliases(keyword)
VALUES
    -- auth
    ('login'),
    ('logout'),
    ('register'),
    ('signup'),
    ('signin'),
    ('auth'),
    ('oauth'),
    ('callback'),
    ('verify'),
    ('verification'),
    ('forgot-password'),
    ('reset-password'),
    ('refresh'),
    ('validate'),
    ('me'),
    ('profile'),

    -- api
    ('api'),
    ('v1'),
    ('v2'),
    ('v3'),

    -- dashboard
    ('dashboard'),
    ('account'),
    ('settings'),
    ('billing'),
    ('subscription'),
    ('plan'),
    ('workspace'),
    ('organization'),
    ('team'),

    -- admin
    ('admin'),
    ('administrator'),
    ('superadmin'),
    ('staff'),
    ('moderator'),

    -- system
    ('health'),
    ('metrics'),
    ('status'),
    ('ping'),
    ('monitor'),
    ('ready'),
    ('live'),

    -- documentation
    ('docs'),
    ('documentation'),
    ('swagger'),
    ('openapi'),
    ('redoc'),

    -- static files
    ('favicon.ico'),
    ('robots.txt'),
    ('sitemap.xml'),
    ('manifest.json'),

    -- assets
    ('assets'),
    ('static'),
    ('public'),
    ('uploads'),
    ('images'),
    ('img'),
    ('css'),
    ('js'),
    ('fonts'),

    -- common pages
    ('home'),
    ('about'),
    ('contact'),
    ('support'),
    ('help'),
    ('pricing'),
    ('terms'),
    ('privacy'),
    ('security'),

    -- user routes
    ('user'),
    ('users'),

    -- url routes
    ('url'),
    ('urls'),
    ('link'),
    ('links'),
    ('short'),
    ('shorten'),
    ('redirect'),

    -- analytics future
    ('analytics'),
    ('reports'),
    ('stats'),

    -- integrations
    ('webhook'),
    ('webhooks'),
    ('integration'),
    ('integrations'),

    -- jobs / workers
    ('worker'),
    ('workers'),
    ('queue'),
    ('queues'),

    -- storage
    ('storage'),
    ('files'),

    -- common exploits
    ('root'),
    ('system'),
    ('internal'),
    ('private'),

    -- legal
    ('tos'),
    ('policy'),
    ('legal'),

    -- auth future
    ('session'),
    ('sessions'),
    ('token'),
    ('tokens'),
    ('access-token'),
    ('refresh-token'),
    ('password'),
    ('change-password'),

    -- oauth providers
    ('google'),
    ('github'),
    ('facebook'),
    ('apple'),
    ('microsoft'),
    ('gitlab'),

    -- admin future
    ('audit'),
    ('logs'),

    -- frontend
    ('app'),
    ('console'),

    -- communication
    ('mail'),
    ('email'),
    ('notification'),
    ('notifications'),

    -- devops
    ('grafana'),
    ('prometheus'),
    ('kibana'),
    ('elk'),

    -- common short aliases
    ('go'),
    ('new'),
    ('create'),
    ('edit'),
    ('update'),
    ('delete'),
    ('remove'),
    ('test'),
    ('dev'),
    ('prod'),
    ('staging'),

    -- search
    ('search'),
    ('explore'),
    ('discover'),

    -- future features
    ('share'),
    ('shares'),
    ('archive'),
    ('archives'),
    ('trash'),
    ('restore'),
    ('export'),
    ('import'),

    -- payments
    ('payment'),
    ('payments'),
    ('invoice'),
    ('invoices'),

    -- security
    ('2fa'),
    ('mfa'),
    ('captcha'),

    -- common redirects
    ('www');

-- =========================================================
-- URLS
-- =========================================================

CREATE TABLE urls (
    id UUID PRIMARY KEY
        DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    short_code VARCHAR(32) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    is_active BOOLEAN NOT NULL
        DEFAULT TRUE,
    click_count BIGINT NOT NULL
        DEFAULT 0,
    password_hash TEXT,
    expired_at BIGINT,
    deleted_at BIGINT,
    created_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT,
    updated_at BIGINT NOT NULL
        DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT
);

-- =========================================================
-- URL INDEXES
-- =========================================================

CREATE UNIQUE INDEX idx_urls_short_code
    ON urls(short_code);

CREATE INDEX idx_urls_user_id
    ON urls(user_id);

CREATE INDEX idx_urls_created_at
    ON urls(created_at);

CREATE INDEX idx_urls_updated_at
    ON urls(updated_at);

CREATE INDEX idx_urls_expired_at
    ON urls(expired_at);

CREATE INDEX idx_urls_deleted_at
    ON urls(deleted_at);

CREATE INDEX idx_urls_user_deleted
    ON urls(user_id, deleted_at);

CREATE INDEX idx_urls_user_created
    ON urls(user_id, created_at);








-- ==================================================================================================================



-- =========================================================
-- COMMENTED FOR NEXT PHASE
-- URL SHORTENER DOMAIN
-- =========================================================


/*

-- =========================================================
-- RESERVED ALIASES
-- =========================================================

CREATE TABLE reserved_aliases (
    id BIGSERIAL PRIMARY KEY,
    keyword VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


INSERT INTO reserved_aliases(keyword)
VALUES
    ('api'),
    ('admin'),
    ('login'),
    ('logout'),
    ('register'),
    ('dashboard'),
    ('settings'),
    ('metrics'),
    ('health'),
    ('docs'),
    ('swagger'),
    ('graphql'),
    ('favicon.ico'),
    ('robots.txt');


-- =========================================================
-- URLS
-- =========================================================

CREATE TABLE urls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID
        REFERENCES users(id)
        ON DELETE CASCADE,

    short_code VARCHAR(32) NOT NULL UNIQUE,

    original_url TEXT NOT NULL,

    title TEXT,

    is_custom BOOLEAN NOT NULL DEFAULT FALSE,

    click_count BIGINT NOT NULL DEFAULT 0,

    password_hash TEXT,

    expired_at TIMESTAMP,

    deleted_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- =========================================================
-- URL INDEXES
-- =========================================================

CREATE UNIQUE INDEX idx_urls_short_code
ON urls(short_code);

CREATE INDEX idx_urls_user_id
ON urls(user_id);

CREATE INDEX idx_urls_created_at
ON urls(created_at);

CREATE INDEX idx_urls_expired_at
ON urls(expired_at);

CREATE INDEX idx_urls_deleted_at
ON urls(deleted_at);


-- =========================================================
-- ANALYTICS EVENTS
-- RAW CLICK EVENTS
-- =========================================================

CREATE TABLE analytics_events (
    id BIGSERIAL PRIMARY KEY,

    url_id UUID
        REFERENCES urls(id)
        ON DELETE CASCADE,

    short_code VARCHAR(32) NOT NULL,

    ip_address INET,

    country VARCHAR(100),

    city VARCHAR(100),

    device VARCHAR(50),

    os VARCHAR(50),

    browser VARCHAR(50),

    referer TEXT,

    source VARCHAR(100),

    user_agent TEXT,

    clicked_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- =========================================================
-- ANALYTICS INDEXES
-- =========================================================

CREATE INDEX idx_analytics_url_id
ON analytics_events(url_id);

CREATE INDEX idx_analytics_short_code
ON analytics_events(short_code);

CREATE INDEX idx_analytics_clicked_at
ON analytics_events(clicked_at);

CREATE INDEX idx_analytics_country
ON analytics_events(country);

CREATE INDEX idx_analytics_source
ON analytics_events(source);


-- =========================================================
-- DAILY STATS
-- AGGREGATED ANALYTICS
-- =========================================================

CREATE TABLE url_daily_stats (
    id BIGSERIAL PRIMARY KEY,

    url_id UUID
        REFERENCES urls(id)
        ON DELETE CASCADE,

    stat_date DATE NOT NULL,

    total_clicks BIGINT NOT NULL DEFAULT 0,

    mobile_clicks BIGINT NOT NULL DEFAULT 0,

    desktop_clicks BIGINT NOT NULL DEFAULT 0,

    tablet_clicks BIGINT NOT NULL DEFAULT 0,

    indonesia_clicks BIGINT NOT NULL DEFAULT 0,

    us_clicks BIGINT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE(url_id, stat_date)
);


-- =========================================================
-- DAILY STATS INDEXES
-- =========================================================

CREATE INDEX idx_url_daily_stats_url_id
ON url_daily_stats(url_id);

CREATE INDEX idx_url_daily_stats_date
ON url_daily_stats(stat_date);


-- =========================================================
-- API KEYS
-- OPTIONAL PUBLIC API ACCESS
-- =========================================================

CREATE TABLE api_keys (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID
        REFERENCES users(id)
        ON DELETE CASCADE,

    api_key TEXT NOT NULL UNIQUE,

    name VARCHAR(255),

    last_used_at TIMESTAMP,

    expired_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- =========================================================
-- API KEY INDEXES
-- =========================================================

CREATE INDEX idx_api_keys_user_id
ON api_keys(user_id);

CREATE INDEX idx_api_keys_expired_at
ON api_keys(expired_at);

*/
