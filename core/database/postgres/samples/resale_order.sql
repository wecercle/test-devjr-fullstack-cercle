-- Resale Order sample data
INSERT INTO cercle_test.resale_order (
    id,
    fk_retailer_id,
    fk_users_id,
    external_id,
    code,
    date,
    current_status,
    delivery_address_zip_code,
    delivery_address_street,
    delivery_address_number,
    delivery_address_district,
    delivery_address_city,
    delivery_address_state,
    delivery_address_country,
    delivery_address_receiver_name,
    payment_status,
    amount_subtotal,
    amount_freight,
    amount_interest,
    amount_total,
    created_at,
    updated_at
) VALUES (
    '6ae28fac-280e-4879-bbe3-9277decb4a06', -- id
    '550e8400-e29b-41d4-a716-446655440002', -- fk_retailer_id (LA ABACATEIRA)
    '0e4a949c-b05c-4778-aeaa-a027ff3d7751', -- fk_users_id (Talia Ekiu)
    'or_TestJY4I8c6w0AB', -- external_id
    'CD-T0001-00001', -- code
    '2026-06-08 16:31:18', -- date
    'PAID', -- current_status
    '13253330', -- delivery_address_zip_code
    'Rua João Pellizzer', -- delivery_address_street
    '111', -- delivery_address_number
    'Lugarzinho', -- delivery_address_district
    'Itatiba', -- delivery_address_city
    'SP', -- delivery_address_state
    'BR', -- delivery_address_country
    'Talia Ekiu', -- delivery_address_receiver_name
    'pagamento-aprovado', -- payment_status
    '538.00', -- amount_subtotal
    '30.00', -- amount_freight
    '9.52', -- amount_interest
    '577.52', -- amount_total
    '2026-06-01 16:31:18', -- created_at
    '2026-06-05 16:31:18' -- updated_at
), (
    '7ae28fac-280e-4879-bbe3-9277decb4a07', -- id
    '550e8400-e29b-41d4-a716-446655440003', -- fk_retailer_id (SPORTS STORE)
    '1e4a949c-b05c-4778-aeaa-a027ff3d7752', -- fk_users_id (Helena Nogueira)
    'or_TestJY4I8c6w0AC', -- external_id
    'SPR-2026-00002', -- code
    '2026-06-10 09:15:00', -- date
    'PAID', -- current_status
    '04538000', -- delivery_address_zip_code
    'Avenida Paulista', -- delivery_address_street
    '1000', -- delivery_address_number
    'Bela Vista', -- delivery_address_district
    'São Paulo', -- delivery_address_city
    'SP', -- delivery_address_state
    'BR', -- delivery_address_country
    'Helena Nogueira', -- delivery_address_receiver_name
    'pagamento-aprovado', -- payment_status
    '349.90', -- amount_subtotal
    '19.90', -- amount_freight
    '0.00', -- amount_interest
    '369.80', -- amount_total
    '2026-06-13 09:15:12.000000', -- created_at
    '2026-06-14 09:17:40.000000' -- updated_at
), (
    '8ae28fac-280e-4879-bbe3-9277decb4a08', -- id
    '550e8400-e29b-41d4-a716-446655440004', -- fk_retailer_id (MM Store)
    '2e4a949c-b05c-4778-aeaa-a027ff3d7753', -- fk_users_id (Bianca Sousa)
    'or_TestJY4I8c6w0AD', -- external_id
    'MM-2026-00003', -- code
    '2026-06-11 11:42:00', -- date
    'CANCELLED', -- current_status
    '13087012', -- delivery_address_zip_code
    'Rua da Abolição', -- delivery_address_street
    '220', -- delivery_address_number
    'Cambuí', -- delivery_address_district
    'Campinas', -- delivery_address_city
    'SP', -- delivery_address_state
    'BR', -- delivery_address_country
    'Bianca Sousa', -- delivery_address_receiver_name
    'cancelado', -- payment_status
    '1299.00', -- amount_subtotal
    '0.00', -- amount_freight
    '129.90', -- amount_interest
    '1428.90', -- amount_total
    '2026-06-13 11:43:10.000000', -- created_at
    '2026-06-13 11:45:55.000000' -- updated_at
), (
    '9ae28fac-280e-4879-bbe3-9277decb4a09', -- id
    '550e8400-e29b-41d4-a716-446655440005', -- fk_retailer_id (Shield Hero)
    '3e4a949c-b05c-4778-aeaa-a027ff3d7754', -- fk_users_id (Camila Freitas)
    'or_TestJY4I8c6w0AE', -- external_id
    'SHD-2026-00004', -- code
    '2026-06-12 14:05:00', -- date
    'PAID', -- current_status
    '30140071', -- delivery_address_zip_code
    'Rua Espírito Santo', -- delivery_address_street
    '501', -- delivery_address_number
    'Funcionários', -- delivery_address_district
    'Belo Horizonte', -- delivery_address_city
    'MG', -- delivery_address_state
    'BR', -- delivery_address_country
    'Camila Freitas', -- delivery_address_receiver_name
    'pagamento-aprovado', -- payment_status
    '589.50', -- amount_subtotal
    '22.50', -- amount_freight
    '0.00', -- amount_interest
    '612.00', -- amount_total
    '2026-06-12 14:05:30.000000', -- created_at
    '2026-06-12 14:06:05.000000' -- updated_at
), (
    'aae28fac-280e-4879-bbe3-9277decb4a10', -- id
    '550e8400-e29b-41d4-a716-446655440002', -- fk_retailer_id (LA ABACATEIRA)
    '4e4a949c-b05c-4778-aeaa-a027ff3d7755', -- fk_users_id (Larissa Melo)
    'or_TestJY4I8c6w0AF', -- external_id
    'LAB-2026-00005', -- code
    '2026-06-13 18:20:00', -- date
    'PAID', -- current_status
    '22250040', -- delivery_address_zip_code
    'Rua Voluntários da Pátria', -- delivery_address_street
    '77', -- delivery_address_number
    'Botafogo', -- delivery_address_district
    'Rio de Janeiro', -- delivery_address_city
    'RJ', -- delivery_address_state
    'BR', -- delivery_address_country
    'Larissa Melo', -- delivery_address_receiver_name
    'pagamento-aprovado', -- payment_status
    '249.99', -- amount_subtotal
    '15.00', -- amount_freight
    '0.00', -- amount_interest
    '264.99', -- amount_total
    '2026-06-13 18:20:10.000000', -- created_at
    '2026-06-13 18:21:00.000000' -- updated_at
), (
    '66666666-6666-6666-6666-666666666666', -- id
    '550e8400-e29b-41d4-a716-446655440003', -- fk_retailer_id (SPORTS STORE)
    '1e4a949c-b05c-4778-aeaa-a027ff3d7752', -- fk_users_id (Helena Nogueira)
    'or_PaidManualTest1', -- external_id
    'PAID-2026-00006', -- code
    '2026-06-17 17:35:00', -- date
    'PAID', -- current_status
    '04538000', -- delivery_address_zip_code
    'Avenida Paulista', -- delivery_address_street
    '1000', -- delivery_address_number
    'Bela Vista', -- delivery_address_district
    'São Paulo', -- delivery_address_city
    'SP', -- delivery_address_state
    'BR', -- delivery_address_country
    'Helena Nogueira', -- delivery_address_receiver_name
    'pagamento-aprovado', -- payment_status
    '199.90', -- amount_subtotal
    '19.90', -- amount_freight
    '0.00', -- amount_interest
    '219.80', -- amount_total
    '2026-06-17 17:35:00.000000', -- created_at
    '2026-06-17 17:35:00.000000' -- updated_at
) ON CONFLICT (id) DO NOTHING;


