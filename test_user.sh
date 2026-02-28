#!/bin/bash
echo "--- Testing Phone Registration ---"
curl -s -X POST http://127.0.0.1:8000/api/v1/user/user/register -H "Content-Type: application/json" -d '{"signType": "PHONE", "phone": "13800138008", "loginPassword": "password123"}'
echo -e "\n--- Testing Web3 Registration ---"
curl -s -X POST http://127.0.0.1:8000/api/v1/user/user/register -H "Content-Type: application/json" -d '{"signType": "ADDRESS", "address": "0x8234567890abcdef"}'
echo -e "\n--- Testing Phone Login ---"
curl -s -X POST http://127.0.0.1:8000/api/v1/user/user/login -H "Content-Type: application/json" -d '{"signType": "PHONE", "phone": "13800138008", "loginPassword": "password123"}'
echo -e "\n--- Testing Web3 Login ---"
curl -s -X POST http://127.0.0.1:8000/api/v1/user/user/login -H "Content-Type: application/json" -d '{"signType": "ADDRESS", "address": "0x8234567890abcdef"}'
echo ""
