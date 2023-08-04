## Deployment notes
### Inject linkerd
Mesh all deployment in the namespace
```
kubectl get -n linkerd-examples deploy -o yaml | linkerd inject - | kubectl apply -f -
```
The output will similar to this
```
deployment "foobar-api" injected
deployment "web-api" injected

deployment.apps/foobar-api configured
deployment.apps/web-api configured
```
The number of containers on each pod will be increased too.