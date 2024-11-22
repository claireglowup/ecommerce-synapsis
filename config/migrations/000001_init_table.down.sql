-- Hapus tabel cart_items terlebih dahulu karena memiliki dependensi pada cart dan product
DROP TABLE IF EXISTS cart_items CASCADE;

-- Hapus tabel carts setelah cart_items
DROP TABLE IF EXISTS carts CASCADE;

-- Hapus tabel products
DROP TABLE IF EXISTS products CASCADE;

-- Hapus tabel users
DROP TABLE IF EXISTS users CASCADE;
