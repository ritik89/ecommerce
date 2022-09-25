CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

drop table if exists users;
drop table if exists merchants;
drop table if exists product_categories;
drop table if exists products;
drop table if exists product_offers;
drop table if exists product_category_offers;

create table users (
                       id uuid primary key,
                       name varchar
);

create table merchants (
                           id uuid primary key,
                           name varchar
);

create table product_categories (
                                    id uuid  primary key,
                                    name varchar,
                                    description varchar
);

create table products (
                          id uuid primary key,
                          name varchar not null,
                          description varchar,
                          product_category uuid references product_categories(id),
                          units int,
                          price float(3),
                          min_sale_price float(3)
);


create table product_offers (
                                id uuid primary key,
                                description varchar,
                                product_id uuid references products(id),
                                discount_percentage int,
                                cashback_amount float(3),
                                voucher_code varchar,
                                min_purchase_amount float(3),
                                max_discount_allowed float(3),
                                valid_from timestamptz,
                                valid_to timestamptz,
                                is_active bool
);

create table product_category_offers (
                                         id uuid primary key,
                                         description varchar,
                                         product_category_id uuid references product_categories(id),
                                         discount_percentage int,
                                         cashback_amount float(3),
                                         voucher_code varchar,
                                         min_purchase_amount float(3),
                                         max_discount_allowed float(3),
                                         valid_from timestamptz,
                                         valid_to timestamptz,
                                         is_active bool
);

create table cart_offers (
                             id uuid primary key,
                             description varchar,
                             discount_percentage int,
                             min_cart_value float(3),
                             max_discount_allowed float(3),
                             valid_from timestamptz,
                             valid_to timestamptz,
                             is_active bool
);

create table custom_offers (
                               id uuid primary key,
                               description varchar,
                               discount_percentage int,
                               cashback_amount float(3),
                               voucher_code varchar,
                               min_purchase_amount float(3),
                               max_discount_allowed float(3),
                               valid_from timestamptz not NUll,
                               valid_to timestamptz not null,
                               is_active bool not null
);

create table custom_offers_product_category_mapping(
                                                       id uuid primary key,
                                                       custom_offer_id uuid references custom_offers(id)
                                                           product_category_id references product_categories(id)
);

create table offer_provider (
                                id uuid primary key,
                                offer_type varchar not null,
                                offer_provider varchar not null
);

create table cart(
                     id uuid primary key,
                     user_id uuid references users(id)
)

create table cart_items(
                           user_id uuid references users(id) primary key,
                           product_id uuid references products(id),
                           offer_type varchar,
                           offer_id uuid
);

INSERT INTO public.product_categories
(id, "name", description)
VALUES(uuid_generate_v4(), 'shoes', 'shoes');

select * from product_categories;

INSERT INTO public.products
(id, "name", description, product_category, units, price, min_sale_price)
VALUES(uuid_generate_v4(), 'nike shoes 9', 'nike shoes 9', '5d95676a-0947-43a5-b526-ef1382b2f295', 5, 6.99, 3.00);

select * from products;

select now();

INSERT INTO public.product_offers
(id, description, product_id, discount_percentage, cashback_amount, min_purchase_amount, max_discount_allowed, valid_from, valid_to, is_active)
VALUES(uuid_generate_v4(), 'nike shoes 9', '81eecfa9-32ac-4cb7-b5cb-21886b4b9d49', 30, 0, 6.00, 3.00, '2022-09-25 17:55:37.959 +0530', '2022-09-30 17:55:37.959 +0530', true);

select * from product_offers;


INSERT INTO public.product_category_offers
(id, description, product_category_id, discount_percentage, cashback_amount, voucher_code, min_purchase_amount, max_discount_allowed, valid_from, valid_to, is_active)
VALUES(uuid_generate_v4(), 'shoes', '5d95676a-0947-43a5-b526-ef1382b2f295', 10, 0, '', 50, 20, '2022-09-25 17:55:37.959 +0530', '2022-09-30 17:55:37.959 +0530', true);

select * from product_category_offers;

select * from product_offers where discount_percentage = 40 or description = 'nike shoes 9';

INSERT INTO public.users
(id, "name")
VALUES(uuid_generate_v4(), 'ritik1');
INSERT INTO public.users
(id, "name")
VALUES(uuid_generate_v4(), 'ritik2');
INSERT INTO public.users
(id, "name")
VALUES(uuid_generate_v4(), 'ritik3');

select * from users;

INSERT INTO public.cart
(id, user_id)
VALUES(uuid_generate_v4(), '1b57d3c9-858c-4091-b6d6-728401764683');
INSERT INTO public.cart
(id, user_id)
VALUES(uuid_generate_v4(), '4b114942-53b8-4ac7-aed8-b0c2913b98b4');
INSERT INTO public.cart
(id, user_id)
VALUES(uuid_generate_v4(), '35dd14be-3d20-42b6-b839-68abd830db72');

INSERT INTO public.cart_items
(user_id, product_id)
VALUES('1b57d3c9-858c-4091-b6d6-728401764683', '81eecfa9-32ac-4cb7-b5cb-21886b4b9d49');

select * from cart;
select * from cart_items;
SELECT id, user_id
FROM public.cart where user_id = '1b57d3c9-858c-4091-b6d6-728401764683';

select cart_items.user_id, cart_items.product_id from cart left join cart_items on cart.user_id = cart_items.user_id where cart.user_id = '1b57d3c9-858c-4091-b6d6-728401764683';

