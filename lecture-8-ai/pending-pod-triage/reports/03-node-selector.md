# Pending: shop/api-6874bb6cc-sclzg

Алерт: PodNotScheduled. Категория: node_selector.

## Что сказал планировщик

0/3 nodes are available: 1 node(s) had untolerated taint(s), 2 node(s) didn't match Pod's node affinity/selector. preemption: 0/3 nodes are available: 3 Preemption is not helpful for scheduling.

## Доказательства

- 1 node(s) had untolerated taint(s) | status.conditions.message
- 2 node(s) didn't match Pod's node affinity/selector | status.conditions.message
- nodeSelector: disktype: nvme, preemptionPolicy: PreemptLowerPriority, priority: 0, restartPolicy: Always, schedulerName: default-scheduler, securityContext: {}, serviceAccount: default, serviceAccountName: default, terminationGracePeriodSeconds: 30 | spec.nodeSelector
- replicas: 2 | values.yaml.replicas
- 3 nodes in the cluster | resources_list.v1.Node
- 1 node with taints | resources_list.v1.Node
- Категория исправлена кодом: модель предложила insufficient_resources, по сообщению планировщика это node_selector

## Расхождение с репозиторием

podAntiAffinity: OTLICHAETSYA; nodeSelector: OTLICHAETSYA

## Что сделать

Проверьте манифест в кластере и values.yaml из репозитория. В манифесте отсутствует podAntiAffinity, а nodeSelector содержит нестандартные параметры. Рекомендуется выкатить сервис из репозитория platform-config и убедиться, что манифест соответствует values.yaml. Если нарушения в values.yaml, исправьте их согласно правилам размещения.

## Источники

- pods_get shop/api-6874bb6cc-sclzg
- events_list shop
- resources_list apps/v1 Deployment shop
- /knowledge/runbook-pending-pods.md
- /knowledge/services/api/values.yaml
