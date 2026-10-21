# Runbook: под висит в Pending

Правила размещения для кластера shop-prod. Платформенная команда (platform@shop) отвечает за ноды, команды сервисов отвечают за свои манифесты.

## Кластер

- Воркеры одинаковые: 4 vCPU, 8 GiB. Control-plane затейнчен, рабочая нагрузка на него не едет.
- Лейблы нод, на которые можно ссылаться в nodeSelector/nodeAffinity: `kubernetes.io/hostname`, `topology.kubernetes.io/zone`, `workload=batch`. Других лейблов на нодах нет, и просить их надо у платформы, а не писать в манифест.

## Правила для сервисов

1. Требуемая (required) podAntiAffinity по hostname разрешена, только если replicas не больше числа воркеров. Для раскладки по нодам используйте preferred anti-affinity или topologySpreadConstraints.
2. Запрос памяти на один под больше 2 GiB согласуется с платформой заранее. Запрос больше 8 GiB никогда не влезет в воркер.
3. nodeSelector и nodeAffinity только по лейблам из списка выше.
4. PriorityClass для сервисов с пользовательским трафиком: `shop-critical`. Фон: `shop-batch`.
5. Запросы ресурсов должны быть такими же, как в values сервиса в репозитории platform-config. Ручные правки через `kubectl edit` считаются дрейфом и откатываются.

## Как разбирать Pending

1. Смотрим `status.conditions` пода: в PodScheduled с reason Unschedulable лежит сообщение планировщика с разбивкой по нодам.
2. Смотрим события пода: FailedScheduling повторяется, пока под не встанет.
3. Сверяем манифест в кластере с values из репозитория. Если расходятся, причина почти всегда в ручной правке.
4. Классифицируем: anti-affinity, insufficient resources, node selector/affinity, taints, PVC, прочее.
5. Пишем автору сервиса, что сломано и как чинить. Если нарушено правило из списка выше, явно это говорим.
