# Overview

## Prerequisites

* Requires `go` setup.

## Local Run

* Build docker image

  ```shell script
  make build-image
  ```

* Run

  ```shell script
   docker run -p 8080:8080 gabbi/docusign-event-gw:kyma-1.15 --basic-auth-enabled=false --app-name=docusign --event-publish-url=<provide a mock bin url>
  ```

* Verify

  ```shell script
  curl -X POST http://localhost:8080  -d @./test-data/1.xml
  ```

* Push the image

  ```shell script
  make push-image
  ```

## Deploy on Kyma Cluster

* Create Secret required for basic auth

  ```shell script
  kc -n <ns> create secret generic docusign-event-gw \
  --from-literal=USERNAME=<user> --from-literal=PASSWORD=<password> \
  --from-literal=APP_NAME=<application-name>
  ```

* Deploy service and deployment along with API Rule

  ```shell script
  kubectl apply -f k8s/
  ```

  The `docusign-event-gw` is accessible on <https://{docusign-event-gw}.cluster-domain>