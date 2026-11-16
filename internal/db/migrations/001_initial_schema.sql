-- Instances table
CREATE TABLE IF NOT EXISTS instances (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    aws_instance_id VARCHAR(50),
    instance_type VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Allowed instance types
CREATE TABLE IF NOT EXISTS allowed_instance_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(20) UNIQUE NOT NULL,
    vcpu INT NOT NULL,
    memory_gb DECIMAL(4,2) NOT NULL
);

-- Insert allowed instance types
INSERT INTO allowed_instance_types (type_name, vcpu, memory_gb) VALUES
    ('t2.micro', 1, 1.00),
    ('t2.small', 1, 2.00),
    ('t2.medium', 2, 4.00),
    ('t3.small', 2, 2.00),
    ('t3.medium', 2, 4.00)
ON CONFLICT (type_name) DO NOTHING;

-- Indexes
CREATE INDEX IF NOT EXISTS idx_instances_status ON instances(status);
