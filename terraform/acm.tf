resource "aws_acm_certificate" "favpic-certificate" {
  domain_name               = aws_route53_zone.favpic.name
  subject_alternative_names = []
  validation_method         = "DNS"

  tags = {
    "Name" = "favpic-acm"
  }
}

# resource "aws_route53_record" "favpic_certificate" {
#   name    = aws_acm_certificate.favpic-certificate.domain_validation_options[0].resource_record_name
#   type    = aws_acm_certificate.favpic-certificate.domain_validation_options[0].resource_record_type
#   records = [aws_acm_certificate.favpic-certificate.domain_validation_options[0].resource_record_value]
#   zone_id = aws_route53_zone.favpic.id
#   ttl     = 60
# }

# resource "aws_acm_certificate_validation" "favpic-certificate-validation" {
#   certificate_arn         = aws_acm_certificate.favpic-certificate.arn
#   validation_record_fqdns = [aws_route53_record.favpic_certificate.fqdn]
# }
