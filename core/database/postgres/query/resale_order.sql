-- name: SelectOrderByCPFAndOrderID :one
SELECT ro.id
FROM cercle_test.resale_order ro
JOIN cercle_test.users u ON u.id = ro.fk_users_id
WHERE ro.id = @order_id
  AND u.document_number = @document_number
  AND ro.deleted_at IS NULL;

-- name: SelectOrderItemsByCPFAndOrderID :many
SELECT
    roi.id,
    roi.fk_resale_order_id,
    roi.sku,
    roi.name,
    roi.quantity,
    roi.amount_value,
    roi.shipping_code,
    roi.shipping_status,
    roi.delivered_at,
    roi.created_at,
    roi.updated_at,
    roi.deleted_at
FROM cercle_test.resale_order_item roi
JOIN cercle_test.resale_order ro ON ro.id = roi.fk_resale_order_id
JOIN cercle_test.users u ON u.id = ro.fk_users_id
WHERE ro.id = @order_id
  AND u.document_number = @document_number
  AND ro.deleted_at IS NULL
  AND roi.deleted_at IS NULL
ORDER BY roi.created_at ASC;
