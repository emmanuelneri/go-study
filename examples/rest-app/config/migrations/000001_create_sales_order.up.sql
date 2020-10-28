create table sales_order
(
    id       integer primary key generated always as identity,
    customer VARCHAR(200),
    total    numeric(19, 2)
);