CREATE TABLE "subscriptions"
(
    "id"                INTEGER PRIMARY KEY AUTOINCREMENT,
    "created_at"        DATETIME,
    "updated_at"        DATETIME,
    "deleted_at"        DATETIME,
    "link"              TEXT,
    "is_aggregate"      BOOLEAN,
    "source"            TEXT,
    "black_list_filter" TEXT,
    "white_list_filter" TEXT,
    "is_enabled"        BOOLEAN
);

CREATE TABLE "subscription_files"
(
    "id"                 INTEGER PRIMARY KEY AUTOINCREMENT,
    "created_at"         DATETIME,
    "updated_at"         DATETIME,
    "deleted_at"         DATETIME,
    "subscription_id"    INTEGER,
    "name"               TEXT,
    "link"               TEXT,
    "is_download"        BOOLEAN,
    "downloader_task_id" TEXT,
    FOREIGN KEY ("subscription_id") REFERENCES "subscriptions" ("id")
);

CREATE TABLE "bangumis"
(
    "id"              INTEGER PRIMARY KEY AUTOINCREMENT,
    "created_at"      DATETIME,
    "updated_at"      DATETIME,
    "deleted_at"      DATETIME,
    "subscription_id" INTEGER,
    "data_provider"   TEXT,
    "provider_id"     TEXT,
    "official_title"  TEXT,
    "first_air_data"  TEXT,
    "season"          INTEGER,
    "group_name"      TEXT,
    "poster"          TEXT,
    "overview"        TEXT,
    FOREIGN KEY ("subscription_id") REFERENCES "subscriptions" ("id")
);

CREATE TABLE "bangumi_episodes"
(
    "id"                   INTEGER PRIMARY KEY AUTOINCREMENT,
    "created_at"           DATETIME,
    "updated_at"           DATETIME,
    "deleted_at"           DATETIME,
    "subscription_file_id" INTEGER,
    "name"                 TEXT,
    "episode"              INTEGER,
    "poster"               TEXT,
    "overview"             TEXT,
    FOREIGN KEY ("subscription_file_id") REFERENCES "subscription_files" ("id")
);
