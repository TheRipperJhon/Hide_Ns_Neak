package main

const state = `
	terraform {
		backend "s3" {
		  bucket         = "hidensneak-terraform"
		  key            = "filename.tfstate"
		  dynamodb_table = "terraform-state-lock-dynamo"
		  region         = "us-east-1"
		  encrypt        = true
		}
	  }
`

const variables = `
	variable "do_token" {}
	variable "aws_access_key" {}
	variable "aws_secret_key" {}
	variable "azure_tenant_id" {}
	variable "azure_client_id" {}
	variable "azure_cosntsclient_secret" {}
	variable "azure_subscription_id" {}
`

///////////////////// MODULES /////////////////////

const ec2Module = `
	module "aws-{{.Region}}" {
		source         = "modules/ec2-deployment"
		aws_region     = "{{.Region}}"
		aws_access_key = "${var.aws_access_key}"
		aws_secret_key = "${var.aws_secret_key}"
		default_sg_name = "{{.SecurityGroup}}"
		aws_keypair_file     = "{{.KeypairFile}}"
		aws_keypair_name     = "{{.KeypairName}}"
		aws_new_keypair      = "{{.NewKeypair}}"
		region_count   = {{.Count}}
	}
`

const azureCdnModule = `
	module "azure-cdn-{{.Endpoint}}" {
		source                  = "modules/azure-cdn-deployment"
		azure_subscription_id   = "${var.azure_subscription_id}"
		azure_tenant_id         = "${var.azure_tenant_id}"
		azure_client_id         = "${var.azure_client_id}"
		azure_client_secret     = "${var.azure_client_secret}"
		azure_cdn_hostname      = "{{.HostName}}"
		azure_cdn_profile_name  = "{{.ProfileName}}"
		azure_cdn_endpoint_name = "{{.EndpointName}}"
		azure_location          = "{{.Location}}"
	}
`

//TODO: need to run removeSpaces() on region
const azureModule = `
	module "azure-{{.LOOKATTODO}}" {
		source                = "modules/azure-deployment"
		azure_subscription_id = "${var.azure_subscription_id}"
		azure_tenant_id       = "${var.azure_tenant_id}"
		azure_client_id       = "${var.azure_client_id}"
		azure_client_secret   = "${var.azure_client_secret}"
		azure_location        = "{{.Region}}"
		azure_instance_count  = {{.InstanceCount}}
	}
`

const cloudfrontModule = `
	module "cloudfront-{{.Origin}}" {
		source            = "modules/cloudfront-deployment"
		cloudfront_origin = "{{.Origin}}"
		aws_access_key    = "${var.aws_access_key}"
		aws_secret_key    = "${var.aws_secret_key}"
	}
`

const digitalOceanModule = `
	module "digital-ocean-{{.Region}}" {
		source          = "modules/droplet-deployment"
		do_token        = "${var.do_token}"
		do_image        = "{{.Image}}"
		pvt_key         = "{{.PrivateKey}}"
		ssh_fingerprint = "{{.SSHFingerprint}}"
		do_region       = "{{.Region}}"
		do_size         = "{{.Size}}"
		do_count        = {{.Count}}
	}
`

const googleCloudModule = `
	module "google-cloud-{{.Region}}" {
		source               = "modules/gcp-deployment"
		gcp_region           = "{{.Region}}"
		gcp_project          = "{{.Project}}"
		gcp_instance_count   = {{.InstanceCount}}
		gcp_ssh_user         = "{{.SSHUser}}"
		gcp_ssh_pub_key_file = "{{.SSHPubKeyFile}}"
	  	gcp_machine_type	 = "{{.MachineType}}"
	}
`

const apiGateway = `
	module "apigateway-{{.TargetUri}}" {
		source 				 = "modules/api-gateway"
		aws_access_key    	 = "${var.aws_access_key}"
		aws_secret_key    	 = "${var.aws_secret_key}"
		aws_api_target_uri 	 = "{{.TargetURI}"
  	}
`
