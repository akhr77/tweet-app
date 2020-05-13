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
    Name = "public-subnet-1a"
  }
}

resource "aws_subnet" "public-subnet-1c" {
  vpc_id                  = aws_vpc.favpic-vpc.id
  cidr_block              = "10.0.11.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1c"
  tags = {
    Name = "public-subnet-1c"
  }
}

resource "aws_subnet" "private-subnet-1a" {
  vpc_id                  = aws_vpc.favpic-vpc.id
  cidr_block              = "10.0.100.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "private-subnet-1a"
  }
}

resource "aws_subnet" "private-subnet-1c" {
  vpc_id                  = aws_vpc.favpic-vpc.id
  cidr_block              = "10.0.101.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1c"
  tags = {
    Name = "private-subnet-1c"
  }
}

resource "aws_internet_gateway" "favpic-ig" {
  vpc_id = aws_vpc.favpic-vpc.id
  tags = {
    Name = "favpic-ig"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.favpic-vpc.id
}

resource "aws_route" "public" {
  route_table_id         = aws_route_table.public.id
  gateway_id             = aws_internet_gateway.favpic-ig.id
  destination_cidr_block = "0.0.0.0/0"
}

# resource "aws_route_table" "favpic-route-table" {
#   vpc_id = aws_vpc.favpic-vpc.id
#   route {
#     cidr_block = "0.0.0.0/0"
#     gateway_id = aws_internet_gateway.favpic-ig.id
#   }
#   tags = {
#     Name = "favpic-route-table"
#   }
# }

resource "aws_route_table_association" "public-1a" {
  subnet_id      = aws_subnet.public-subnet-1a.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public-1c" {
  subnet_id      = aws_subnet.public-subnet-1c.id
  route_table_id = aws_route_table.public.id
}
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.favpic-vpc.id
}

resource "aws_route_table_association" "private" {
  subnet_id      = aws_subnet.private-subnet-1a.id
  route_table_id = aws_route_table.private.id
}

resource "aws_eip" "nat_gateway" {
  vpc        = true
  depends_on = [aws_internet_gateway.favpic-ig]
}

resource "aws_nat_gateway" "favpic" {
  allocation_id = aws_eip.nat_gateway.id
  subnet_id     = aws_subnet.public-subnet-1a.id
  depends_on    = [aws_internet_gateway.favpic-ig]
}

resource "aws_route" "private" {
  route_table_id         = aws_route_table.private.id
  nat_gateway_id         = aws_nat_gateway.favpic.id
  destination_cidr_block = "0.0.0.0/0"
}








