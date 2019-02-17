---
title: First Kubernetes Deployment
date: 2019-02-01
tags: ["kubernetes", "code"]
---

First application on Kubernetes using Kubernetes deployments

<!--more-->

```sh
    $ kubectl run kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1 --port=8080
```

Now, check whether it is running:
```sh
    $ kubectl get pods
    NAME                                   READY     STATUS    RESTARTS   AGE
    kubernetes-bootcamp-5c69669756-wv2rp   1/1       Running   0          11s
```

We can check application logs:
```sh
$ kubectl logs kubernetes-bootcamp-5c69669756-wv2rp
Kubernetes Bootcamp App Started At: 2018-10-20T13:38:41.537Z | Running On:  kubernetes-bootcamp-5c69669756-wv2rp
```
