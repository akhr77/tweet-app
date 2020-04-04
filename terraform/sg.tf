module "http_sg" {
  source      = "./security_group"
  name        = "http-sg"
  vpc_id      = aws_vpc.favpic-vpc.id
  port        = 80
  cidr_blocks = ["0.0.0.0/0"]
}

module "https_sg" {
  source      = "./security_group"
  name        = "https-sg"
  vpc_id      = aws_vpc.favpic-vpc.id
  port        = 443
  cidr_blocks = ["0.0.0.0/0"]
}

module "web_front_sg" {
  source      = "./security_group"
  name        = "web_front_sg"
  vpc_id      = aws_vpc.favpic-vpc.id
  port        = 8080
  cidr_blocks = ["0.0.0.0/0"]
}

module "web_backend_sg" {
  source      = "./security_group"
  name        = "web_backend_sg"
  vpc_id      = aws_vpc.favpic-vpc.id
  port        = 8082
  cidr_blocks = ["0.0.0.0/0"]
}

module "http_redirect_sg" {
  source      = "./security_group"
  name        = "http_redirect_sg"
  vpc_id      = aws_vpc.favpic-vpc.id
  port        = 8085
  cidr_blocks = ["0.0.0.0/0"]
}


