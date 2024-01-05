CREATE SEQUENCE product_id_seq;

create table product (
	id bigint primary key not null default nextval('product_id_seq'),
    name varchar(200) not null,
    description varchar(300),
    price decimal(10,2),
    quantity bigint
)

ALTER SEQUENCE product_id_seq
OWNED BY product.id;