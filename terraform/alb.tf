resource "aws_lb" "favpic-lb" {
  name                       = "favpic-alb"
  load_balancer_type         = "application"
  internal                   = false
  idle_timeout               = 60
  enable_deletion_protection = true

  subnets = [
    aws_subnet.public-subnet-1a.id,
    aws_subnet.public-subnet-1c.id
  ]



  security_groups = [
    module.http_sg.security_group_id,
    module.https_sg.security_group_id,
    module.http_redirect_sg.security_group_id
  ]
}

resource "aws_lb_target_group" "favpic" {
  name                 = "favpic"
  target_type          = "ip"
  vpc_id               = aws_vpc.favpic-vpc.id
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 300

  health_check {
    path                = "/"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = "200"
    port                = "traffic-port"
    protocol            = "HTTP"
  }
  depends_on = [aws_lb.favpic-lb]
}


resource "aws_lb_listener" "favpic-https" {
  load_balancer_arn = aws_lb.favpic-lb.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.favpic-certificate.arn
  ssl_policy        = "ELBSecurityPolicy-2016-08"

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "これはHTTPSです"
      status_code  = "200"
    }
  }
}

resource "aws_lb_listener" "favpic-http" {
  load_balancer_arn = aws_lb.favpic-lb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "これはHTTPです"
      status_code  = "200"
    }
  }
}

resource "aws_lb_listener" "redirect_http_to_https" {
  load_balancer_arn = aws_lb.favpic-lb.arn
  port              = 8085
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = 443
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_listener_rule" "favpic" {
  listener_arn = aws_lb_listener.favpic-https.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.favpic.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/*"]
  }
}

output "alb_dns_name" {
  value = aws_lb.favpic-lb.dns_name
}
