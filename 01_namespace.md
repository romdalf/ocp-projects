---
title: Projects
layout: default
nav_order: 1
---

A ```project``` is a logical space created on an OpenShift cluster to host one or more pods that will have direct interaction with each other. 
A good example of this would be a web front-end pod and a database pod being deployed within a project called "my-webapp".

```
oc new-project my-webapp --display-name="External web services" --description="External customer facing web service" 
```

```YAML
kind: Project
apiVersion: project.openshift.io/v1
metadata:
  name: my-webapp
  annotations:
    openshift.io/description: External customer facing web service
    openshift.io/display-name: External web services
```

Considering an application lifecycle, and our interest in testing, we can create three projects;   
* ```my-webapp-dev```; where all the development will happen and each pull request from a git perspective will trigger a pipeline to build and deploy a new version of the application.

```YAML
kind: Project
apiVersion: project.openshift.io/v1
metadata:
  name: my-webapp-dev
  annotations:
    openshift.io/description: Development of External customer facing web service
    openshift.io/display-name: DEV EWS
```

* ```my-webapp-tst```; where a release candidate would be deployed and goes through the functional and non-functional testing.

```YAML
kind: Project
apiVersion: project.openshift.io/v1
metadata:
  name: my-webapp-tst
  annotations:
    openshift.io/description: Test environment for External customer facing web service
    openshift.io/display-name: TST EWS
```
  
* ```my-webapp-prd```; where a release candidate has been validated and cut as the next go to release.

```YAML
kind: Project
apiVersion: project.openshift.io/v1
metadata:
  name: my-webapp-prd
  annotations:
    openshift.io/description: Production Environment for External customer facing web service
    openshift.io/display-name: PRD EWS
```