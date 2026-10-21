#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

CLUSTER=lab8
CONTEXT="kind-${CLUSTER}"

kubectl --context "$CONTEXT" apply -f mcp/rbac.yaml

TOKEN=$(kubectl --context "$CONTEXT" -n mcp create token mcp-reader --duration=168h)
CA=$(kubectl config view --raw -o jsonpath="{.clusters[?(@.name==\"${CONTEXT}\")].cluster.certificate-authority-data}")

cat > mcp/kubeconfig <<KUBECONFIG
apiVersion: v1
kind: Config
clusters:
  - name: ${CLUSTER}
    cluster:
      server: https://${CLUSTER}-control-plane:6443
      certificate-authority-data: ${CA}
users:
  - name: mcp-reader
    user:
      token: ${TOKEN}
contexts:
  - name: ${CLUSTER}
    context:
      cluster: ${CLUSTER}
      user: mcp-reader
      namespace: shop
current-context: ${CLUSTER}
KUBECONFIG

chmod 0600 mcp/kubeconfig
echo "mcp/kubeconfig written (server https://${CLUSTER}-control-plane:6443, token valid 7d)"
