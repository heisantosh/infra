resource "google_storage_bucket" "loki_storage_bucket" {
  name     = "${var.gcp_project_id}-loki-storage"
  location = var.gcp_region

  public_access_prevention    = "enforced"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  labels = var.labels

  soft_delete_policy {
    retention_duration_seconds = 0
  }

  lifecycle_rule {
    condition {
      age = 8
    }

    action {
      type = "Delete"
    }
  }
}

# TODO: Not needed anymore, BUT! we may need it for rebuilding templates
resource "google_storage_bucket" "envs_docker_context" {
  name     = "${var.gcp_project_id}-envs-docker-context"
  location = var.gcp_region

  public_access_prevention    = "enforced"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  labels = var.labels
}

resource "google_storage_bucket" "setup_bucket" {
  location = var.gcp_region
  name     = "${var.gcp_project_id}-instance-setup"

  public_access_prevention    = "enforced"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  labels = var.labels
}

resource "google_storage_bucket" "fc_kernels_bucket" {
  location = var.gcp_region
  name     = "${var.gcp_project_id}-fc-kernels"

  public_access_prevention    = "enforced"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  labels = var.labels
}

resource "google_storage_bucket" "fc_versions_bucket" {
  location = var.gcp_region
  name     = "${var.gcp_project_id}-fc-versions"

  public_access_prevention    = "enforced"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  labels = var.labels
}

resource "google_storage_bucket" "fc_env_pipeline_bucket" {
  location = var.gcp_region
  name     = "${var.gcp_project_id}-fc-env-pipeline"

  public_access_prevention    = "enforced"
  storage_class               = "STANDARD"
  uniform_bucket_level_access = true

  labels = var.labels
}

data "google_storage_bucket" "fc_template_bucket" {
  name = var.fc_template_bucket_name
}

resource "google_storage_bucket_iam_member" "loki_storage_iam" {
  bucket = google_storage_bucket.loki_storage_bucket.name
  role   = "roles/storage.objectUser"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "envs_docker_context_iam" {
  bucket = google_storage_bucket.envs_docker_context.name
  role   = "roles/storage.objectUser"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "envs_pipeline_iam" {
  bucket = google_storage_bucket.fc_env_pipeline_bucket.name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "instance_setup_bucket_iam" {
  bucket = google_storage_bucket.setup_bucket.name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "fc_kernels_bucket_iam" {
  bucket = google_storage_bucket.fc_kernels_bucket.name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "fc_versions_bucket_iam" {
  bucket = google_storage_bucket.fc_versions_bucket.name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "fc_template_bucket_iam" {
  bucket = data.google_storage_bucket.fc_template_bucket.name
  role   = "roles/storage.objectUser"
  member = "serviceAccount:${var.gcp_service_account_email}"
}

resource "google_storage_bucket_iam_member" "fc_template_bucket_iam_reader" {
  bucket = data.google_storage_bucket.fc_template_bucket.name
  role   = "roles/storage.legacyBucketReader"
  member = "serviceAccount:${var.gcp_service_account_email}"
}
