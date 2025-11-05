-- Add category reference to products
ALTER TABLE products
    ADD COLUMN IF NOT EXISTS category_id INTEGER NULL REFERENCES categories(id);

CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);

