#!/bin/bash

BASE_URL="http://localhost:8000/api/v2/records"

echo "Creating business histories..."

#
# BUSINESS 1 — ACME PLUMBING
#

echo "Creating Acme Plumbing history..."

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "business_name": "Acme Plumbing LLC",
  "industry": "Plumbing",
  "employee_count": "10",
  "annual_payroll": "500000",
  "general_liability_limit": "1000000",
  "business_hours": "9am-5pm"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "employee_count": "15",
  "annual_payroll": "750000"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "business_hours": "7am-7pm"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "general_liability_limit": "2000000"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "employee_count": "25",
  "annual_payroll": "1500000"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "delivery_services": "yes"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "employee_count": "40"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "business_hours": "24/7"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "general_liability_limit": "5000000"
}'

sleep 1

curl -s -X POST $BASE_URL/1 \
-H "Content-Type: application/json" \
-d '{
  "hazardous_materials": "yes"
}'

sleep 1

#
# BUSINESS 2 — BLUE SKY BAKERY
#

echo ""
echo "Creating Blue Sky Bakery history..."

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "business_name": "Blue Sky Bakery",
  "industry": "Bakery",
  "employee_count": "5",
  "annual_payroll": "200000",
  "general_liability_limit": "500000",
  "business_hours": "6am-3pm"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "employee_count": "8"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "delivery_services": "yes"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "annual_payroll": "350000"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "business_hours": "5am-8pm"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "employee_count": "12"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "general_liability_limit": "1000000"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "catering_services": "yes"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "employee_count": "20",
  "annual_payroll": "750000"
}'

sleep 1

curl -s -X POST $BASE_URL/2 \
-H "Content-Type: application/json" \
-d '{
  "overnight_baking": "yes"
}'

echo ""
echo "Done seeding historical business records."