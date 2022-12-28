#!/bin/bash

##### Development #####

# Frontend
export ENV_FILE="frontend/.env"

[ -f "$ENV_FILE" ] && rm "$ENV_FILE"

echo 'VUE_APP_MSC_EXPERIMENTER_URL=http://127.0.0.1:8001' >> $ENV_FILE
echo 'VUE_APP_MSC_ANALYZER_URL=http://127.0.0.1:8000' >> $ENV_FILE

# Analyzer
export ENV_FILE="analyzer/.env"

[ -f "$ENV_FILE" ] && rm "$ENV_FILE"

echo 'SECRET_KEY="test secret key"' >> $ENV_FILE
echo 'DEBUG=True' >> $ENV_FILE

##### Kubernetes #####

export ENV_FILE="env_vars.sh"

[ -f "$ENV_FILE" ] && rm "$ENV_FILE"

echo '#!/bin/bash' >> env_vars.sh

# General
echo '# General' >> $ENV_FILE
echo 'export MSC_TAG_PREFIX="gcr.io/mastership"' >> $ENV_FILE
echo 'export MSC_TAG_PREFIX_FOR_REPLACEMENT="gcr.io\/mastership"' >> $ENV_FILE
echo 'export MSC_TAG_VERSION="1.0"' >> $ENV_FILE
echo 'export MSC_IMAGE_PULL_POLICY="Always"' >> $ENV_FILE

# MSC Analyzer
echo '# MSC Analyzer' >> $ENV_FILE
echo 'export MSC_ANALYZER_SECRET_KEY=test-secret-key' >> $ENV_FILE

# MSC Reverse Proxy
echo '# MSC Reverse Proxy' >> $ENV_FILE
## Minikube
echo '## Minikube' >> $ENV_FILE
echo 'export MSC_MINIKUBE_CLUSTER_IP=' >> $ENV_FILE

chmod +x $ENV_FILE
