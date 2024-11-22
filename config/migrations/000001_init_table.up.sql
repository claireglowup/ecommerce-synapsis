-- users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,           
    name VARCHAR(255) NOT NULL,     
    email VARCHAR(255) UNIQUE NOT NULL, 
    password VARCHAR(255) NOT NULL  
);

-- products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,         
    name VARCHAR(255) NOT NULL,    
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,            
    description TEXT                
);

-- carts table
CREATE TABLE IF NOT EXISTS carts (
    id SERIAL PRIMARY KEY,          
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP     
);

-- cart_items table (Relasi Many-to-Many antara Cart dan Product)
CREATE TABLE IF NOT EXISTS cart_items (
    id SERIAL PRIMARY KEY,           
    cart_id INT REFERENCES carts(id) ON DELETE CASCADE,  
    product_id INT REFERENCES products(id) ON DELETE CASCADE, 
    quantity INT NOT NULL,           
    UNIQUE(cart_id, product_id)      
);

-- Menambahkan Indeks untuk kinerja
CREATE INDEX IF NOT EXISTS idx_user_id ON carts(user_id);
CREATE INDEX IF NOT EXISTS idx_product_id ON cart_items(product_id);
