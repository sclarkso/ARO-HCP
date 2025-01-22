using '../templates/mgmt-cluster.bicep'

// AKS
param kubernetesVersion = '{{ .mgmt.aks.kubernetesVersion }}'
param vnetAddressPrefix = '{{ .mgmt.aks.vnetAddressPrefix }}'
param subnetPrefix = '{{ .mgmt.aks.subnetPrefix }}'
param podSubnetPrefix = '{{ .mgmt.aks.podSubnetPrefix }}'
param aksClusterName = '{{ .aksName }}'
param aksKeyVaultName = '{{ .mgmt.aks.etcd.kvName }}'
param aksEtcdKVEnableSoftDelete = {{ .mgmt.aks.etcd.kvSoftDelete }}
param systemAgentMinCount = {{ .mgmt.aks.systemAgentPool.minCount}}
param systemAgentMaxCount = {{ .mgmt.aks.systemAgentPool.maxCount }}
param systemAgentVMSize = '{{ .mgmt.aks.systemAgentPool.vmSize }}'
param aksSystemOsDiskSizeGB = {{ .mgmt.aks.systemAgentPool.osDiskSizeGB }}
param userAgentMinCount = {{ .mgmt.aks.userAgentPool.minCount }}
param userAgentMaxCount = {{ .mgmt.aks.userAgentPool.maxCount }}
param userAgentVMSize = '{{ .mgmt.aks.userAgentPool.vmSize }}'
param userAgentPoolAZCount = {{ .mgmt.aks.userAgentPool.azCount }}
param aksUserOsDiskSizeGB = {{ .mgmt.aks.userAgentPool.osDiskSizeGB }}
param aksClusterOutboundIPAddressIPTags = '{{ .mgmt.aks.clusterOutboundIPAddressIPTags }}'

// Maestro
param maestroConsumerName = '{{ .maestro.consumerName }}'
param maestroEventGridNamespacesName = '{{ .maestro.eventGrid.name }}'
param maestroCertDomain = '{{ .maestro.certDomain }}'

// ACR
param acrPullResourceGroups = ['{{ .global.rg }}']

// Region
param regionalResourceGroup = '{{ .regionRG }}'

// CX KV
param cxKeyVaultName = '{{ .cxKeyVault.name }}'

// MSI KV
param msiKeyVaultName = '{{ .msiKeyVault.name }}'

// MGMT KV
param mgmtKeyVaultName = '{{ .mgmtKeyVault.name }}'

// MI for deployment scripts
param aroDevopsMsiId = '{{ .aroDevopsMsiId }}'

// Azure Monitor Workspace
param azureMonitorWorkspaceName = '{{ .monitoring.workspaceName }}'
