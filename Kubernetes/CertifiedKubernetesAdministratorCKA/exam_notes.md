### 22: Practice test - Pods
- OK

### 26: Practice test - ReplicaSets
- OK

### 30: Practice test - Deployments
- OK

### 32: Practice test - Namespaces
- OK

### 35: Practice test - Services
- OK

### 37: Practice test - Imperative commands
- OK

### 43: Practice test - Manual scheduling
- OK

### 45: Practice test - Labels and Selectors
- OK

### 48: Practice test - Taints and Tolerations
- OK

### 51: Practice test - Node Affinity
- OK

### 55: Practice test - Resource requirements and Limits
- OK

### 58: Practice test - Daemonsets
- OK

### 60: Practice test - Static Pods
- OK

### 62: Practice test - Multiple Schedulers
- Revisit running multiple schedulers
- https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/

### 69: Practice test - Monitoring
- OK

### 72: Practice test - Monitor application logs
- OK

### 78: Practice test - Rolling Updates and Rollbacks 
- OK

### 83: Practice test - Commands and Arguments
- OK

### 86: Practice test - Environment variables
- OK

### 89: Practice test - Secrets
- OK 

### 92: Practice test - Multi-container pods
- OK

### 95: Practice test - Init Containers
- OK

### 101: Practice test - OS Upgrades
- OK

### 105: Practice test - Cluster upgrade
- OK. Do an upgrade on your own if you want more proficiency (with no assists on question text)

### 107: Practice test - Backup and Restore methods
- etcd backup restore with etcdctl is difficult. You need to be able to backup & restore only with the standard doc!

### 121 : Practice test - View certificates
- Renewing a certificate with openssl was very difficult with the standard doc. Not sure if they're gonna ask about that on the exam...

### 123: Practice test - Certificates API
- OK

### 125: Practice test - Kubeconfig
- OK

### 129: Practice test - RBAC
- OK

### 131: Cluster roles and Role Bindings
- OK

### 133: Practice test - Image security
- OK

### 135: Practice test - Security contexts
- OK

### 137: Practice test - Network Policy
- OK

### 148: Practice test: Persistent Volumes and Persistent Volume Claims
- OK

### 161: Practice test - Explore Kubernetes Environment
- OK

### 164: Practice test - Explore CNI Weave
- OK

### 166: Practice test - Explore CNI Weave - 2
- OK

### 167: Practice test - Deploy network solution
- OK

### 170: Practice test - Service Networking
- OK

### 173: Practice test - Explore DNS
- OK

### 175: Practice test - Ingress - 1
- OK

### 177: Practice test - Ingress - 2
- OK

### 203: Practice test - Deploy a Kubernetes Cluster using Kubeadm
- OK - Run it again with flannel if you're curious, exam uses flannel.

### 212: Practice test - Application Failure
- OK - Redo!

### 215: Practice test - Control Plane Failure
- OK - Redo!

### 218: Practice test - Worker Node Failure
- OK. Redo!

### 223: Practice test - Advanced Kubectl Commands
- OK. Practice queries more if you feel like it.

### 225: Mock Exam - 1
- OK - 100
### 226: Mock Exam - 2
- Redo CSR/RBAC and pod DNS - 70
- Pod DNS: How?

### 228: Mock Exam - 3
- Your ass got whooped by service accounts, cluster roles and cluster role bindings
- JSONPath with range some more with spaces and/or tabs
- NetworkPolicies breh
- Iterate over `kubectl api-resources`
- Go through kubectl with just code examples, no lenghty explanations

### Extra:
- When done with Mock exams 2 and 3, make a list of what they covered, and deduce what they did not cover, and practice with that too!

### Things you have to memorize, no two ways about it:
- `ETCDCTL_API=3 ./etcdctl snapshot save snapshot.db --cacert /etc/ssl/etcd/ca.crt --cert /etc/ssl/etcd/client.crt --key /etc/ssl/etcd/client.key` 