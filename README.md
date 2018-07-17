# kubenetes-CRD
Kubernetes CRD demo

### `crd/queuejob.go` 
QueueJob: 定义的是QueueJob的结构，类型、包含库k8s的属性、QueueJob的行为规范、QueueJob的状态

QueueJobSpec: 定义QueueJob的行为状态 1.SchedSpec 指定调度参数 2. TaskSpeces 指定QueueJob的任务规范

TaskSpec: 定义任务的行为状态 1.Seteclor 通过label进行选择 2. Replicas QueueJob中Task数量 3.Template 沿用k8s的pod的template

QueueJobStatus: 定义QueueJob的状态，1-5分别为 Pending、Running、Successed、Failed、MinAvailable

QueueJobList: QueueJob 展示时用的