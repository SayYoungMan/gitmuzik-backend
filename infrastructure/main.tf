# Terraform file to run Docker Image on ECS Cluster
# Followed the tutorial: https://medium.com/devops-engineer-documentation/terraform-deploying-a-docker-image-to-an-aws-ecs-cluster-3931337e82fb

# Configure the AWS Provider
provider "aws" {
    region = "eu-west-2"
}

# Create ECR Repository to store Docker Image
resource "aws_ecr_repository" "gitmuzik-backend" {
    name = "gitmuzik-backend"
}

# Create ECS Cluster
resource "aws_ecs_cluster" "gitmuzik-cluster" {
    name = "gitmuzik-cluster"
}

# Create task definition that describes the container
resource "aws_ecs_task_definition" "gitmuzik-backend-task" {
    family = "gitmuzik-backend-task"
    container_definitions = <<DEFINITION
    [
        {
            "name": "gitmuzik-backend-task",
            "image": "${aws_ecr_repository.gitmuzik-backend.repository_url}",
            "essential": true,
            "memory": 512,
            "cpu": 256
        }
    ]
    DEFINITION
    requires_compatibilities = ["FARGATE"]
    network_mode = "awsvpc"
    memory = 512
    cpu = 256
    execution_role_arn = "${aws_iam_role.ecsTaskExecutionRole.arn}"
}

# Provide IAM roles so that tasks have correct permission to execute
resource "aws_iam_role" "ecsTaskExecutionRole" {
    name = "ecsTaskExecutionRole"
    assume_role_policy = "${data.aws_iam_policy_document.assume_role_policy.json}"
}

data "aws_iam_policy_document" "assume_role_policy" {
    statement {
      actions = ["sts:AssumeRole"]

      principals {
          type = "Service"
          identifiers = [ "ecs-tasks.amazonaws.com" ]
      }
    }
}

# Create services for container
resource "aws_ecs_service" "gitmuzik-service" {
    name = "gitmuzik-service"
    cluster = "${aws_ecs_cluster.gitmuzik-cluster.id}"
    task_definition = "${aws_ecs_task_definition.gitmuzik-backend-task.arn}"
    launch_type = "FARGATE"
    desired_count = 1

    network_configuration {
        subnets = ["${aws_default_subnet.default_subnet_a.id}", "${aws_default_subnet.default_subnet_b.id}", "${aws_default_subnet.default_subnet_c.id}"]
        assign_public_ip = true
    }
}

# Reference to our default VPC
resource "aws_default_vpc" "default_vpc" {
}

# Reference to our default subnets
resource "aws_default_subnet" "default_subnet_a" {
    availability_zone = "eu-west-2a"
}
resource "aws_default_subnet" "default_subnet_b" {
    availability_zone = "eu-west-2b"
}
resource "aws_default_subnet" "default_subnet_c" {
    availability_zone = "eu-west-2c"
}