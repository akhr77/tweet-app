resource "aws_vpc" "favpic-vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "favpic-vpc"
  }
}

resource "aws_subnet" "public-subnet-1a" {
  vpc_id                  = aws_vpc.favpic-vpc.id
  cidr_block              = "10.0.10.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "vpc-sub-pub-favpic-1a"
  }
}

resource "aws_subnet" "private-subnet-1a" {
  vpc_id                  = aws_vpc.favpic-vpc.id
  cidr_block              = "10.0.100.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "vpc-sub-pri-favpic-1a"
  }
}

resource "aws_subnet" "private-subnet-1c" {
  vpc_id                  = aws_vpc.favpic-vpc.id
  cidr_block              = "10.0.101.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1c"
  tags = {
    Name = "vpc-sub-pri-favpic-1c"
  }
}

resource "aws_internet_gateway" "favpic-ig" {
  vpc_id = aws_vpc.favpic-vpc.id
  tags = {
    Name = "favpic-ig"
  }
}

resource "aws_route_table" "favpic-route-table" {
  vpc_id = aws_vpc.favpic-vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.favpic-ig.id
  }
  tags = {
    Name = "favpic-route-table"
  }
}

resource "aws_route_table_association" "fabpic-sbn-rtb" {
  subnet_id      = aws_subnet.public-subnet-1a.id
  route_table_id = aws_route_table.favpic-route-table.id
}





