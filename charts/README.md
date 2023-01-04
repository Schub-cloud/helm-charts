# Deploy Sample Application for RabbitMQ to test the HPA with custom metrics

## Deploy Charts
```
helm upgrade --install prometheus-operator ./prometheus-operator-crds/
helm upgrade --install rabbitmq ./rabbitmq -n monitoring
helm upgrade --install prometheus ./prometheus/ -n monitoring
helm upgrade --install prometheus-adapter ./prometheus-adapter/  -n monitoing
```
## Deploy Grafana
```
helm upgrade --install grafana ./grafana/ -n monitoring
```

## Deploy Sample application
```
helm upgrade --install rabbitmq-scaling-demo-app ./rabbitmq-sample-app -n monitoing
```
## Deploy HPA
```
kubectl apply -f ./hpa-custom-metric.yaml -n monitoing
```

## Priority Class
To use this priority class in you pods we need to update the parameter in values.yaml file to **priorityClassName:""** to the name of priority class 
```
kubectl apply -f ./priority-class.yaml -n monitoring
```
