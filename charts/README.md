# Deploy Sample Application for RabbitMQ to test the HPA with custom metrics

## Deploy Charts
```
helm upgrade --install prometheus-operator ./charts/prometheus-operator-crds/
helm upgrade --install rabbitmq ./charts/rabbitmq -n monitoring
helm upgrade --install prometheus ./charts/prometheus/ -n monitoring
helm upgrade --install prometheus-adapter ./charts/prometheus-adapter/  -n monitoing
helm upgrade --install grafana ./charts/grafana/ -n monitoring
```

## Deploy application
```
helm upgrade --install rabbitmq-scaling-demo-app ./charts/rabbitmq-sample-app -n monitoing
```
## Deploy HPA
```
kubectl apply -f ./charts/hpa-custom-metric.yaml -n monitoing
```

## Priority Class
To use this priority class in you pods we need to update the parameter in values.yaml file to **priorityClassName:""** to the name of priority class 
```
kubectl apply -f ./charts/priority-class.yaml -n monitoring
```
