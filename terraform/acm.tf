resource "aws_acm_certificate" "favpic-certificate" {
  domain_name = aws_route53_zone.favpic.name

  tags = {
    "favpic" = "website"
  }
}
