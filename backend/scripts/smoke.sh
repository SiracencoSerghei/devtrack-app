#!/bin/bash

set -e

BASE_URL="http://localhost:8080"

echo "→ Health check..."
curl -s -f "$BASE_URL/health" | grep -q "OK" || {
  echo "❌ Health failed"
  exit 1
}

echo "→ Root check..."
curl -s -f "$BASE_URL/" | grep -q "DevTrack" || {
  echo "❌ Root failed"
  exit 1
}

echo "→ Login route check (expected 400/401 ok)"
curl -s -o /dev/null -w "%{http_code}" -X POST "$BASE_URL/api/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"x","password":"x"}' | grep -E "400|401" >/dev/null || {
  echo "❌ Login route broken"
  exit 1
}

echo "✅ SMOKE TEST PASSED"

# chmod +x backend/scripts/smoke.sh