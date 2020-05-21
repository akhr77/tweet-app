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

resource "aws_ecr_repository" "favpic-nginx" {
  name = "favpic-nginx"
}

resource "aws_ecr_lifecycle_policy" "favpic-nginx" {
  repository = aws_ecr_repository.favpic-nginx.name

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

resource "aws_ecr_repository" "favpic-db" {
  name = "favpic-db"
}

resource "aws_ecr_lifecycle_policy" "favpic-db" {
  repository = aws_ecr_repository.favpic-db.name

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




