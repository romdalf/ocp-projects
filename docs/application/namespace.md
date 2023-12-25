

A ```project``` is a logical space created on an OpenShift cluster corresponding to a ```namespace``` to host one or more pods that will have direct interaction with each other. 
As an example, a web front-end pod and a database pod could be deployed within a project called ```my-webapp```.

The CLI command to create a ```project``` is:   
```
oc new-project my-webapp --display-name="External web services" --description="External customer facing web service" 
```

The YAML API object definition is:  
```yaml 
--8<-- "sources/project.yaml"
```

Considering an application lifecycle, three projects could be created on a single OCP cluster:  

* ```my-webapp-dev```; where all the development will happen and each pull request from a git perspective could trigger pipeline update to build and deploy a new version of the application.

```yaml
--8<-- "sources/gitops/dev/project.yaml"
```

* ```my-webapp-tst```; where a release candidate would be deployed and goes through the functional and non-functional testing.

```yaml
--8<-- "sources/gitops/tst/project.yaml"
```
  
* ```my-webapp-prd```; where a release candidate has been validated and cut as the next go to release.

```yaml
--8<-- "sources/gitops/prd/project.yaml"
```