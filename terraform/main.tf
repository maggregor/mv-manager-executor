variable "project_id" {
  type = string
}

variable "mmvs" {
  type = list(map(any))
  default = [{
    "dataset_name" = "value",
    "mmv_name"     = "value",
    "statement"    = "value"
  }]
}

variable "service_account" {
  type = string
}

provider "google" {
  project     = var.project_id
  credentials = var.service_account
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
