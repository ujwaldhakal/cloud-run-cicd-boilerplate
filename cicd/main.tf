provider "google-beta" {
  credentials = "${file(var.credential_file_path)}"
  project = var.project
  region = var.region
  version = "~> 3.43"
}

terraform {
  backend "gcs" {
    prefix = "api"
  }
}

locals {
  app_image = "gcr.io/${var.github_repo_name}/${var.github_repo}:${var.image_tag}"
}
resource "google_cloud_run_service" "default" {
  name = var.app_name
  location = "us-central1"
  provider = "google-beta"

  template {
    spec {
      containers {
        image = local.app_image
      }
    }
  }

  traffic {
    percent = 100
    latest_revision = true
  }
}


data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = "us-central1"
  project     = "pvstaging"
  provider = "google-beta"
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

