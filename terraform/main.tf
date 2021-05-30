variable "project_id" {
  type = string
}

variable "region_id" {
  type = string
}

variable "dataset_id" {
  type = string
}

variable "queries" {
  type = map
}

provider "google" {
  project     = var.project_id
  region      = var.region_id
  credentials = "/home/nico/Workspace/Achilio/dev/achilio-dev-919cf3ede1d0.json"
  #   access_token = "ya29.a0AfH6SMDvPkDfUqxyeE0G6TjdxLHVpj4ECCr20oHHoNx1ox9ICu3D006xicIGtjRx6AJye1_1p5k-J7AMvzea74UdK0J95MX_n_Sul1oha6sm789aQiEDTqcYj8xHXTBtfOEWHfLngst5vn-xgo2LHAXbGsW0"
}

terraform {
  backend "s3" {
    bucket = "achilio-tfstate"
    key    = "terraform/state"
    region = "eu-west-1"
  }
}

resource "google_bigquery_table" "mmv" {
  dataset_id          = var.dataset_id
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
