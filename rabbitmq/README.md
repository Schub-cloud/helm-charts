# RabbitMQ Cluster

This Helm chart creates an instance of the RabbitmqCluster CRD from the RabbitMQ Operator. Since no official Helm chart is available yet, we have created and maintained this chart to provide to provision the instance and include additional features.

Addtional Features:
- ServiceMonitor for Prometheus-based monitoring.
- PodDisruptionBudgets (PDBs)
- Kyverno Policy to propagate RabbitMQ user credentials across multiple namespaces.
