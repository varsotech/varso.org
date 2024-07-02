FROM grafana/loki:2.9.4

WORKDIR /varsoloki

COPY ./scripts/apply_env.sh /varsoloki/
COPY ./src/services/loki/entrypoint.sh /varsoloki/
COPY ./src/services/loki/minio-config.yaml.template /etc/loki/

ENTRYPOINT [ "/varsoloki/entrypoint.sh" ]
CMD ["-config.file=/etc/loki/minio-config.yaml"]