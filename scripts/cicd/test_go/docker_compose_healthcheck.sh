
check_service_health() {
    local service_name="$1"
    local start_time=$(date +%s)
    local end_time=$((start_time + 60))  # 60 seconds timeout

    while true; do
        echo "Checking $service_name health.."
        local health_status=$(docker-compose ps -q "$service_name" | xargs docker inspect --format='{{.State.Health.Status}}' 2>/dev/null)

        if [[ "$health_status" == "healthy" ]]; then
            echo "$service_name is healthy"
            break
        fi
        
        local current_time=$(date +%s)
        if [[ $current_time -ge $end_time ]]; then
            docker compose logs $service_name
            echo "Timeout: $service_name did not become healthy within 1 minute."
            exit 1
            break
        fi

        sleep 1  # Adjust sleep time if needed
    done
}


# Check services health
services=$(docker-compose config --services)
for service in $services; do
    check_service_health "$service"
done
