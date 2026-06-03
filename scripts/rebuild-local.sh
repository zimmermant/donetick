#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
FRONTEND_DIR="${FRONTEND_DIR:-"${REPO_ROOT}/../donetick-frontend"}"
GOOS="${GOOS:-linux}"

if [[ ! -d "${FRONTEND_DIR}" ]]; then
  echo "Frontend directory not found: ${FRONTEND_DIR}" >&2
  echo "Set FRONTEND_DIR=/path/to/donetick-frontend and try again." >&2
  exit 1
fi

if [[ -z "${GOARCH:-}" ]]; then
  case "$(uname -m)" in
    arm64|aarch64)
      GOARCH="arm64"
      ;;
    x86_64|amd64)
      GOARCH="amd64"
      ;;
    *)
      echo "Unsupported host architecture: $(uname -m)" >&2
      echo "Set GOARCH manually, for example GOARCH=amd64 or GOARCH=arm64." >&2
      exit 1
      ;;
  esac
fi

export GOOS GOARCH CGO_ENABLED="${CGO_ENABLED:-0}"
export GOCACHE="${GOCACHE:-/private/tmp/donetick-go-build-cache}"

echo "Building frontend from ${FRONTEND_DIR}"
(
  cd "${FRONTEND_DIR}"
  npm run build
)

echo "Syncing frontend build into backend frontend/dist"
mkdir -p "${REPO_ROOT}/frontend/dist"
rsync -a --delete "${FRONTEND_DIR}/dist/" "${REPO_ROOT}/frontend/dist/"

echo "Building Donetick binary for ${GOOS}/${GOARCH}"
(
  cd "${REPO_ROOT}"
  go build -o donetick .
)

echo "Building Docker image from docker-compose.yaml"
(
  cd "${REPO_ROOT}"
  docker compose build
)

echo
echo "Done. Start Donetick with:"
echo "  docker compose up -d"
echo
echo "Then open:"
echo "  http://localhost:2021"
