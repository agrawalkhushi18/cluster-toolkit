<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
Copyright 2024 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3 |
| <a name="requirement_kubernetes"></a> [kubernetes](#requirement\_kubernetes) | >= 2.20 |
| <a name="requirement_local"></a> [local](#requirement\_local) | >= 2.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_kubernetes.in_cluster"></a> [kubernetes.in\_cluster](#provider\_kubernetes.in\_cluster) | >= 2.20 |
| <a name="provider_local"></a> [local](#provider\_local) | >= 2.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [kubernetes_manifest.this](https://registry.terraform.io/providers/hashicorp/kubernetes/latest/docs/resources/manifest) | resource |
| [local_file.kubeconfig](https://registry.terraform.io/providers/hashicorp/local/latest/docs/resources/file) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cluster_ca_certificate"></a> [cluster\_ca\_certificate](#input\_cluster\_ca\_certificate) | Base64 encoded certificate authority data for the GKE cluster. | `string` | n/a | yes |
| <a name="input_cluster_dependency"></a> [cluster\_dependency](#input\_cluster\_dependency) | A resource from the cluster to depend on, ensuring it's created before this module runs. | `any` | n/a | yes |
| <a name="input_cluster_endpoint"></a> [cluster\_endpoint](#input\_cluster\_endpoint) | The endpoint for the GKE cluster's API server. | `string` | n/a | yes |
| <a name="input_cluster_token"></a> [cluster\_token](#input\_cluster\_token) | A valid authentication token for the GKE cluster. | `string` | n/a | yes |
| <a name="input_content"></a> [content](#input\_content) | Direct content of a YAML manifest. Has precedence over source\_path. | `string` | `null` | no |
| <a name="input_source_path"></a> [source\_path](#input\_source\_path) | Path to a single manifest file (.yaml or .tftpl) or a directory of manifests. For a directory, the path must end with a '/'. | `string` | `null` | no |
| <a name="input_template_vars"></a> [template\_vars](#input\_template\_vars) | A map of variables to be used when rendering .tftpl template files. | `map(any)` | `{}` | no |

## Outputs

No outputs.
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
