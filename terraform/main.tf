variable "project_id" {
  type = string
}

variable "region_id" {
  type = string
}

variable "dataset_name" {
  type = string
}

variable "queries" {
  type = map
}

variable "access_token" {
  type = string
}

provider "google" {
  project     = var.project_id
  # region      = var.region_id
  # credentials = "/home/nico/Workspace/Achilio/dev/achilio-dev-919cf3ede1d0.json"
  access_token = var.access_token
}

terraform {
  backend "s3" {
    bucket = "achilio-tfstate"
    key    = "terraform/state"
    region = "eu-west-1"
  }
}

resource "google_bigquery_table" "mmv" {
  dataset_id          = var.dataset_name
  for_each            = var.queries
  table_id            = each.key
  deletion_protection = false

  labels = {
    env          = terraform.workspace
    generated_by = "achilio"
  }
  materialized_view {
    query = each.value
  }
}
