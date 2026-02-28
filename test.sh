#!/bin/bash
echo "Depositing 100 USDT:"
curl -s -X POST http://127.0.0.1:8000/api/admin/v1/asset/asset/subAmount \
  -H "Content-Type: application/json" \
  -d '{"userId": 1, "symbol": "USDT", "amount": 100, "recordType": 1, "remark": "Deposit"}'
echo -e "\n\nWithdrawing 200 USDT:"
curl -s -X POST http://127.0.0.1:8000/api/admin/v1/asset/asset/subAmount \
  -H "Content-Type: application/json" \
  -d '{"userId": 1, "symbol": "USDT", "amount": -200, "recordType": 2, "remark": "Withdraw"}'
echo ""
