# Pending: shop/api-c8cbf8567-lkrg9

Алерт: PodNotScheduled. Категория: anti_affinity.

## Что сказал планировщик

0/3 nodes are available: 1 node(s) had untolerated taint(s), 2 node(s) didn't match pod anti-affinity rules. preemption: 0/3 nodes are available: 1 Preemption is not helpful for scheduling, 2 No preemption victims found for incoming pod.

## Доказательства

- replicas в кластере 3, в values 2: ОТЛИЧАЕТСЯ | Сравнение манифеста в кластере с values.yaml из репозитория
- podAntiAffinity в кластере required, в values preferred: ОТЛИЧАЕТСЯ | Сравнение манифеста в кластере с values.yaml из репозитория
- nodeSelector в кластере нет, в values нет: совпадает | Сравнение манифеста в кластере с values.yaml из репозитория
- requests cpu 50m, memory 64Mi в кластере и в values совпадают | Сравнение манифеста в кластере с values.yaml из репозитория
- 1 нода с taint, 2 воркера | Под shop/api-c8cbf8567-lkrg9

## Расхождение с репозиторием

replicas в кластере 3, в values 2: ОТЛИЧАЕТСЯ
podAntiAffinity в кластере required, в values preferred: ОТЛИЧАЕТСЯ

## Нарушено правило размещения

1. Требуемая (required) podAntiAffinity по hostname разрешена, только если replicas не больше числа воркеров. Для раскладки по нодам используйте preferred anti-affinity или topologySpreadConstraints.

## Что сделать

В values.yaml сервиса api в репозитории platform-config указано required podAntiAffinity, что противоречит правилу 1 из runbook. В кластере же указано preferred podAntiAffinity. Рекомендуется выкатить сервис api из репозитория platform-config и убедиться, что в манифесте указано preferred podAntiAffinity. Если нарушение вызвано ручной правкой в кластере, рекомендуется выкатить сервис заново и найти источник ручного изменения.

## Источники

- pods_get shop/api-c8cbf8567-lkrg9
- events_list shop
- resources_list apps/v1 Deployment shop
- /knowledge/runbook-pending-pods.md
- /knowledge/services/api/values.yaml
