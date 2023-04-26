CREATE TABLE IF NOT EXISTS orders
(
    order_uid          varchar PRIMARY KEY NOT NULL,
    track_number       varchar,
    entry              varchar,
    locale             varchar,
    internal_signature varchar,
    customer_id        varchar,
    delivery_service   varchar,
    shardkey           varchar,
    sm_id              bigint,
    date_created       timestamp,
    oof_shard          varchar
);

CREATE TABLE IF NOT EXISTS deliveries
(
    id serial PRIMARY KEY NOT NULL,
    order_uid   varchar NOT NULL,
    name        varchar,
    phone       varchar,
    zip         varchar,
    city        varchar,
    address     varchar,
    region      varchar,
    email       varchar,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS payments
(
    order_uid     varchar NOT NULL,
    transaction   varchar PRIMARY KEY NOT NULL,
    request_id    varchar,
    currency      varchar,
    provider      varchar,
    amount        bigint,
    payment_dt    bigint,
    bank          varchar,
    delivery_cost bigint,
    goods_total   bigint,
    custom_fee    bigint,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS items
(
    item_id	     serial PRIMARY KEY NOT NULL,
    order_uid    varchar NOT NULL,
    chrt_id      bigint,
    track_number varchar,
    price        bigint,
    rid          varchar,
    name         varchar,
    sale         bigint,
    size         varchar,
    total_price  bigint,
    nm_id        bigint,
    brand        varchar,
    status       bigint,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);