# Terraform Infrastructure for DevOps Go Lab

This Terraform code creates an AWS EC2 instance to host the DevOps Go Lab application.

## Requirements

- Terraform >= 1.0 installed
- AWS CLI configured with permissions to create instances

## Variables

- `region` — AWS region (default: us-east-1)
- `instance_type` — EC2 instance type (default: t2.micro)
- `instance_name` — Name of the instance (default: devops-go-lab-instance)

## Usage

1. Initialize Terraform:

```bash
terraform init
```

2. Review the execution plan:

```bash
terraform plan
```

3. Apply the configuration:

```bash
terraform apply
```

4. Destroy the infrastructure:

```bash
terraform destroy
```