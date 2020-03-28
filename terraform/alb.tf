resource "aws_lb" "favpic-lb" {
  name                       = "favpicLB"
  load_balancer_type         = "application"
  internal                   = false
  idle_timeout               = 60
  enable_deletion_protection = true

  subnets = [

  ]
}

resource "aws_lb_target_group" "favpic-http" {
  name     = "favpicGP"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.favpic-vpc.id

  health_check {
    enabled             = true
    interval            = 30
    path                = "/helth.html"
    port                = 80
    protocol            = "HTTP"
    matcher             = "200"
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 2
  }
}

resource "aws_lb_listener" "favpic-https" {
  load_balancer_arn = aws_lb.favpic-lb.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.favpic-certificate.arn

  default_action {
    target_group_arn = aws_lb_target_group.favpic-http.arn
    type             = "forward"

  }
}
