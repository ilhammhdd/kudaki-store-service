-- retrieve items to be updated from owner order approved
SELECT
  ci.total_item AS ci_total_item,
  i.uuid AS i_uuid,
  i.amount AS i_amount,
  sf.uuid AS sf_uuid,
  sf.total_item AS sf_total_item
FROM
  kudaki_order.owner_orders oo
  JOIN kudaki_order.orders o ON oo.order_uuid = o.uuid
  JOIN kudaki_rental.carts c ON o.cart_uuid = o.cart_uuid
  JOIN kudaki_rental.cart_items ci ON c.uuid = ci.cart_uuid
  JOIN kudaki_store.items i ON ci.item_uuid = i.uuid
  JOIN kudaki_store.storefronts sf ON i.storefront_uuid = sf.uuid
WHERE
  oo.uuid = "7877e8d4-cee6-48b0-abc5-30029078d47a";
-- update all item quantity in items table
INSERT INTO kudaki_store.storefronts(uuid, total_item) VALUES(?, ?) ON DUPLICATE KEY UPDATE total_item = ?;
-- update all item quantity in storefronts table
INSERT INTO kudaki_store.items(uuid, amount) VALUES(?, ?) ON DUPLICATE KEY UPDATE amount = ?;