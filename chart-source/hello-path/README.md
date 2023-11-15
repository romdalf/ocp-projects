# hello-path helm charts
Simple Golang app deployed using Helm charts. 

## Sources  
Source: https://github.com/romdalf/hello-path
Images: https://github.com/users/romdalf/packages/container/package/hello 

## Deploying
Add the chart repository:
```
helm repo add ocp-projects https://romdalf.github.io/ocp-projects/
```

Install the chart:
```
helm install ocp-projects/hello-path
```




