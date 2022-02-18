variable "project_id" {
  type = string
}

variable "env" {
  type    = string
  default = "dev"
}

variable "mmvs" {
  type = list(map(any))
  default = [{
    "dataset_name" = "value",
    "mmv_name"     = "value",
    "statement"    = "value"
  }]
}

variable "access_token" {
  type = string
}

locals {
  achilio_env_project_id = format("achilio-%s", var.env)
  achilio_env_bucket     = format("achilio_%s_executor_tfstate", var.env)
  achilio_env_prefix     = format("terraform_%s/tfstate", var.env)
}

provider "google" {
  project      = var.project_id
  access_token = var.access_token
}

terraform {
  backend "gcs" {
    bucket = "achilio_executor_tfstate"
    prefix = "terraform/state"
  }
}

resource "google_bigquery_table" "mmv" {
  count               = length(var.mmvs)
  dataset_id          = var.mmvs[count.index]["dataset_name"]
  table_id            = var.mmvs[count.index]["mmv_name"]
  deletion_protection = false

  labels = {
    env          = terraform.workspace
    generated_by = "achilio"
  }
  materialized_view {
    query = var.mmvs[count.index]["statement"]
  }
}
