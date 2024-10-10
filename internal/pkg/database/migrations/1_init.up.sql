CREATE TABLE Subscription
(
    id                INTEGER PRIMARY KEY AUTOINCREMENT,
    link              TEXT,
    is_aggregate      BOOLEAN,
    parser            TEXT,
    black_list_filter TEXT,
    white_list_filter TEXT,
    is_enabled        BOOLEAN,
    created_at        DATETIME,
    updated_at        DATETIME,
    deleted_at        DATETIME
);

CREATE TABLE SubscriptionFile
(
    id                 INTEGER PRIMARY KEY AUTOINCREMENT,
    subscription_id    INTEGER,
    name               TEXT,
    link               TEXT,
    is_download        BOOLEAN,
    downloader_task_id TEXT,
    created_at         DATETIME,
    updated_at         DATETIME,
    deleted_at         DATETIME,
    FOREIGN KEY (subscription_id) REFERENCES Subscription (id)
);


CREATE TABLE Bangumi
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    subscription_id INTEGER,
    data_provider   TEXT,
    provider_id     TEXT,
    official_title  TEXT,
    first_air_data  TEXT,
    season          INTEGER,
    group_name      TEXT,
    poster          TEXT,
    overview        TEXT,
    created_at      DATETIME,
    updated_at      DATETIME,
    deleted_at      DATETIME,
    FOREIGN KEY (subscription_id) REFERENCES Subscription (id)
);

CREATE TABLE BangumiEpisode
(
    id                   INTEGER PRIMARY KEY AUTOINCREMENT,
    subscription_file_id INTEGER,
    name                 TEXT,
    episode              INTEGER,
    poster               TEXT,
    overview             TEXT,
    created_at           DATETIME,
    updated_at           DATETIME,
    deleted_at           DATETIME,
    FOREIGN KEY (subscription_file_id) REFERENCES SubscriptionFile (id)
);

