#!/usr/bin/env bash
# upload
curl -X POST http://localhost:8080/documents -d @fake-data.json -H "Content-Type: application/json"

# search
curl http://localhost:8080/search?query=exercitation+est+officia