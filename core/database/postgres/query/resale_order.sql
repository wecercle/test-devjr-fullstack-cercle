-- name: GetOrderItemsByCPFAndOrderID :many
SELECT roi.id, roi.fk_resale_order_id, roi.sku, roi.name, roi.quantity, roi.amount_value, roi.shipping_code, roi.shipping_status, roi.delivered_at
FROM cercle_test.resale_order_item roi
JOIN cercle_test.resale_order ro ON ro.id = roi.fk_resale_order_id
JOIN cercle_test.users u ON u.id = ro.fk_users_id
WHERE u.document_number = @document_number
  AND ro.id = @order_id
  AND roi.deleted_at IS NULL
  AND ro.deleted_at IS NULL
ORDER BY roi.created_at ASC;

-- name: GetOrderItemForCancellation :one
SELECT roi.id, roi.shipping_status, roi.delivered_at
FROM cercle_test.resale_order_item roi
JOIN cercle_test.resale_order ro ON ro.id = roi.fk_resale_order_id
JOIN cercle_test.users u ON u.id = ro.fk_users_id
WHERE u.document_number = @document_number
  AND ro.id = @order_id
  AND roi.id = @item_id
  AND roi.deleted_at IS NULL
  AND ro.deleted_at IS NULL;

-- name: CancelOrderItem :exec
UPDATE cercle_test.resale_order_item
SET shipping_status = 'RETURNED', updated_at = NOW()
WHERE id = @item_id AND deleted_at IS NULL;