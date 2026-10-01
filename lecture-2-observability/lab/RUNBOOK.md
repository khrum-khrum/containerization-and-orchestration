# Запуск лабораторной 2 — пошаговая инструкция
---

## Шаг 0. Предусловия

```bash
minikube version   # 1.34+
kubectl version --client
helm version       # 3.14+ или 4.x
docker version
```

Все команды ниже выполняются из каталога `lecture-2-observability/lab`:

```bash
cd lecture-2-observability/lab
```

## Шаг 1. Поднять кластер

```bash
minikube start --driver docker
```

## Шаг 2. Собрать образ `api` внутрь кластера

```bash
docker build -t api:0.1.0 ./api
minikube image load api:0.1.0
```

## Шаг 3. Создать неймспейс

```bash
kubectl create namespace lab-2
```

Удобно:
```bash
export NS=lab-2
```

## Шаг 4. Подключить helm-репозитории

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add grafana https://grafana.github.io/helm-charts
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo update
```

## Шаг 5. Prometheus + Grafana + Alertmanager

```bash
helm upgrade --install kube-prometheus-stack \
  prometheus-community/kube-prometheus-stack \
  --version 91.5.2 \
  -n $NS \
  -f values/kube-prometheus-stack.yaml \
  --wait --timeout 10m
```

## Шаг 6. Loki

```bash
helm upgrade --install loki grafana/loki \
  --version 7.3.0 \
  -n $NS \
  -f values/loki.yaml \
  --wait --timeout 10m
```

## Шаг 7. Агент сбора логов (Grafana Alloy)

```bash
helm upgrade --install alloy grafana/k8s-monitoring \
  --version 4.5.2 \
  -n $NS \
  -f values/k8s-monitoring.yaml \
  --wait --timeout 10m
```

## Шаг 8. Jaeger

```bash
helm upgrade --install jaeger jaegertracing/jaeger \
  --version 4.14.0 \
  -n $NS \
  -f values/jaeger.yaml \
  --wait --timeout 10m
```

## Шаг 9. Сервис `api`

```bash
helm upgrade --install api ./charts/api \
  -n $NS \
  --wait --timeout 5m
```

## Шаг 10. Алерты и приёмник

```bash
helm upgrade --install alerting ./charts/alerting -n $NS --wait
```

## Шаг 11. Karma

```bash
helm upgrade --install karma ./charts/karma -n $NS --wait
```
