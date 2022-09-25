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
                     user_id uuid references users(id),
                     cart_total float(3)
)

create table cart_products(
                              id uuid primary key,
                              cart_id uuid references cart(id),
                              product_id uuid references products(id),
                              offer_type varchar not null,
                              offer_id uuid not null
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
