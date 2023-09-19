create table orders (
    id                bigserial not null primary key,
    OrderUID          varchar(128),
    Entry             varchar(128),
    InternalSignature varchar(128),
    payment_id_fk     varchar(128),
    Locale            varchar(128),
    CustomerID        varchar(128),
    TrackNumber       varchar(128),
    DeliveryService   varchar(128),
    Shardkey          varchar(128),
    SmID              int,
    OffShard          varchar(128)
);

create table delivery (
    id     bigserial not null primary key,
    Name   varchar(50),
    Phone  varchar(50),
    Zip    varchar(128),
    City   varchar(128),
    Addres varchar(128),
    Region varchar(128),
    Email  varchar(128)
);

create table payments (
    id           bigserial not null primary key,
    Transaction  varchar(256),
    Currency     varchar(128),
    Provider     varchar(128),
    Amount       int,
    PaymentDt    int,
    Bank         varchar(128),
    DeliveryCost int,
    GoodsTotal   int,
    CustomFee    int
);

create table items (
    id         bigserial not null primary key,
    ChrtID     int,
    Price      int,
    Rid        varchar(256),
    Name       varchar(128),
    Sale       int,
    Size       varchar(128),
    TotalPrice int,
    NmID       int,
    Brand      varchar(128),
    Status     int
);

