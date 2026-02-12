#!/bin/bash

echo "=== Testing User API ==="
echo ""

# Test 1: Create a user
echo "1. Creating a new user..."
curl -X POST http://localhost:4000/order \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
echo -e "\n"

# Test 2: Create another user
echo "2. Creating another user..."
curl -X POST http://localhost:4000/order \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com",
    "password": "password456"
  }'
echo -e "\n"

# Test 3: Get all users
echo "3. Getting all users..."
curl -X GET http://localhost:4000/users
echo -e "\n"

echo ""
echo "=== Test Complete ==="
