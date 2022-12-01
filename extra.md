# Extra

```bash
istioctl install --set values.global.proxy.privileged=true
kubectl exec --stdin --tty deployment/msc-analyser-deployment --container istio-proxy -- /bin/bash
tcpdump -vvvv -A -i -eth0 '((dst port 9080) and (net 10.56.3.235))'
```
