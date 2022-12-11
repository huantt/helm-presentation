## Create new helm chart
```shell
helm create {NAME} [flags]
```

This command creates a chart directory along with the common files and
directories used in a chart.

For example, 'helm create foo' will create a directory structure that looks
something like this:
```
    foo/
    ├── .helmignore   # Contains patterns to ignore when packaging Helm charts.
    ├── Chart.yaml    # Information about your chart
    ├── values.yaml   # The default values for your templates
    ├── charts/       # Charts that this chart depends on
    └── templates/    # The template files
        └── tests/    # The test files
```

## Build dependencies
Build is used to reconstruct a chart's dependencies to the state specified in the lock file. 
In the directory of chart, run the following command:
```shell
helm dependency build
```

## Render chart templates locally and display the output
In the directory of chart, run the following command:
```shell
helm template ./ --debug > preview.yaml
```

## Install chart
```shell
 helm install -f {PATH_TO_VALUES_FILE} {RELEASE_NAME} {CHART}
```
The `CHART` argument must be a chart reference, a path to a packaged chart, a path to an unpacked chart directory or a URL.

### For example:
```shell
 helm install -f aggregator/values.yaml aggreagtor-rest-service ./aggregator
```

## Upgrade chart
```shell
helm upgrade --set {KEY}={VALUE} -f {PATH_TO_VALUES_FILE} {RELEASE_NAME} {CHART}
```
You can specify the `--set` and `-f` flags multiple times.

### For example:
```shell
helm upgrade --set backend.deployment.replicas=2 aggreagtor-rest-service ./aggregator
```