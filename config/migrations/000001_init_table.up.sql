-- Aktifkan ekstensi pgcrypto untuk UUID
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),           
    name VARCHAR(255) NOT NULL,     
    email VARCHAR(255) UNIQUE NOT NULL, 
    password VARCHAR(255) NOT NULL  
);

-- products table
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),         
    name VARCHAR(255) NOT NULL,    
    price BIGINT NOT NULL,
    stock INT NOT NULL,
    category VARCHAR(40) NOT NULL,            
    description TEXT                
);

-- carts table
CREATE TABLE IF NOT EXISTS carts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),          
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     
);

-- cart_items table (Relasi Many-to-Many antara Cart dan Product)
CREATE TABLE IF NOT EXISTS cart_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),           
    cart_id UUID REFERENCES carts(id) ON DELETE CASCADE,  
    product_id UUID REFERENCES products(id) ON DELETE CASCADE, 
    quantity INT NOT NULL,           
    UNIQUE(cart_id, product_id)      
);

-- Menambahkan Indeks untuk kinerja
CREATE INDEX IF NOT EXISTS idx_user_id ON carts(user_id);
CREATE INDEX IF NOT EXISTS idx_product_id ON cart_items(product_id);

-- seed
INSERT INTO products (name, price, stock, category, description)
VALUES
    ('Nike SB Blazer', 1200000, 2, 'shoes', 'Nike shoes for Skateboarding'),
    ('Vans Old Skool Skate', 1200000, 3, 'shoes', 'Vans Old Skool For Skateboarding'),
    ('Deck Baker T-Funk', 1200000, 3, 'deck', 'Deck Size 8.25'),
    ('Independent Stage 11 Trucks', 800000, 4, 'trucks', 'Premium Independent trucks for stability and grind durability'),
    ('Spitfire Formula Four Wheels', 700000, 5, 'wheels', 'High-performance Spitfire wheels for smooth riding'),
    ('Bones Reds Bearings', 300000, 10, 'bearings', 'Affordable and high-quality bearings from Bones'),
    ('Element Section Complete', 2500000, 2, 'complete', 'Complete skateboard by Element with quality hardware'),
    ('Santa Cruz Screaming Hand Deck', 1300000, 3, 'deck', 'Classic Santa Cruz deck with iconic graphics'),
    ('Etnies Marana Skate Shoes', 1400000, 4, 'shoes', 'Durable Etnies Marana for impact resistance and comfort'),
    ('Thunder Polished Trucks', 900000, 6, 'trucks', 'Durable and lightweight Thunder trucks for all types of skating');


   
