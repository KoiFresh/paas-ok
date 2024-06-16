resource "null_resource" "website" {
  provisioner "local-exec" {
    command = "docker run -v $PWD/../web:/web --workdir /web --entrypoint '' node sh -c 'yarn install && yarn generate'"
  }
}
