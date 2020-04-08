resource "aws_route53_zone" "favpic" {
  name = "favpic.tech"
}

# resource "aws_acm_certificate" "favpic-acm-certificate" {
#   domain_name       = aws_route53_record.favpic-route53-record.name
#   validation_method = "DNS"
# }

resource "aws_route53_record" "favpic-route53-record" {
  zone_id = aws_route53_zone.favpic.zone_id
  name    = aws_route53_zone.favpic.name
  type    = "A"

  alias {
    name                   = aws_lb.favpic-lb.dns_name
    zone_id                = aws_lb.favpic-lb.zone_id
    evaluate_target_health = true
  }
}

output "domain_name" {
  value = aws_route53_record.favpic-route53-record.name
}
