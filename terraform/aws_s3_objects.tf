
resource "aws_s3_object" "website" {
  for_each = fileset("../web/.output/public", "**")

  bucket = aws_s3_bucket.website.id
  key    = each.value
  source = "../web/.output/public/${each.value}"
  # etag makes the file update when it changes; see https://stackoverflow.com/questions/56107258/terraform-upload-file-to-s3-on-every-apply
  etag         = filemd5("../web/.output/public/${each.value}")
  content_type = "text/html"

  depends_on = [null_resource.website]
}
