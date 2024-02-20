CREATE VIEW "shopping_cart_with_cart_items" AS
SELECT
    "shopping_carts".*,
    JSONB_AGG("cart_items".*) AS "cart_items"
FROM
    "shopping_carts"
LEFT JOIN "cart_items" ON "shopping_carts"."id" = "cart_items"."cart_id"
GROUP BY
    "shopping_carts"."id";

CREATE VIEW "order_with_order_items" AS
SELECT
    "orders".*,
    JSONB_AGG("order_items".*) AS "order_items"
FROM
    "orders"
LEFT JOIN "order_items" ON "orders"."id" = "order_items"."order_id"
GROUP BY
    "orders"."id";