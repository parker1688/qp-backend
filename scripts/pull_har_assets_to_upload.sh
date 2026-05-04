#!/usr/bin/env bash
set -euo pipefail

# Pull all /api/upload/har-assets/* resources referenced by DB text columns,
# then materialize files into backend upload directory so existing DB paths work.

DB_DEFAULTS_FILE="${DB_DEFAULTS_FILE:-/etc/mysql/debian.cnf}"
DB_NAME="${DB_NAME:-hy_game}"
BACKEND_CWD="${BACKEND_CWD:-}"
TMP_DIR="${TMP_DIR:-/tmp/qp-har-assets}"

mkdir -p "$TMP_DIR"
URL_FILE="$TMP_DIR/har_asset_urls.txt"
SQL_FILE="$TMP_DIR/har_asset_union.sql"

if [[ -z "$BACKEND_CWD" ]]; then
  WEB_PID="$(pgrep -f 'cmd/web/main.go' | head -n1 || true)"
  API_PID="$(pgrep -f 'cmd/api/main.go' | head -n1 || true)"
  PID="${WEB_PID:-$API_PID}"
  if [[ -z "$PID" ]]; then
    echo "[ERR] Cannot detect backend process cwd. Set BACKEND_CWD explicitly."
    exit 1
  fi
  BACKEND_CWD="$(readlink -f "/proc/$PID/cwd")"
fi

UPLOAD_DIR="$BACKEND_CWD/upload/har-assets"
mkdir -p "$UPLOAD_DIR"

echo "[INFO] DB: $DB_NAME"
echo "[INFO] Backend cwd: $BACKEND_CWD"
echo "[INFO] Upload dir: $UPLOAD_DIR"

# Build a UNION SQL over all textual columns that may contain har asset paths.
mysql --defaults-file="$DB_DEFAULTS_FILE" -N -B -e "
SELECT CONCAT(
  'SELECT DISTINCT ',
  CONCAT('`', COLUMN_NAME, '`'),
  ' AS url FROM ',
  CONCAT('`', TABLE_SCHEMA, '`.`', TABLE_NAME, '`'),
  ' WHERE ',
  CONCAT('`', COLUMN_NAME, '`'),
  \" LIKE '/api/upload/har-assets/%'\"
)
FROM information_schema.COLUMNS
WHERE TABLE_SCHEMA='${DB_NAME}'
  AND DATA_TYPE IN ('char','varchar','tinytext','text','mediumtext','longtext');
" > "$SQL_FILE"

if [[ ! -s "$SQL_FILE" ]]; then
  echo "[ERR] No candidate text columns found."
  exit 1
fi

# Merge as one query and extract unique URL list.
{
  awk 'NR==1{print;next}{print "UNION\n"$0}' "$SQL_FILE"
} | mysql --defaults-file="$DB_DEFAULTS_FILE" -N -B "$DB_NAME" \
  | sed '/^$/d' \
  | sort -u > "$URL_FILE"

COUNT="$(wc -l < "$URL_FILE" | tr -d ' ')"
echo "[INFO] Found $COUNT DB URLs under /api/upload/har-assets/"

if [[ "$COUNT" -eq 0 ]]; then
  echo "[DONE] Nothing to download."
  exit 0
fi

ok=0
fail=0

while IFS= read -r db_url; do
  rel="${db_url#/api/upload/har-assets/}"
  dst="$UPLOAD_DIR/$rel"
  mkdir -p "$(dirname "$dst")"

  if [[ -s "$dst" ]]; then
    ok=$((ok + 1))
    continue
  fi

  src_https="https://$rel"
  src_http="http://$rel"

  if curl -kfsSL "$src_https" -o "$dst"; then
    ok=$((ok + 1))
    continue
  fi

  if curl -kfsSL "$src_http" -o "$dst"; then
    ok=$((ok + 1))
    continue
  fi

  echo "[WARN] download failed: $db_url"
  rm -f "$dst"
  fail=$((fail + 1))
done < "$URL_FILE"

echo "[DONE] success=$ok failed=$fail total=$COUNT"
if [[ "$fail" -gt 0 ]]; then
  echo "[INFO] Failed URL list:"
  grep -v '^$' "$URL_FILE" > "$TMP_DIR/all_urls.txt"
  find "$UPLOAD_DIR" -type f | sed "s#^$UPLOAD_DIR/##" | awk '{print "/api/upload/har-assets/"$0}' | sort -u > "$TMP_DIR/success_urls.txt"
  comm -23 "$TMP_DIR/all_urls.txt" "$TMP_DIR/success_urls.txt" || true
fi
