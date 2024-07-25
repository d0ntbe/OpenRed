#!/usr/bin/env bash

OpenRed="OpenRed.go"

strings=('target1.com' 'target2.ru' 'target3.gg' )

for string in "${strings[@]}"; do
  echo "Testing target : $string"
  go run "$OpenRed" "$string"
done