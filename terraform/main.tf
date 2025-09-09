
resource "null_resource" "deploy" {
  # このリソースは、トリガーが変更されるたびに再作成されます
  triggers = {
    always_run = timestamp()
  }

  connection {
    type        = "ssh"
    user        = var.server_user
    host        = var.server_ip
    port        = var.server_ssh_port
    private_key = var.ssh_private_key
  }

  provisioner "remote-exec" {
    inline = [
      "cd /home/dokkiitech/BridgeMe-Back", 
      "git pull",
      "docker compose build",
      "docker compose up -d",
      "docker exec -it swagger-ui nginx -s reload",
      "echo 'Deployment completed successfully!'"
    ]
  }
}
