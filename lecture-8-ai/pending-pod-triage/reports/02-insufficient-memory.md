# Pending: shop/api-5c7bcfbffb-skxgd

Алерт: PodNotScheduled. Категория: insufficient_resources.

## Что сказал планировщик

0/3 nodes are available: 1 node(s) had untolerated taint(s), 2 Insufficient memory. preemption: 0/3 nodes are available: 3 Preemption is not helpful for scheduling.

## Доказательства

- 2 Insufficient memory. | scheduler_message
- 2 | 2 | replicas comparison
- 50m | 50m | resources comparison
- 64Gi | 64Mi | resources comparison

## Расхождение с репозиторием

requests: cpu 50m, memory 64Gi in cluster, cpu 50m, memory 64Mi in values: ОТЛИЧАЕТСЯ

## Что сделать

В манифесте в кластере запрос памяти указан как 64Gi, в values.yaml как 64Mi. Рекомендуется выкатить сервис из репозитория platform-config и проверить, что запрос памяти в манифесте соответствует значению в values.yaml.

## Источники

- Kubernetes_MCP_pods_get
- Kubernetes_MCP_events_list
- Platform_repo_MCP_read_text_file
- Platform_repo_MCP_read_text_file
- pods_get shop/api-5c7bcfbffb-skxgd
- events_list shop
- resources_list apps/v1 Deployment shop
- /knowledge/runbook-pending-pods.md
- /knowledge/services/api/values.yaml
