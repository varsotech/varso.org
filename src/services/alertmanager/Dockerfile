FROM prom/alertmanager

COPY ./src/services/alertmanager/ /etc/alertmanager/
COPY ./scripts/ /etc/scripts/

ENTRYPOINT sh -c "/etc/scripts/apply_env.sh /etc/alertmanager/config.yaml" && /bin/alertmanager --config.file=/etc/alertmanager/config.yaml --storage.path=/alertmanager