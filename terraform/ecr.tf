resource "aws_ecr_repository" "favpic-vue" {
  name = "favpic-vue"
}

resource "aws_ecr_lifecycle_policy" "favpic-vue" {
  repository = aws_ecr_repository.favpic-vue.name

  policy = <<EOF
  {
    "rules": [
      {
        "rulePriority": 1,
        "description": "Keep last 30 release tagged images",
        "selection": {
          "tagStatus": "tagged",
          "tagPrefixList": ["release"],
          "countType": "imageCountMoreThan",
          "countNumber": 30
        },
        "action": {
          "type": "expire"
        }
      }
    ]
  }
EOF
}

resource "aws_ecr_repository" "favpic-go" {
  name = "favpic-go"
}

resource "aws_ecr_lifecycle_policy" "favpic-go" {
  repository = aws_ecr_repository.favpic-go.name

  policy = <<EOF
  {
    "rules": [
      {
        "rulePriority": 1,
        "description": "Keep last 30 release tagged images",
        "selection": {
          "tagStatus": "tagged",
          "tagPrefixList": ["release"],
          "countType": "imageCountMoreThan",
          "countNumber": 30
        },
        "action": {
          "type": "expire"
        }
      }
    ]
  }
EOF
}




