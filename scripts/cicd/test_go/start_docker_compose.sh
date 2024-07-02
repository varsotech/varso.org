set -e

# Install yq - used for parsing docker-compose.yaml
if ! command -v yq &> /dev/null; then
    wget https://github.com/mikefarah/yq/releases/download/${VERSION}/${BINARY} -O /usr/bin/yq && chmod +x /usr/bin/yq
fi

echo Image Template: ${CONTAINER_REGISTRY}/${GIT_FULL_REPO_NAME}/service:${DOCKER_IMAGE_TAG}

# Boot up services
docker compose up -d