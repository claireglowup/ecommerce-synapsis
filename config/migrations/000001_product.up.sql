CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,                -- ID unik untuk setiap produk
    name VARCHAR(100) NOT NULL,           -- Nama produk
    description TEXT,                     -- Deskripsi produk
    price DECIMAL(10, 2) NOT NULL,        -- Harga produk (contoh: 9999.99)
    stock INT DEFAULT 0,                  -- Jumlah stok produk, default 0
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Tanggal dan waktu produk ditambahkan
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- Tanggal dan waktu terakhir kali produk diperbarui
);
