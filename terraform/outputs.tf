output "s3_bucket_id" {
  value = "http://${aws_s3_bucket.website.id}.s3-website.eu-central-1.amazonaws.com/"
}
