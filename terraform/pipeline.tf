# data "aws_iam_policy_document" "codebuild" {
#   statement {
#     effect    = "Allow"
#     resources = ["*"]

#     actions = [
#       "s3:PutObject",
#       "s3:GetObject",
#       "s3:GetObjectVersion",
#       "logs:CreateLogGroup",
#       "logs:CreateLogStream",
#       "logs:PutLogEvents",
#       "ecr:GetAuthorizationToken",
#       "ecr:BatchCheckLayerAvailability",
#       "ecr:GetDownloadUrlForLayer",
#       "ecr:GetRepositoryPolicy",
#       "ecr:DescribeRepositories",
#       "ecr:ListImages",
#       "ecr:DescribeImages",
#       "ecr:BatchGetImage",
#       "ecr:InitiateLayerUpload",
#       "ecr:UploadLayerPart",
#       "ecr:CompleteLayerUpload",
#       "ecr:PutImage",
#     ]
#   }
# }

# data "aws_iam_policy_document" "codepipeline" {
#   statement {
#     effect    = "Allow"
#     resources = ["*"]

#     actions = [
#       "s3:PutObject",
#       "s3:GetObject",
#       "s3:GetObjectVersion",
#       "s3:GetBucketVersioning",
#       "codebuild:BatchGetBuilds",
#       "codebuild:StartBuild",
#       "ecs:DescribeServices",
#       "ecs:DescribeTaskDefinition",
#       "ecs:DescribeTasks",
#       "ecs:ListTasks",
#       "ecs:RegisterTaskDefinition",
#       "ecs:UpdateService",
#       "iam:PassRole",
#     ]
#   }
# }

# module "codebuild_role" {
#   source     = "./iam_role"
#   name       = "codebuild"
#   identifier = "codebuild.amazonaws.com"
#   policy     = data.aws_iam_policy_document.codebuild.json
# }

# module "codepipeline_role" {
#   source     = "./iam_role"
#   name       = "codepipeline"
#   identifier = "codepipeline.amazonaws.com"
#   policy     = data.aws_iam_policy_document.codepipeline.json
# }

# resource "aws_codebuild_project" "favpic" {
#   name         = "favpic"
#   service_role = module.codebuild_role.iam_role_arn

#   source {
#     type = "CODEPIPELINE"
#   }

#   artifacts {
#     type = "CODEPIPELINE"
#   }

#   environment {
#     type            = "LINUX_CONTAINER"
#     compute_type    = "BUILD_GENERAL1_SMALL"
#     image           = "aws/codebuild/standard:2.0"
#     privileged_mode = true
#   }
# }

# resource "aws_s3_bucket" "artifact" {
#   bucket = "artifact-pragmatic-terraform"

#   lifecycle_rule {
#     enabled = true

#     expiration {
#       days = "180"
#     }
#   }
# }

# resource "aws_codepipeline" "favpic" {
#   name     = "example"
#   role_arn = module.codepipeline_role.iam_role_arn

#   stage {
#     name = "Source"

#     action {
#       name             = "Source"
#       category         = "Source"
#       owner            = "ThirdParty"
#       provider         = "GitHub"
#       version          = 1
#       output_artifacts = ["Source"]

#       configuration = {
#         Owner                = "your-github-name"
#         Repo                 = "your-repository"
#         Branch               = "master"
#         PollForSourceChanges = false
#       }
#     }
#   }

#   stage {
#     name = "Build"

#     action {
#       name             = "Build"
#       category         = "Build"
#       owner            = "AWS"
#       provider         = "CodeBuild"
#       version          = 1
#       input_artifacts  = ["Source"]
#       output_artifacts = ["Build"]

#       configuration = {
#         ProjectName = aws_codebuild_project.favpic.id
#       }
#     }
#   }

#   stage {
#     name = "Deploy"

#     action {
#       name            = "Deploy"
#       category        = "Deploy"
#       owner           = "AWS"
#       provider        = "ECS"
#       version         = 1
#       input_artifacts = ["Build"]

#       configuration = {
#         ClusterName = aws_ecs_cluster.favpic-ecs-cluster.name
#         ServiceName = aws_ecs_service.favpic.name
#         FileName    = "imagedefinitions.json"
#       }
#     }
#   }

#   artifact_store {
#     location = aws_s3_bucket.artifact.id
#     type     = "S3"
#   }
# }

# resource "aws_codepipeline_webhook" "favpic" {
#   name            = "favpic"
#   target_pipeline = aws_codepipeline.favpic.name
#   target_action   = "Source"
#   authentication  = "GITHUB_HMAC"

#   authentication_configuration {
#     secret_token = "VeryRandomStringMoreThan20Byte!"
#   }

#   filter {
#     json_path    = "$.ref"
#     match_equals = "refs/heads/{Branch}"
#   }
# }

# provider "github" {
#   organization = "akhr77"
# }

# resource "github_repository_webhook" "favpic" {
#   repository = "favpic"

#   configuration {
#     url          = aws_codepipeline_webhook.favpic.url
#     secret       = "VeryRandomStringMoreThan20Byte!"
#     content_type = "json"
#     insecure_ssl = false
#   }

#   events = ["push"]
# }


