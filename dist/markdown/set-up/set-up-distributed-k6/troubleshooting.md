---
aliases:
  - ./troubleshooting # docs/k6/<K6_VERSION>/set-up/set-up-distributed-k6/troubleshooting
weight: 400
title: Troubleshoot
---

# Troubleshoot

This topic includes instructions to help you troubleshoot common issues with the k6 Operator.

If you're using Private Load Zones in Grafana Cloud k6, refer to [Troubleshoot Private Load Zones](https://grafana.com/docs/grafana-cloud/testing/k6/author-run/private-load-zone/troubleshoot/).

## How to troubleshoot


### Test your script locally

Always run your script locally before trying to run it with the k6 Operator:

```bash
k6 run script.js
```

If you're using environment variables or CLI options, pass them in as well:

```bash
MY_ENV_VAR=foo k6 run script.js --tag my_tag=bar
```

That ensures that the script has correct syntax and can be parsed with k6 in the first place. Additionally, running locally can help you check if the configured options are doing what you expect. If there are any errors or unexpected results in the output of `k6 run`, make sure to fix those prior to deploying the script elsewhere.

### `TestRun` deployment

#### The Jobs and Pods

In case of one `TestRun` Custom Resource (CR) creation with `parallelism: n`, there are certain repeating patterns:

1. There will be `n + 2` Jobs (with corresponding Pods) created: initializer, starter, and `n` runners.
1. If any of these Jobs didn't result in a Pod being deployed, there must be an issue with that Job. Some commands that can help here:

   ```bash
   kubectl get jobs -A
   kubectl describe job mytest-initializer
   ```

1. If one of the Pods was deployed but finished with `Error`, you can check its logs with the following command:

   ```bash
   kubectl logs mytest-initializer-xxxxx
   ```

#### `TestRun` with `cleanup` option

If a `TestRun` has the [`spec.cleanup` option](https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/usage/executing-k6-scripts-with-testrun-crd/#clean-up-resources) set, as [`PrivateLoadZone`](https://grafana.com/docs/grafana-cloud/testing/k6/author-run/private-load-zone/) tests always do, for example, it may be harder to locate and analyze the Pod before it's deleted.

In that case, we recommend using observability solutions, like Prometheus and Loki, to store metrics and logs for later analysis.

As an alternative, it's also possible to watch for the resources manually with the following commands:

  ```bash
  kubectl get jobs -n my-namespace -w
  kubectl get pods -n my-namespace -w

  # To get detailed information (this one is quite verbose so use with caution):
  kubectl get pods -n my-namespace -w -o yaml
  ```

#### k6 Operator

Another source of info is the k6 Operator itself. It's deployed as a Kubernetes `Deployment`, with `replicas: 1` by default, and its logs together with observations about the Pods from the previous section usually contain enough information to help you diagnose any issues. With the standard deployment, the logs of the k6 Operator can be checked with:

```bash
kubectl -n k6-operator-system -c manager logs k6-operator-controller-manager-xxxxxxxx-xxxxx
```

#### Inspect `TestRun` resource

After you or `PrivateLoadZone` deployed a `TestRun` CR, you can inspect it the same way as any other resource:

```bash
kubectl describe testrun my-testrun
```

Firstly, check if the spec is as expected. Then, see the current status:

```yaml
Status:
  Conditions:
    Last Transition Time:  2024-01-17T10:30:01Z
    Message:
    Reason:                CloudTestRunFalse
    Status:                False
    Type:                  CloudTestRun
    Last Transition Time:  2024-01-17T10:29:58Z
    Message:
    Reason:                TestRunPreparation
    Status:                Unknown
    Type:                  TestRunRunning
    Last Transition Time:  2024-01-17T10:29:58Z
    Message:
    Reason:                CloudTestRunAbortedFalse
    Status:                False
    Type:                  CloudTestRunAborted
    Last Transition Time:  2024-01-17T10:29:58Z
    Message:
    Reason:                CloudPLZTestRunFalse
    Status:                False
    Type:                  CloudPLZTestRun
  Stage:                   error
```

If `Stage` is equal to `error`, you can check the logs of k6 Operator.

Conditions can be used as a source of info as well, but it's a more advanced troubleshooting option that should be used if the previous steps weren't enough to diagnose the issue. Note that conditions that start with the `Cloud` prefix only matter in the setting of k6 Cloud test runs, for example, for cloud output and PLZ test runs.

#### Debugging k6 process

If the script is working locally as expected, and the previous steps show no errors as well, yet you don't see an expected result of a test and suspect k6 process is at fault, you can use the k6 [verbose option](https://grafana.com/docs/k6/<K6_VERSION>/using-k6/k6-options/#options) in the `TestRun` spec:

```yaml
apiVersion: k6.io/v1alpha1
kind: TestRun
metadata:
  name: k6-sample
spec:
  parallelism: 2
  script:
    configMap:
      name: 'test'
      file: 'test.js'
  arguments: --verbose
```

## Common errors

### Issues with environment variables

Refer to [Environment variables](https://github.com/grafana/k6-operator/blob/main/docs/env-vars.md) for details on how to pass environment variables to the k6 Operator.

### Tags not working

Tags are a rather common source of errors when using the k6 Operator. For example, the following tags would lead to parsing errors:

```yaml
  arguments: --tag product_id="Test A"
  # or
  arguments: --tag foo=\"bar\"
```

You can see those errors in the logs of either the initializer or the runner Pod, for example:

```bash
time="2024-01-11T11:11:27Z" level=error msg="invalid argument \"product_id=\\\"Test\" for \"--tag\" flag: parse error on line 1, column 12: bare \" in non-quoted-field"
```

This is a common problem with escaping the characters. You can find an [issue](https://github.com/grafana/k6-operator/issues/211) in the k6 Operator repository that can be upvoted.

### An error on reading output of the initializer Pod

The k6 runners fail to start, and in the k6 Operator logs, you see the `unable to marshal` error. This can happen for several reasons:

1. Your Kubernetes setup includes some tool that is implicitly adding symbols to the log output of Pods. You can verify this case by checking the logs of the initializer Pod: they should contain valid JSON, generated by k6. Currently, to fix this, the tool adding symbols must be switched off for the k6 Operator workloads.

1. Multi-file script includes many files which all must be fully accessible from the runner Pod. You can verify this case by checking the logs of the initializer Pod: there will be an error about some file not being found. To fix this, refer to [Multi-file tests](https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/usage/executing-k6-scripts-with-testrun-crd/#multi-file-tests) on how to configure multi-file tests in `TestRun`.

1. There are problems with environment variables or with importing an extension. Following the steps found in [testing locally](#test-your-script-locally) and in [troubleshooting extensions](https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/usage/extensions#troubleshooting) can help debug this issue. One additional command that you can use to help diagnose issues with your script is the following:

```bash
k6 inspect --execution-requirements script.js
```

That command is a shortened version of what the initializer Pod is executing. If the command produces an error, there's a problem with the script itself and it should be solved outside of the k6 Operator. The error itself may contain a hint to what's wrong, such as a syntax error.

If the standalone `k6 inspect --execution-requirements` executes successfully, then it's likely a problem with `TestRun` deployment specific to your Kubernetes setup.

### An issue with `volumeClaim`

Storing k6 scripts on a persistent volume is one approach to [multi-file tests](https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/usage/executing-k6-scripts-with-testrun-crd/#multi-file-tests). However, errors can occur due to misconfiguration of the volume. These errors are not within the purview of the k6 Operator; they are inherent to the Kubernetes setup itself, as the k6 Operator only mounts volumes to the Pods. However, here are some general recommendations to help debug such errors.

The `volumeClaim` option is expecting a persistent volume claim, so first, check the [Kubernetes documentation](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) and your infrastructure provider documentation to confirm if the volume is indeed set up correctly and can be mounted by Kubernetes pods.

Then, if the volume appears to be correct and is mounted to the k6 Pods without an issue, yet the `TestRun` fails with an error like the following:

```bash
The moduleSpecifier \"/test/utils.js\" couldn't be found on local disk.
```

This error implies that either the file was not written successfully to the Volume or there is a misconfiguration with a path. So it makes sense to create a separate debug Pod, for example, with the [`busybox` image](https://hub.docker.com/_/busybox) to confirm that the Volume contains the script and all its dependencies. Such a Pod should have a configuration similar to this one:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox
spec:
  volumes:
  - name: test-volume
    volumeSource:
      persistentVolumeClaim:
        claimName: test-pvc
        readOnly: false
  containers:
  - image: busybox
    name: busybox
    imagePullPolicy: IfNotPresent
    command:
      - sleep
      - "3600"
    volumeMounts:
    - mountPath: /test
      name: test-volume
  restartPolicy: Always
```

Then execute `ls /test` on this debug Pod to see which files are present.


### k6 runners do not start

The k6 runners fail to start, and in the k6 Operator logs, you see the error `Waiting for initializing pod to finish`.

In this case, it's most likely that an initializer Pod was not able to start for some reason.

#### How to fix

Refer to [The Jobs and Pods](#the-jobs-and-pods) section to see how to:

1. Check if the initializer Pod has started and finished.
1. See an issue in the initializer Job's description that prevents a Pod from being scheduled.

Once the error preventing the initializer Pod from starting and completing is resolved, redeploy the `TestRun` or, in case of a `PrivateLoadZone` test, restart the k6 process.

### Non-existent ServiceAccount

A ServiceAccount can be defined as `serviceAccountName` in a PrivateLoadZone, and as `runner.serviceAccountName` in a TestRun CRD. If the specified ServiceAccount doesn't exist, k6 Operator will successfully create Jobs but corresponding Pods will fail to be deployed, and the k6 Operator will wait indefinitely for Pods to be `Ready`. This error can be best seen in the events of the Job:

```bash
kubectl describe job plz-test-xxxxxx-initializer
...
Events:
  Warning  FailedCreate  57s (x4 over 2m7s)  job-controller  Error creating: pods "plz-test-xxxxxx-initializer-" is forbidden: error looking up service account plz-ns/plz-sa: serviceaccount "plz-sa" not found
```

k6 Operator doesn't try to analyze such scenarios on its own, but you can refer to the following [issue](https://github.com/grafana/k6-operator/issues/260) for improvements.

#### How to fix

To fix this issue, the incorrect `serviceAccountName` must be corrected, and the `TestRun` or `PrivateLoadZone` resource must be re-deployed.

### Non-existent `nodeSelector`

`nodeSelector` can be defined as `nodeSelector` in a PrivateLoadZone, and as `runner.nodeSelector` in the TestRun CRD.

This case is very similar to the [ServiceAccount](#non-existent-serviceaccount): the Pod creation will fail, but the error is slightly different:

```bash
kubectl describe pod plz-test-xxxxxx-initializer-xxxxx
...
Events:
  Warning  FailedScheduling  48s (x5 over 4m6s)  default-scheduler  0/1 nodes are available: 1 node(s) didn't match Pod's node affinity/selector.
```

#### How to fix

To fix this issue, the incorrect `nodeSelector` must be corrected and the `TestRun` or `PrivateLoadZone` resource must be re-deployed.

### Insufficient resources

A related problem can happen when the cluster does not have sufficient resources to deploy the runners. There's a higher probability of hitting this issue when setting small CPU and memory limits for runners or using options like `nodeSelector`, `runner.affinity` or `runner.topologySpreadConstraints`, and not having a set of nodes matching the spec. Alternatively, it can happen if there is a high number of runners required for the test (via `parallelism` in TestRun or during PLZ test run) and autoscaling of the cluster has limits on the maximum number of nodes, and can't provide the required resources on time or at all.

This case is somewhat similar to the previous two: the k6 Operator will wait indefinitely and can be monitored with events in Jobs and Pods. If it's possible to fix the issue with insufficient resources on-the-fly, for example, by adding more nodes, k6 Operator will attempt to continue executing a test run.

### OOM of a runner Pod

If there's at least one runner Pod that OOM-ed, the whole test will be [stuck](https://github.com/grafana/k6-operator/issues/251) and will have to be deleted manually:

```bash
kubectl delete testrun my-test
```

A `PrivateLoadZone` test or a `TestRun` [with cloud output](https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/usage/k6-operator-to-gck6/#cloud-output) will be aborted by Grafana Cloud k6 after its expected duration is up.

#### How to fix

In the case of OOM, review your k6 script to understand what kind of resource usage the script requires. It may be that the k6 script can be improved to be more performant. Then, set the `spec.runner.resources` in the `TestRun` CRD, or `spec.resources` in the `PrivateLoadZone` CRD accordingly.

### Disruption of the k6 runners

A k6 test can be executed for a long time. But depending on the Kubernetes setup, it's possible that the Pods running k6 are disrupted and moved elsewhere during execution. This will skew the test results. In the case of a `PrivateLoadZone` test or a `TestRun` [with cloud output](https://grafana.com/docs/k6/latest/set-up/set-up-distributed-k6/usage/k6-operator-to-gck6/#cloud-output), the test run may additionally be aborted by Grafana Cloud k6 once its expected duration is up, regardless of the exact state of k6 processes.

#### How to fix

Ensure that k6 Pods can't be disrupted by the Kubernetes setup, for example, with [PodDisruptionBudget](https://kubernetes.io/docs/concepts/workloads/pods/disruptions/) and a less aggressive configuration of the autoscaler.

